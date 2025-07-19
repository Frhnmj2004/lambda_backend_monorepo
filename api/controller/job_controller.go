package controller

import (
	"encoding/json"
	"strconv"
	"time"

	"lamda_backend/internal/job_dispatcher"
	"lamda_backend/pkg/logger"
	"lamda_backend/pkg/nats"

	"github.com/gofiber/fiber/v2"
)

// JobController handles HTTP requests for job operations
type JobController struct {
	natsClient *nats.NATSClient
	logger     *logger.Logger
}

// NewJobController creates a new job controller
func NewJobController(natsClient *nats.NATSClient, logger *logger.Logger) *JobController {
	return &JobController{
		natsClient: natsClient,
		logger:     logger.WithService("job-controller"),
	}
}

// GetJobs handles GET /api/v1/jobs
func (jc *JobController) GetJobs(c *fiber.Ctx) error {
	// Parse query parameters
	query := job_dispatcher.JobQuery{}

	// Parse renter_address
	if renterAddress := c.Query("renter_address"); renterAddress != "" {
		query.RenterAddress = renterAddress
	}

	// Parse provider_address
	if providerAddress := c.Query("provider_address"); providerAddress != "" {
		query.ProviderAddress = providerAddress
	}

	// Parse status
	if status := c.Query("status"); status != "" {
		query.Status = job_dispatcher.JobStatus(status)
	}

	// Parse limit
	if limitStr := c.Query("limit"); limitStr != "" {
		if limit, err := strconv.Atoi(limitStr); err == nil {
			query.Limit = limit
		}
	}

	// Parse offset
	if offsetStr := c.Query("offset"); offsetStr != "" {
		if offset, err := strconv.Atoi(offsetStr); err == nil {
			query.Offset = offset
		}
	}

	// Set default limit if not specified
	if query.Limit <= 0 {
		query.Limit = 100
	}

	// Query jobs via NATS
	queryData, err := json.Marshal(query)
	if err != nil {
		jc.logger.Error("Failed to marshal query", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to process query",
		})
	}

	responseData, err := jc.natsClient.PublishWithReply("jobs.query", queryData, 10*time.Second)
	if err != nil {
		jc.logger.Error("Failed to query jobs", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to query jobs",
		})
	}

	var response job_dispatcher.JobsResponse
	if err := json.Unmarshal(responseData, &response); err != nil {
		jc.logger.Error("Failed to unmarshal response", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to process response",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    response,
	})
}

// GetJobByID handles GET /api/v1/jobs/:id
func (jc *JobController) GetJobByID(c *fiber.Ctx) error {
	jobID := c.Params("id")
	if jobID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Job ID is required",
		})
	}

	// Create a query for a specific job
	query := job_dispatcher.JobQuery{
		Limit: 1,
	}

	// Query jobs via NATS
	queryData, err := json.Marshal(query)
	if err != nil {
		jc.logger.Error("Failed to marshal query", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to process query",
		})
	}

	responseData, err := jc.natsClient.PublishWithReply("jobs.query", queryData, 10*time.Second)
	if err != nil {
		jc.logger.Error("Failed to query jobs", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to query jobs",
		})
	}

	var response job_dispatcher.JobsResponse
	if err := json.Unmarshal(responseData, &response); err != nil {
		jc.logger.Error("Failed to unmarshal response", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to process response",
		})
	}

	// Find the specific job
	for _, job := range response.Jobs {
		if job.ID == jobID {
			return c.JSON(fiber.Map{
				"success": true,
				"data":    job,
			})
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "Job not found",
	})
}

// GetJobsByRenter handles GET /api/v1/jobs/renter/:address
func (jc *JobController) GetJobsByRenter(c *fiber.Ctx) error {
	renterAddress := c.Params("address")
	if renterAddress == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Renter address is required",
		})
	}

	// Parse limit and offset
	limit := 100
	offset := 0

	if limitStr := c.Query("limit"); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil {
			limit = l
		}
	}

	if offsetStr := c.Query("offset"); offsetStr != "" {
		if o, err := strconv.Atoi(offsetStr); err == nil {
			offset = o
		}
	}

	// Create query for renter's jobs
	query := job_dispatcher.JobQuery{
		RenterAddress: renterAddress,
		Limit:         limit,
		Offset:        offset,
	}

	// Query jobs via NATS
	queryData, err := json.Marshal(query)
	if err != nil {
		jc.logger.Error("Failed to marshal query", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to process query",
		})
	}

	responseData, err := jc.natsClient.PublishWithReply("jobs.query", queryData, 10*time.Second)
	if err != nil {
		jc.logger.Error("Failed to query jobs", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to query jobs",
		})
	}

	var response job_dispatcher.JobsResponse
	if err := json.Unmarshal(responseData, &response); err != nil {
		jc.logger.Error("Failed to unmarshal response", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to process response",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    response,
	})
}

