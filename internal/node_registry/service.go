package node_registry

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

// Service handles node registry operations
type Service struct {
	db                     *gorm.DB
	natsClient             *nats.NATSClient
	blockchain             *blockchain.EVMClient
	logger                 *logger.Logger
	contractAddr           string
	nodeReputationContract *contracts.NodeReputation
}

// NewService creates a new node registry service
func NewService(db *gorm.DB, natsClient *nats.NATSClient, blockchain *blockchain.EVMClient, logger *logger.Logger, contractAddr string) *Service {
	return &Service{
		db:           db,
		natsClient:   natsClient,
		blockchain:   blockchain,
		logger:       logger.WithService("node-registry"),
		contractAddr: contractAddr,
	}
}

// Start starts the node registry service
func (s *Service) Start(ctx context.Context) error {
	s.logger.Info("Starting node registry service")

	// Initialize the NodeReputation contract
	contractAddress := common.HexToAddress(s.contractAddr)
	nodeReputationContract, err := contracts.NewNodeReputation(contractAddress, s.blockchain.GetClient())
	if err != nil {
		return fmt.Errorf("failed to initialize NodeReputation contract: %w", err)
	}
	s.nodeReputationContract = nodeReputationContract

	// Subscribe to NATS queries
	if err := s.subscribeToQueries(); err != nil {
		return fmt.Errorf("failed to subscribe to queries: %w", err)
	}

	// Start blockchain event listener
	go s.listenToBlockchainEvents(ctx)

	s.logger.Info("Node registry service started successfully")
	return nil
}

// subscribeToQueries subscribes to NATS queries for node information
func (s *Service) subscribeToQueries() error {
	// Subscribe to nodes.query subject
	_, err := s.natsClient.SubscribeWithReply("nodes.query", s.handleNodeQuery)
	if err != nil {
		return fmt.Errorf("failed to subscribe to nodes.query: %w", err)
	}

	s.logger.Info("Subscribed to nodes.query")
	return nil
}

