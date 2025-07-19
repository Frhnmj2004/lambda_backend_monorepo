package router

import (
	"lamda_backend/api/controller"
	"lamda_backend/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberlogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// SetupRoutes configures all API routes
func SetupRoutes(app *fiber.App, nodeController *controller.NodeController, jobController *controller.JobController, log *logger.Logger) {
	// Middleware
	app.Use(recover.New())
	app.Use(fiberlogger.New(fiberlogger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "healthy",
			"service": "lamda-backend",
		})
	})

	// API v1 routes
	api := app.Group("/api/v1")

	// Node routes
	nodes := api.Group("/nodes")
	nodes.Get("/", nodeController.GetActiveNodes)
	nodes.Get("/stats", nodeController.GetNodeStats)
	nodes.Get("/:address", nodeController.GetNodeByAddress)

	// Job routes
	jobs := api.Group("/jobs")
	jobs.Get("/", jobController.GetJobs)
	jobs.Get("/stats", jobController.GetJobStats)
	jobs.Get("/:id", jobController.GetJobByID)
	jobs.Get("/renter/:address", jobController.GetJobsByRenter)
	jobs.Get("/provider/:address", jobController.GetJobsByProvider)

	// 404 handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Route not found",
		})
	})
}
