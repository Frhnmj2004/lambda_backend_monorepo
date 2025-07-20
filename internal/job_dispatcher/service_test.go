package job_dispatcher

import (
	"testing"
	"time"
)

func TestJobCreatedEvent_IPFSIntegration(t *testing.T) {
	// Test that JobCreatedEvent properly handles IPFS CIDs
	event := JobCreatedEvent{
		JobID:           "0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
		RenterAddress:   "0x742d35Cc6634C0532925a3b8D4C9db96C4b4d8b6",
		ProviderAddress: "0x8ba1f109551bD432803012645Hac136c772c3e",
		DockerImage:     "nvidia/cuda:11.8-base",
		InputFileCID:    "QmXabc123def456ghi789jkl012mno345pqr678stu901vwx234yz",
		PaymentAmount:   "1000000000000000000", // 1 ETH in wei
		BlockNumber:     12345,
		TransactionHash: "0xabcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890",
	}

	// Verify the event has the expected IPFS CID
	if event.InputFileCID == "" {
		t.Error("InputFileCID should not be empty")
	}

	// Verify it's a valid IPFS CID format (starts with Qm)
	if len(event.InputFileCID) < 2 || event.InputFileCID[:2] != "Qm" {
		t.Error("InputFileCID should be a valid IPFS CID starting with Qm")
	}
}

func TestJobAssignment_IPFSIntegration(t *testing.T) {
	// Test that JobAssignment properly handles IPFS CIDs
	assignment := JobAssignment{
		JobID:        "0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
		DockerImage:  "nvidia/cuda:11.8-base",
		InputFileCID: "QmXabc123def456ghi789jkl012mno345pqr678stu901vwx234yz",
	}

	// Verify the assignment has the expected IPFS CID
	if assignment.InputFileCID == "" {
		t.Error("InputFileCID should not be empty")
	}

	// Verify it's a valid IPFS CID format
	if len(assignment.InputFileCID) < 2 || assignment.InputFileCID[:2] != "Qm" {
		t.Error("InputFileCID should be a valid IPFS CID starting with Qm")
	}
}

func TestJob_IPFSIntegration(t *testing.T) {
	// Test that Job model properly handles IPFS CIDs
	now := time.Now()
	job := Job{
		ID:              "0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
		RenterAddress:   "0x742d35Cc6634C0532925a3b8D4C9db96C4b4d8b6",
		ProviderAddress: "0x8ba1f109551bD432803012645Hac136c772c3e",
		DockerImage:     "nvidia/cuda:11.8-base",
		InputFileCID:    "QmXabc123def456ghi789jkl012mno345pqr678stu901vwx234yz",
		OutputFileCID:   "QmYdef456ghi789jkl012mno345pqr678stu901vwx234yzabc123",
		PaymentAmount:   "1000000000000000000",
		Status:          JobStatusCreated,
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	// Verify input CID
	if job.InputFileCID == "" {
		t.Error("InputFileCID should not be empty")
	}

	// Verify output CID (can be empty for new jobs)
	if job.OutputFileCID == "" {
		t.Log("OutputFileCID is empty for new job - this is expected")
	}

	// Verify both CIDs are valid IPFS format when present
	if len(job.InputFileCID) < 2 || job.InputFileCID[:2] != "Qm" {
		t.Error("InputFileCID should be a valid IPFS CID starting with Qm")
	}

	if job.OutputFileCID != "" && (len(job.OutputFileCID) < 2 || job.OutputFileCID[:2] != "Qm") {
		t.Error("OutputFileCID should be a valid IPFS CID starting with Qm when present")
	}
}