// GetJobsByProvider handles GET /api/v1/jobs/provider/:address
func (jc *JobController) GetJobsByProvider(c *fiber.Ctx) error {
	providerAddress := c.Params("address")
	if providerAddress == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Provider address is required",
		})
	}

	// Parse limit and offset
	limit := 100
	offset := 0

	if limitStr := c.Query("limit"); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil {
			limit = l
		}
	}

	if offsetStr := c.Query("offset"); offsetStr != "" {
		if o, err := strconv.Atoi(offsetStr); err == nil {
			offset = o
		}
	}

	// Create query for provider's jobs
	query := job_dispatcher.JobQuery{
		ProviderAddress: providerAddress,
		Limit:           limit,
		Offset:          offset,
	}

	// Query jobs via NATS
	queryData, err := json.Marshal(query)
	if err != nil {
		jc.logger.Error("Failed to marshal query", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to process query",
		})
	}

	responseData, err := jc.natsClient.PublishWithReply("jobs.query", queryData, 10*time.Second)
	if err != nil {
		jc.logger.Error("Failed to query jobs", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to query jobs",
		})
	}

	var response job_dispatcher.JobsResponse
	if err := json.Unmarshal(responseData, &response); err != nil {
		jc.logger.Error("Failed to unmarshal response", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to process response",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    response,
	})
}

// GetJobStats handles GET /api/v1/jobs/stats
func (jc *JobController) GetJobStats(c *fiber.Ctx) error {
	// Query all jobs for stats
	query := job_dispatcher.JobQuery{
		Limit: 1000, // Get all jobs for stats
	}

	queryData, err := json.Marshal(query)
	if err != nil {
		jc.logger.Error("Failed to marshal query", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to process query",
		})
	}

	responseData, err := jc.natsClient.PublishWithReply("jobs.query", queryData, 10*time.Second)
	if err != nil {
		jc.logger.Error("Failed to query jobs", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to query jobs",
		})
	}

	var response job_dispatcher.JobsResponse
	if err := json.Unmarshal(responseData, &response); err != nil {
		jc.logger.Error("Failed to unmarshal response", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to process response",
		})
	}

	// Calculate stats
	stats := calculateJobStats(response.Jobs)

	return c.JSON(fiber.Map{
		"success": true,
		"data":    stats,
	})
}

// calculateJobStats calculates statistics from a list of jobs
func calculateJobStats(jobs []job_dispatcher.Job) fiber.Map {
	if len(jobs) == 0 {
		return fiber.Map{
			"total_jobs":       0,
			"completed_jobs":   0,
			"failed_jobs":      0,
			"running_jobs":     0,
			"pending_jobs":     0,
			"status_breakdown": map[string]int{},
		}
	}

	statusBreakdown := make(map[string]int)
	completedJobs := 0
	failedJobs := 0
	runningJobs := 0
	pendingJobs := 0

	for _, job := range jobs {
		status := string(job.Status)
		statusBreakdown[status]++

		switch job.Status {
		case job_dispatcher.JobStatusCompleted:
			completedJobs++
		case job_dispatcher.JobStatusFailed, job_dispatcher.JobStatusCancelled:
			failedJobs++
		case job_dispatcher.JobStatusRunning:
			runningJobs++
		case job_dispatcher.JobStatusCreated, job_dispatcher.JobStatusAssigned:
			pendingJobs++
		}
	}

	return fiber.Map{
		"total_jobs":       len(jobs),
		"completed_jobs":   completedJobs,
		"failed_jobs":      failedJobs,
		"running_jobs":     runningJobs,
		"pending_jobs":     pendingJobs,
		"status_breakdown": statusBreakdown,
	}
}
