package job_dispatcher

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"lamda_backend/pkg/blockchain"
	"lamda_backend/pkg/contracts"
	"lamda_backend/pkg/logger"
	"lamda_backend/pkg/nats"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"
)

// Service handles job dispatching operations
type Service struct {
	db                 *gorm.DB
	natsClient         *nats.NATSClient
	blockchain         *blockchain.EVMClient
	logger             *logger.Logger
	contractAddr       string
	jobManagerContract *contracts.JobManager
}

// NewService creates a new job dispatcher service
func NewService(db *gorm.DB, natsClient *nats.NATSClient, blockchain *blockchain.EVMClient, logger *logger.Logger, contractAddr string) *Service {
	return &Service{
		db:           db,
		natsClient:   natsClient,
		blockchain:   blockchain,
		logger:       logger.WithService("job-dispatcher"),
		contractAddr: contractAddr,
	}
}

// Start starts the job dispatcher service
func (s *Service) Start(ctx context.Context) error {
	s.logger.Info("Starting job dispatcher service")

	// Initialize the JobManager contract
	contractAddress := common.HexToAddress(s.contractAddr)
	jobManagerContract, err := contracts.NewJobManager(contractAddress, s.blockchain.GetClient())
	if err != nil {
		return fmt.Errorf("failed to initialize JobManager contract: %w", err)
	}
	s.jobManagerContract = jobManagerContract

	// Subscribe to NATS queries
	if err := s.subscribeToQueries(); err != nil {
		return fmt.Errorf("failed to subscribe to queries: %w", err)
	}

	// Start blockchain event listener
	go s.listenToBlockchainEvents(ctx)

	s.logger.Info("Job dispatcher service started successfully")
	return nil
}

// subscribeToQueries subscribes to NATS queries for job information
func (s *Service) subscribeToQueries() error {
	// Subscribe to jobs.query subject
	_, err := s.natsClient.SubscribeWithReply("jobs.query", s.handleJobQuery)
	if err != nil {
		return fmt.Errorf("failed to subscribe to jobs.query: %w", err)
	}

	s.logger.Info("Subscribed to jobs.query")
	return nil
}

// handleJobQuery handles queries for jobs
func (s *Service) handleJobQuery(data []byte) ([]byte, error) {
	var query JobQuery
	if err := json.Unmarshal(data, &query); err != nil {
		return nil, fmt.Errorf("failed to unmarshal query: %w", err)
	}

	jobs, err := s.GetJobs(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get jobs: %w", err)
	}

	response := JobsResponse{
		Jobs:  jobs,
		Count: len(jobs),
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response: %w", err)
	}

	return responseData, nil
}

// listenToBlockchainEvents listens for blockchain events
func (s *Service) listenToBlockchainEvents(ctx context.Context) {
	s.logger.Info("Starting blockchain event listener (polling mode)")

	// Get the latest block number to start listening from
	latestBlock, err := s.blockchain.GetLatestBlockNumber(ctx)
	if err != nil {
		s.logger.Error("Failed to get latest block number", "error", err)
		return
	}

	// Start from the latest block
	fromBlock := uint64(latestBlock)
	s.logger.Info("Starting event listener from block", "block", fromBlock)

	// Poll for events every 10 seconds
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			s.logger.Info("Stopping blockchain event listener")
			return
		case <-ticker.C:
			// Poll for new events
			if err := s.pollForEvents(ctx, fromBlock); err != nil {
				s.logger.Error("Failed to poll for events", "error", err)
			} else {
				// Update fromBlock to current block
				if currentBlock, err := s.blockchain.GetLatestBlockNumber(ctx); err == nil {
					fromBlock = uint64(currentBlock)
				}
			}
		}
	}
}

// pollForEvents polls for blockchain events
func (s *Service) pollForEvents(ctx context.Context, fromBlock uint64) error {
	// Get current block
	currentBlock, err := s.blockchain.GetLatestBlockNumber(ctx)
	if err != nil {
		return fmt.Errorf("failed to get current block: %w", err)
	}

	// Only process if there are new blocks
	if uint64(currentBlock) <= fromBlock {
		return nil
	}

	// Query for JobCreated events
	jobCreatedEvents, err := s.jobManagerContract.FilterJobCreated(&bind.FilterOpts{
		Start:   fromBlock,
		End:     &currentBlock,
		Context: ctx,
	}, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to filter JobCreated events: %w", err)
	}

	// Process JobCreated events
	for jobCreatedEvents.Next() {
		event := jobCreatedEvents.Event
		s.logger.Info("Received JobCreated event", "job_id", fmt.Sprintf("0x%x", event.JobId))
		if err := s.processJobCreatedEvent(event); err != nil {
			s.logger.Error("Failed to process JobCreated event", "error", err, "job_id", fmt.Sprintf("0x%x", event.JobId))
		}
	}

	return nil
}

// processJobCreatedEvent processes a JobCreated event from the blockchain
func (s *Service) processJobCreatedEvent(event *contracts.JobManagerJobCreated) error {
	s.logger.Info("Processing JobCreated event", "job_id", fmt.Sprintf("0x%x", event.JobId))

	// Convert the event to our internal format
	jobEvent := JobCreatedEvent{
		JobID:           fmt.Sprintf("0x%x", event.JobId),
		RenterAddress:   event.Renter.Hex(),
		ProviderAddress: event.Provider.Hex(),
		PaymentAmount:   event.Payment.String(),
		BlockNumber:     event.Raw.BlockNumber,
		TransactionHash: event.Raw.TxHash.Hex(),
	}

	// Process the event using existing logic
	return s.ProcessJobCreatedEvent(jobEvent)
}

