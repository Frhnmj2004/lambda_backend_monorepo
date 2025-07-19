package node_registry

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

// Service handles node registry operations
type Service struct {
	db           *gorm.DB
	natsClient   *nats.NATSClient
	blockchain   *blockchain.EVMClient
	logger       *logger.Logger
	contractAddr string
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
	s.logger.Info("Starting blockchain event listener")

	// TODO: Implement actual blockchain event listening
	// This would involve:
	// 1. Setting up event filters for NodeRegistered and NodeHeartbeat events
	// 2. Polling for new events
	// 3. Processing events and updating the database

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
