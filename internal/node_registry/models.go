package node_registry

import (
	"time"

	"gorm.io/gorm"
)

// Provider represents a GPU provider in the Lamda network
type Provider struct {
	ID                 uint      `json:"id" gorm:"primaryKey"`
	WalletAddress      string    `json:"wallet_address" gorm:"uniqueIndex;not null"`
	GPUModel           string    `json:"gpu_model" gorm:"not null"`
	VRAM               int       `json:"vram" gorm:"not null"`
	LastSeen           time.Time `json:"last_seen" gorm:"not null"`
	IsOnline           bool      `json:"is_online" gorm:"default:true"`
	TotalJobsCompleted int       `json:"total_jobs_completed" gorm:"default:0"`
	ReputationScore    int       `json:"reputation_score" gorm:"default:0"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

// TableName specifies the table name for the Provider model
func (Provider) TableName() string {
	return "providers"
}

// BeforeUpdate is a GORM hook that updates the UpdatedAt field
func (p *Provider) BeforeUpdate(tx *gorm.DB) error {
	p.UpdatedAt = time.Now()
	return nil
}

// NodeRegisteredEvent represents the NodeRegistered event from the smart contract
type NodeRegisteredEvent struct {
	ProviderAddress string `json:"provider_address"`
	GPUModel        string `json:"gpu_model"`
	VRAM            int    `json:"vram"`
	BlockNumber     uint64 `json:"block_number"`
	TransactionHash string `json:"transaction_hash"`
}

// NodeHeartbeatEvent represents the NodeHeartbeat event from the smart contract
type NodeHeartbeatEvent struct {
	ProviderAddress string `json:"provider_address"`
	BlockNumber     uint64 `json:"block_number"`
	TransactionHash string `json:"transaction_hash"`
}

// ActiveNodesResponse represents the response for active nodes query
type ActiveNodesResponse struct {
	Nodes []Provider `json:"nodes"`
	Count int        `json:"count"`
}

// NodeQuery represents a query for nodes
type NodeQuery struct {
	MinVRAM            *int   `json:"min_vram,omitempty"`
	GPUModel           string `json:"gpu_model,omitempty"`
	MinReputationScore *int   `json:"min_reputation_score,omitempty"`
	Limit              int    `json:"limit,omitempty"`
	Offset             int    `json:"offset,omitempty"`
}
