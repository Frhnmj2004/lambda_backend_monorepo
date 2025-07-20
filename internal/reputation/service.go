package reputation

import (
	"context"
	"fmt"
	"time"

	"lamda_backend/pkg/blockchain"
	"lamda_backend/pkg/contracts"
	"lamda_backend/pkg/logger"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// JobConfirmedEvent represents the JobConfirmed event from the JobManager contract
type JobConfirmedEvent struct {
	JobID           string `json:"job_id"`
	ProviderAddress string `json:"provider_address"`
	RenterAddress   string `json:"renter_address"`
	BlockNumber     uint64 `json:"block_number"`
	TransactionHash string `json:"transaction_hash"`
}

// Service handles reputation operations
type Service struct {
	bscClient                  *blockchain.EVMClient
	opBNBClient                *blockchain.EVMClient
	logger                     *logger.Logger
	jobManagerContractAddr     string
	nodeReputationContractAddr string
	adminKey                   string
	jobManagerContract         *contracts.JobManager
	nodeReputationContract     *contracts.NodeReputation
}

// NewService creates a new reputation service
func NewService(bscClient, opBNBClient *blockchain.EVMClient, logger *logger.Logger, jobManagerContractAddr, nodeReputationContractAddr, adminKey string) *Service {
	return &Service{
		bscClient:                  bscClient,
		opBNBClient:                opBNBClient,
		logger:                     logger.WithService("reputation"),
		jobManagerContractAddr:     jobManagerContractAddr,
		nodeReputationContractAddr: nodeReputationContractAddr,
		adminKey:                   adminKey,
	}
}

// Start starts the reputation service
func (s *Service) Start(ctx context.Context) error {
	s.logger.Info("Starting reputation service")

	// Initialize the JobManager contract (BSC)
	jobManagerAddress := common.HexToAddress(s.jobManagerContractAddr)
	jobManagerContract, err := contracts.NewJobManager(jobManagerAddress, s.bscClient.GetClient())
	if err != nil {
		return fmt.Errorf("failed to initialize JobManager contract: %w", err)
	}
	s.jobManagerContract = jobManagerContract

	// Initialize the NodeReputation contract (opBNB)
	nodeReputationAddress := common.HexToAddress(s.nodeReputationContractAddr)
	nodeReputationContract, err := contracts.NewNodeReputation(nodeReputationAddress, s.opBNBClient.GetClient())
	if err != nil {
		return fmt.Errorf("failed to initialize NodeReputation contract: %w", err)
	}
	s.nodeReputationContract = nodeReputationContract

	// Start blockchain event listener
	go s.listenToBlockchainEvents(ctx)

	s.logger.Info("Reputation service started successfully")
	return nil
}

// listenToBlockchainEvents listens for blockchain events
func (s *Service) listenToBlockchainEvents(ctx context.Context) {
	s.logger.Info("Starting blockchain event listener (polling mode)")

	// Get the latest block number to start listening from
	latestBlock, err := s.bscClient.GetLatestBlockNumber(ctx)
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
				if currentBlock, err := s.bscClient.GetLatestBlockNumber(ctx); err == nil {
					fromBlock = uint64(currentBlock)
				}
			}
		}
	}
}

// pollForEvents polls for blockchain events
func (s *Service) pollForEvents(ctx context.Context, fromBlock uint64) error {
	// Get current block
	currentBlock, err := s.bscClient.GetLatestBlockNumber(ctx)
	if err != nil {
		return fmt.Errorf("failed to get current block: %w", err)
	}

	// Only process if there are new blocks
	if uint64(currentBlock) <= fromBlock {
		return nil
	}

	// Query for JobConfirmed events
	jobConfirmedEvents, err := s.jobManagerContract.FilterJobConfirmed(&bind.FilterOpts{
		Start:   fromBlock,
		End:     &currentBlock,
		Context: ctx,
	}, nil)
	if err != nil {
		return fmt.Errorf("failed to filter JobConfirmed events: %w", err)
	}

	// Process JobConfirmed events
	for jobConfirmedEvents.Next() {
		event := jobConfirmedEvents.Event
		s.logger.Info("Received JobConfirmed event", "job_id", fmt.Sprintf("0x%x", event.JobId))
		if err := s.processJobConfirmedEvent(event); err != nil {
			s.logger.Error("Failed to process JobConfirmed event", "error", err, "job_id", fmt.Sprintf("0x%x", event.JobId))
		}
	}

	return nil
}