// ProcessJobCreatedEvent processes a JobCreated event
func (s *Service) ProcessJobCreatedEvent(event JobCreatedEvent) error {
	s.logger.Info("Processing JobCreated event", "job_id", event.JobID)

	// Create job record
	job := &Job{
		ID:              event.JobID,
		RenterAddress:   event.RenterAddress,
		ProviderAddress: event.ProviderAddress,
		DockerImage:     event.DockerImage,
		InputFileCID:    event.InputFileCID,
		PaymentAmount:   event.PaymentAmount,
		Status:          JobStatusCreated,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	// Save job to database
	if err := s.db.Create(job).Error; err != nil {
		return fmt.Errorf("failed to create job: %w", err)
	}

	// Dispatch job to provider
	if err := s.dispatchJobToProvider(event); err != nil {
		return fmt.Errorf("failed to dispatch job: %w", err)
	}

	s.logger.Info("Job created and dispatched successfully", "job_id", event.JobID)
	return nil
}

// dispatchJobToProvider dispatches a job to a specific provider
func (s *Service) dispatchJobToProvider(event JobCreatedEvent) error {
	// Create job assignment
	assignment := JobAssignment{
		JobID:        event.JobID,
		DockerImage:  event.DockerImage,
		InputFileCID: event.InputFileCID,
	}

	// Publish to provider-specific subject
	subject := fmt.Sprintf("jobs.dispatch.%s", event.ProviderAddress)
	if err := s.natsClient.Publish(subject, assignment); err != nil {
		return fmt.Errorf("failed to publish job assignment: %w", err)
	}

	// Update job status to assigned
	now := time.Now()
	if err := s.db.Model(&Job{}).
		Where("id = ?", event.JobID).
		Updates(map[string]interface{}{
			"status":      JobStatusAssigned,
			"assigned_at": now,
			"updated_at":  now,
		}).Error; err != nil {
		return fmt.Errorf("failed to update job status: %w", err)
	}

	s.logger.Info("Job dispatched to provider", "job_id", event.JobID, "provider", event.ProviderAddress)
	return nil
}

// GetJobs retrieves jobs based on query criteria
func (s *Service) GetJobs(query JobQuery) ([]Job, error) {
	var jobs []Job

	db := s.db

	if query.RenterAddress != "" {
		db = db.Where("renter_address = ?", query.RenterAddress)
	}

	if query.ProviderAddress != "" {
		db = db.Where("provider_address = ?", query.ProviderAddress)
	}

	if query.Status != "" {
		db = db.Where("status = ?", query.Status)
	}

	// Set default limit if not specified
	limit := query.Limit
	if limit <= 0 {
		limit = 100
	}

	// Order by creation time descending
	db = db.Order("created_at DESC")

	if err := db.Offset(query.Offset).Limit(limit).Find(&jobs).Error; err != nil {
		return nil, fmt.Errorf("failed to get jobs: %w", err)
	}

	return jobs, nil
}

// GetJobByID retrieves a job by its ID
func (s *Service) GetJobByID(jobID string) (*Job, error) {
	var job Job
	if err := s.db.Where("id = ?", jobID).First(&job).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("job not found: %s", jobID)
		}
		return nil, fmt.Errorf("failed to get job: %w", err)
	}

	return &job, nil
}

// UpdateJobStatus updates the status of a job
func (s *Service) UpdateJobStatus(jobID string, status JobStatus, errorMessage string) error {
	updates := map[string]interface{}{
		"status":     status,
		"updated_at": time.Now(),
	}

	switch status {
	case JobStatusRunning:
		// No additional fields needed
	case JobStatusCompleted:
		now := time.Now()
		updates["completed_at"] = now
	case JobStatusFailed:
		now := time.Now()
		updates["failed_at"] = now
		updates["error_message"] = errorMessage
	case JobStatusCancelled:
		now := time.Now()
		updates["failed_at"] = now
		updates["error_message"] = errorMessage
	}

	if err := s.db.Model(&Job{}).
		Where("id = ?", jobID).
		Updates(updates).Error; err != nil {
		return fmt.Errorf("failed to update job status: %w", err)
	}

	s.logger.Info("Updated job status", "job_id", jobID, "status", status)
	return nil
}

// GetJobsByProvider retrieves all jobs for a specific provider
func (s *Service) GetJobsByProvider(providerAddress string, limit, offset int) ([]Job, error) {
	var jobs []Job

	if limit <= 0 {
		limit = 100
	}

	if err := s.db.Where("provider_address = ?", providerAddress).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&jobs).Error; err != nil {
		return nil, fmt.Errorf("failed to get jobs by provider: %w", err)
	}

	return jobs, nil
}

// GetJobsByRenter retrieves all jobs for a specific renter
func (s *Service) GetJobsByRenter(renterAddress string, limit, offset int) ([]Job, error) {
	var jobs []Job

	if limit <= 0 {
		limit = 100
	}

	if err := s.db.Where("renter_address = ?", renterAddress).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&jobs).Error; err != nil {
		return nil, fmt.Errorf("failed to get jobs by renter: %w", err)
	}

	return jobs, nil
}
