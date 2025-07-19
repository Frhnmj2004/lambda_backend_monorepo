package controller

import (
	"encoding/json"
	"strconv"
	"time"

	"lamda_backend/internal/node_registry"
	"lamda_backend/pkg/logger"
	"lamda_backend/pkg/nats"

	"github.com/gofiber/fiber/v2"
)

// NodeController handles HTTP requests for node operations
type NodeController struct {
	natsClient *nats.NATSClient
	logger     *logger.Logger
}

// NewNodeController creates a new node controller
func NewNodeController(natsClient *nats.NATSClient, logger *logger.Logger) *NodeController {
	return &NodeController{
		natsClient: natsClient,
		logger:     logger.WithService("node-controller"),
	}
}

// GetActiveNodes handles GET /api/v1/nodes
func (nc *NodeController) GetActiveNodes(c *fiber.Ctx) error {
	// Parse query parameters
	query := node_registry.NodeQuery{}

	// Parse min_vram
	if minVRAMStr := c.Query("min_vram"); minVRAMStr != "" {
		if minVRAM, err := strconv.Atoi(minVRAMStr); err == nil {
			query.MinVRAM = &minVRAM
		}
	}

	// Parse gpu_model
	if gpuModel := c.Query("gpu_model"); gpuModel != "" {
		query.GPUModel = gpuModel
	}

	// Parse min_reputation_score
	if minReputationStr := c.Query("min_reputation_score"); minReputationStr != "" {
		if minReputation, err := strconv.Atoi(minReputationStr); err == nil {
			query.MinReputationScore = &minReputation
		}
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

	// Query nodes via NATS
	queryData, err := json.Marshal(query)
	if err != nil {
		nc.logger.Error("Failed to marshal query", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to process query",
		})
	}

	responseData, err := nc.natsClient.PublishWithReply("nodes.query", queryData, 10*time.Second)
	if err != nil {
		nc.logger.Error("Failed to query nodes", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to query nodes",
		})
	}

	var response node_registry.ActiveNodesResponse
	if err := json.Unmarshal(responseData, &response); err != nil {
		nc.logger.Error("Failed to unmarshal response", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to process response",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    response,
	})
}

// GetNodeByAddress handles GET /api/v1/nodes/:address
func (nc *NodeController) GetNodeByAddress(c *fiber.Ctx) error {
	address := c.Params("address")
	if address == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Node address is required",
		})
	}

	// Create a query for a specific node
	query := node_registry.NodeQuery{
		Limit: 1,
	}

	// Query nodes via NATS
	queryData, err := json.Marshal(query)
	if err != nil {
		nc.logger.Error("Failed to marshal query", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to process query",
		})
	}

	responseData, err := nc.natsClient.PublishWithReply("nodes.query", queryData, 10*time.Second)
	if err != nil {
		nc.logger.Error("Failed to query nodes", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to query nodes",
		})
	}

	var response node_registry.ActiveNodesResponse
	if err := json.Unmarshal(responseData, &response); err != nil {
		nc.logger.Error("Failed to unmarshal response", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to process response",
		})
	}

	// Find the specific node
	for _, node := range response.Nodes {
		if node.WalletAddress == address {
			return c.JSON(fiber.Map{
				"success": true,
				"data":    node,
			})
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "Node not found",
	})
}

// GetNodeStats handles GET /api/v1/nodes/stats
func (nc *NodeController) GetNodeStats(c *fiber.Ctx) error {
	// Query all active nodes
	query := node_registry.NodeQuery{
		Limit: 1000, // Get all nodes for stats
	}

	queryData, err := json.Marshal(query)
	if err != nil {
		nc.logger.Error("Failed to marshal query", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to process query",
		})
	}

	responseData, err := nc.natsClient.PublishWithReply("nodes.query", queryData, 10*time.Second)
	if err != nil {
		nc.logger.Error("Failed to query nodes", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to query nodes",
		})
	}

	var response node_registry.ActiveNodesResponse
	if err := json.Unmarshal(responseData, &response); err != nil {
		nc.logger.Error("Failed to unmarshal response", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to process response",
		})
	}

	// Calculate stats
	stats := calculateNodeStats(response.Nodes)

	return c.JSON(fiber.Map{
		"success": true,
		"data":    stats,
	})
}

// calculateNodeStats calculates statistics from a list of nodes
func calculateNodeStats(nodes []node_registry.Provider) fiber.Map {
	if len(nodes) == 0 {
		return fiber.Map{
			"total_nodes":    0,
			"online_nodes":   0,
			"total_vram":     0,
			"avg_reputation": 0,
			"gpu_models":     []string{},
			"total_jobs":     0,
		}
	}

	totalVRAM := 0
	totalReputation := 0
	totalJobs := 0
	onlineNodes := 0
	gpuModels := make(map[string]int)

	for _, node := range nodes {
		totalVRAM += node.VRAM
		totalReputation += node.ReputationScore
		totalJobs += node.TotalJobsCompleted
		gpuModels[node.GPUModel]++

		if node.IsOnline {
			onlineNodes++
		}
	}

	// Convert gpu models map to slice
	gpuModelList := make([]string, 0, len(gpuModels))
	for model := range gpuModels {
		gpuModelList = append(gpuModelList, model)
	}

	return fiber.Map{
		"total_nodes":      len(nodes),
		"online_nodes":     onlineNodes,
		"total_vram":       totalVRAM,
		"avg_reputation":   float64(totalReputation) / float64(len(nodes)),
		"gpu_models":       gpuModelList,
		"total_jobs":       totalJobs,
		"gpu_model_counts": gpuModels,
	}
}