// processJobConfirmedEvent processes a JobConfirmed event from the blockchain
func (s *Service) processJobConfirmedEvent(event *contracts.JobManagerJobConfirmed) error {
	s.logger.Info("Processing JobConfirmed event", "job_id", fmt.Sprintf("0x%x", event.JobId))

	// Get the job details to find the provider address
	jobInfo, err := s.jobManagerContract.GetJobInfo(&bind.CallOpts{}, event.JobId)
	if err != nil {
		return fmt.Errorf("failed to get job info: %w", err)
	}

	// Convert the event to our internal format
	confirmedEvent := JobConfirmedEvent{
		JobID:           fmt.Sprintf("0x%x", event.JobId),
		ProviderAddress: jobInfo.Provider.Hex(),
		RenterAddress:   jobInfo.Renter.Hex(),
		BlockNumber:     event.Raw.BlockNumber,
		TransactionHash: event.Raw.TxHash.Hex(),
	}

	// Process the event using existing logic
	return s.ProcessJobConfirmedEvent(confirmedEvent)
}

// ProcessJobConfirmedEvent processes a JobConfirmed event
func (s *Service) ProcessJobConfirmedEvent(event JobConfirmedEvent) error {
	s.logger.Info("Processing JobConfirmed event", "job_id", event.JobID, "provider", event.ProviderAddress)

	// Call the incrementJobs function on the NodeReputation contract
	if err := s.incrementJobsOnChain(event.ProviderAddress); err != nil {
		return fmt.Errorf("failed to increment jobs on chain: %w", err)
	}

	s.logger.Info("Job confirmed and reputation updated", "job_id", event.JobID, "provider", event.ProviderAddress)
	return nil
}

// incrementJobsOnChain calls the incrementJobs function on the NodeReputation contract
func (s *Service) incrementJobsOnChain(providerAddress string) error {
	s.logger.Info("Incrementing jobs on chain for provider", "provider", providerAddress)

	// Create transaction options
	auth, err := s.opBNBClient.CreateTransactOpts(context.Background(), s.adminKey)
	if err != nil {
		return fmt.Errorf("failed to create transaction options: %w", err)
	}

	// Convert provider address to common.Address
	providerAddr := common.HexToAddress(providerAddress)

	// Call the incrementJobs function
	tx, err := s.nodeReputationContract.IncrementJobs(auth, providerAddr)
	if err != nil {
		return fmt.Errorf("failed to call incrementJobs: %w", err)
	}

	s.logger.Info("IncrementJobs transaction sent", "tx_hash", tx.Hash().Hex(), "provider", providerAddress)

	// Wait for transaction to be mined
	receipt, err := s.opBNBClient.WaitForTransaction(context.Background(), tx)
	if err != nil {
		return fmt.Errorf("failed to wait for transaction: %w", err)
	}

	if receipt.Status == 0 {
		return fmt.Errorf("transaction failed: %s", tx.Hash().Hex())
	}

	s.logger.Info("IncrementJobs transaction confirmed", "tx_hash", tx.Hash().Hex(), "block", receipt.BlockNumber)
	return nil
}

// GetReputationScore retrieves the reputation score for a provider
func (s *Service) GetReputationScore(providerAddress string) (int, error) {
	s.logger.Info("Getting reputation score for provider", "provider", providerAddress)

	// Convert provider address to common.Address
	providerAddr := common.HexToAddress(providerAddress)

	// Call the getProviderInfo function
	providerInfo, err := s.nodeReputationContract.GetProviderInfo(&bind.CallOpts{}, providerAddr)
	if err != nil {
		return 0, fmt.Errorf("failed to get provider info: %w", err)
	}

	// The job count serves as the reputation score
	reputationScore := int(providerInfo.JobCount.Int64())
	s.logger.Info("Retrieved reputation score", "provider", providerAddress, "score", reputationScore)
	return reputationScore, nil
}

// GetJobsCompleted retrieves the number of jobs completed by a provider
func (s *Service) GetJobsCompleted(providerAddress string) (int, error) {
	s.logger.Info("Getting jobs completed for provider", "provider", providerAddress)

	// Convert provider address to common.Address
	providerAddr := common.HexToAddress(providerAddress)

	// Call the getProviderInfo function
	providerInfo, err := s.nodeReputationContract.GetProviderInfo(&bind.CallOpts{}, providerAddr)
	if err != nil {
		return 0, fmt.Errorf("failed to get provider info: %w", err)
	}

	// Return the job count
	jobsCompleted := int(providerInfo.JobCount.Int64())
	s.logger.Info("Retrieved jobs completed", "provider", providerAddress, "count", jobsCompleted)
	return jobsCompleted, nil
}
