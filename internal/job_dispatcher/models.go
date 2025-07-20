package job_dispatcher

import (
	"time"
)

// JobCreatedEvent represents the JobCreated event from the JobManager contract
type JobCreatedEvent struct {
	JobID           string `json:"job_id"`
	RenterAddress   string `json:"renter_address"`
	ProviderAddress string `json:"provider_address"`
	DockerImage     string `json:"docker_image"`
	InputFileCID    string `json:"input_file_cid"`
	PaymentAmount   string `json:"payment_amount"`
	BlockNumber     uint64 `json:"block_number"`
	TransactionHash string `json:"transaction_hash"`
}

// JobAssignment represents a job assignment sent to a provider via NATS
type JobAssignment struct {
	JobID        string `json:"jobId"`
	DockerImage  string `json:"dockerImage"`
	InputFileCID string `json:"inputFileCID"`
}

// JobStatus represents the status of a job
type JobStatus string

const (
	JobStatusCreated   JobStatus = "created"
	JobStatusAssigned  JobStatus = "assigned"
	JobStatusRunning   JobStatus = "running"
	JobStatusCompleted JobStatus = "completed"
	JobStatusFailed    JobStatus = "failed"
	JobStatusCancelled JobStatus = "cancelled"
)

// Job represents a job in the system
type Job struct {
	ID              string     `json:"id"`
	RenterAddress   string     `json:"renter_address"`
	ProviderAddress string     `json:"provider_address"`
	DockerImage     string     `json:"docker_image"`
	InputFileCID    string     `json:"input_file_cid"`
	OutputFileCID   string     `json:"output_file_cid,omitempty"`
	PaymentAmount   string     `json:"payment_amount"`
	Status          JobStatus  `json:"status"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	AssignedAt      *time.Time `json:"assigned_at,omitempty"`
	CompletedAt     *time.Time `json:"completed_at,omitempty"`
	FailedAt        *time.Time `json:"failed_at,omitempty"`
	ErrorMessage    string     `json:"error_message,omitempty"`
}

// JobQuery represents a query for jobs
type JobQuery struct {
	RenterAddress   string    `json:"renter_address,omitempty"`
	ProviderAddress string    `json:"provider_address,omitempty"`
	Status          JobStatus `json:"status,omitempty"`
	Limit           int       `json:"limit,omitempty"`
	Offset          int       `json:"offset,omitempty"`
}

// JobsResponse represents the response for jobs query
type JobsResponse struct {
	Jobs  []Job `json:"jobs"`
	Count int   `json:"count"`
}