// handleNodeQuery handles queries for active nodes
func (s *Service) handleNodeQuery(data []byte) ([]byte, error) {
	var query NodeQuery
	if err := json.Unmarshal(data, &query); err != nil {
		return nil, fmt.Errorf("failed to unmarshal query: %w", err)
	}

	nodes, err := s.GetActiveNodes(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get active nodes: %w", err)
	}

	response := ActiveNodesResponse{
		Nodes: nodes,
		Count: len(nodes),
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

	// Also poll for offline providers every 5 minutes
	offlineTicker := time.NewTicker(5 * time.Minute)
	defer offlineTicker.Stop()

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
		case <-offlineTicker.C:
			if err := s.MarkOfflineProviders(); err != nil {
				s.logger.Error("Failed to mark offline providers", "error", err)
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

	// Query for NodeRegistered events
	nodeRegisteredEvents, err := s.nodeReputationContract.FilterNodeRegistered(&bind.FilterOpts{
		Start:   fromBlock,
		End:     &currentBlock,
		Context: ctx,
	}, nil)
	if err != nil {
		return fmt.Errorf("failed to filter NodeRegistered events: %w", err)
	}

	// Process NodeRegistered events
	for nodeRegisteredEvents.Next() {
		event := nodeRegisteredEvents.Event
		s.logger.Info("Received NodeRegistered event", "provider", event.Provider.Hex())
		if err := s.processNodeRegisteredEvent(event); err != nil {
			s.logger.Error("Failed to process NodeRegistered event", "error", err, "provider", event.Provider.Hex())
		}
	}

	// Query for NodeHeartbeat events
	nodeHeartbeatEvents, err := s.nodeReputationContract.FilterNodeHeartbeat(&bind.FilterOpts{
		Start:   fromBlock,
		End:     &currentBlock,
		Context: ctx,
	}, nil)
	if err != nil {
		return fmt.Errorf("failed to filter NodeHeartbeat events: %w", err)
	}

	// Process NodeHeartbeat events
	for nodeHeartbeatEvents.Next() {
		event := nodeHeartbeatEvents.Event
		s.logger.Debug("Received NodeHeartbeat event", "provider", event.Provider.Hex())
		if err := s.processNodeHeartbeatEvent(event); err != nil {
			s.logger.Error("Failed to process NodeHeartbeat event", "error", err, "provider", event.Provider.Hex())
		}
	}

	return nil
}

// processNodeRegisteredEvent processes a NodeRegistered event from the blockchain
func (s *Service) processNodeRegisteredEvent(event *contracts.NodeReputationNodeRegistered) error {
	s.logger.Info("Processing NodeRegistered event", "provider", event.Provider.Hex())

	// Convert the event to our internal format
	nodeEvent := NodeRegisteredEvent{
		ProviderAddress: event.Provider.Hex(),
		GPUModel:        event.GpuModel,
		VRAM:            int(event.Vram.Int64()),
		BlockNumber:     event.Raw.BlockNumber,
		TransactionHash: event.Raw.TxHash.Hex(),
	}

	// Process the event using existing logic
	return s.ProcessNodeRegisteredEvent(nodeEvent)
}

// processNodeHeartbeatEvent processes a NodeHeartbeat event from the blockchain
func (s *Service) processNodeHeartbeatEvent(event *contracts.NodeReputationNodeHeartbeat) error {
	s.logger.Debug("Processing NodeHeartbeat event", "provider", event.Provider.Hex())

	// Convert the event to our internal format
	heartbeatEvent := NodeHeartbeatEvent{
		ProviderAddress: event.Provider.Hex(),
		BlockNumber:     event.Raw.BlockNumber,
		TransactionHash: event.Raw.TxHash.Hex(),
	}

	// Process the event using existing logic
	return s.ProcessNodeHeartbeatEvent(heartbeatEvent)
}

// ProcessNodeRegisteredEvent processes a NodeRegistered event
func (s *Service) ProcessNodeRegisteredEvent(event NodeRegisteredEvent) error {
	s.logger.Info("Processing NodeRegistered event", "provider", event.ProviderAddress)

	provider := &Provider{
		WalletAddress: event.ProviderAddress,
		GPUModel:      event.GPUModel,
		VRAM:          event.VRAM,
		LastSeen:      time.Now(),
		IsOnline:      true,
	}

	// Upsert the provider
	result := s.db.Where(Provider{WalletAddress: event.ProviderAddress}).
		Assign(provider).
		FirstOrCreate(provider)

	if result.Error != nil {
		return fmt.Errorf("failed to upsert provider: %w", result.Error)
	}

	s.logger.Info("Provider registered successfully", "provider", event.ProviderAddress)
	return nil
}

// ProcessNodeHeartbeatEvent processes a NodeHeartbeat event
func (s *Service) ProcessNodeHeartbeatEvent(event NodeHeartbeatEvent) error {
	s.logger.Debug("Processing NodeHeartbeat event", "provider", event.ProviderAddress)

	// Update the provider's last seen time
	result := s.db.Model(&Provider{}).
		Where("wallet_address = ?", event.ProviderAddress).
		Updates(map[string]interface{}{
			"last_seen": time.Now(),
			"is_online": true,
		})

	if result.Error != nil {
		return fmt.Errorf("failed to update provider heartbeat: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		s.logger.Warn("Provider not found for heartbeat", "provider", event.ProviderAddress)
	}

	return nil
}

// GetActiveNodes retrieves active nodes based on query criteria
func (s *Service) GetActiveNodes(query NodeQuery) ([]Provider, error) {
	var providers []Provider

	db := s.db.Where("is_online = ?", true)

	if query.MinVRAM != nil {
		db = db.Where("vram >= ?", *query.MinVRAM)
	}

	if query.GPUModel != "" {
		db = db.Where("gpu_model = ?", query.GPUModel)
	}

	if query.MinReputationScore != nil {
		db = db.Where("reputation_score >= ?", *query.MinReputationScore)
	}

	// Set default limit if not specified
	limit := query.Limit
	if limit <= 0 {
		limit = 100
	}

	// Order by reputation score descending
	db = db.Order("reputation_score DESC")

	if err := db.Offset(query.Offset).Limit(limit).Find(&providers).Error; err != nil {
		return nil, fmt.Errorf("failed to get active nodes: %w", err)
	}

	return providers, nil
}

// MarkOfflineProviders marks providers as offline if they haven't sent a heartbeat recently
func (s *Service) MarkOfflineProviders() error {
	// Mark providers as offline if they haven't been seen in the last 5 minutes
	threshold := time.Now().Add(-5 * time.Minute)

	result := s.db.Model(&Provider{}).
		Where("is_online = ? AND last_seen < ?", true, threshold).
		Update("is_online", false)

	if result.Error != nil {
		return fmt.Errorf("failed to mark offline providers: %w", result.Error)
	}

	if result.RowsAffected > 0 {
		s.logger.Info("Marked providers as offline", "count", result.RowsAffected)
	}

	return nil
}

// UpdateReputationScore updates a provider's reputation score
func (s *Service) UpdateReputationScore(walletAddress string, score int) error {
	result := s.db.Model(&Provider{}).
		Where("wallet_address = ?", walletAddress).
		Update("reputation_score", score)

	if result.Error != nil {
		return fmt.Errorf("failed to update reputation score: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("provider not found: %s", walletAddress)
	}

	s.logger.Info("Updated reputation score", "provider", walletAddress, "score", score)
	return nil
}

// IncrementJobsCompleted increments the total jobs completed for a provider
func (s *Service) IncrementJobsCompleted(walletAddress string) error {
	result := s.db.Model(&Provider{}).
		Where("wallet_address = ?", walletAddress).
		UpdateColumn("total_jobs_completed", gorm.Expr("total_jobs_completed + ?", 1))

	if result.Error != nil {
		return fmt.Errorf("failed to increment jobs completed: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("provider not found: %s", walletAddress)
	}

	s.logger.Info("Incremented jobs completed", "provider", walletAddress)
	return nil
}
