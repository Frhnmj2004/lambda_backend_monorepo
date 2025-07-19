package job_dispatcher

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"lamda_backend/pkg/blockchain"
	"lamda_backend/pkg/logger"
	"lamda_backend/pkg/nats"

	"gorm.io/gorm"
)

// Service handles job dispatching operations
type Service struct {
	db           *gorm.DB
	natsClient   *nats.NATSClient
	blockchain   *blockchain.EVMClient
	logger       *logger.Logger
	contractAddr string
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
	s.logger.Info("Starting blockchain event listener")

	// TODO: Implement actual blockchain event listening
	// This would involve:
	// 1. Setting up event filters for JobCreated events
	// 2. Polling for new events
	// 3. Processing events and dispatching jobs

	// For now, we'll simulate the event processing
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			s.logger.Info("Stopping blockchain event listener")
			return
		case <-ticker.C:
			// Simulate processing events
			s.logger.Debug("Checking for blockchain events")
		}
	}
}

// ProcessJobCreatedEvent processes a JobCreated event
func (s *Service) ProcessJobCreatedEvent(event JobCreatedEvent) error {
	s.logger.Info("Processing JobCreated event", "job_id", event.JobID)

	// Create job record
	job := &Job{
		ID:                     event.JobID,
		RenterAddress:          event.RenterAddress,
		ProviderAddress:        event.ProviderAddress,
		DockerImage:            event.DockerImage,
		GreenfieldInputURL:     event.GreenfieldInputURL,
		GreenfieldOutputBucket: event.GreenfieldOutputBucket,
		GreenfieldOutputName:   event.GreenfieldOutputName,
		PaymentAmount:          event.PaymentAmount,
		Status:                 JobStatusCreated,
		CreatedAt:              time.Now(),
		UpdatedAt:              time.Now(),
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
		JobID:                  event.JobID,
		DockerImage:            event.DockerImage,
		GreenfieldInputURL:     event.GreenfieldInputURL,
		GreenfieldOutputBucket: event.GreenfieldOutputBucket,
		GreenfieldOutputName:   event.GreenfieldOutputName,
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
