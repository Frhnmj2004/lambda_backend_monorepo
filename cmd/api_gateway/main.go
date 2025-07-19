package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"lamda_backend/api/controller"
	"lamda_backend/api/router"
	"lamda_backend/config"
	"lamda_backend/pkg/logger"
	"lamda_backend/pkg/nats"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		os.Exit(1)
	}

	// Initialize logger
	log := logger.New("info").WithService("api-gateway")
	log.Info("Starting Lamda API Gateway")

	// Connect to NATS
	natsClient, err := nats.NewNATSConnection(cfg.NATSURL)
	if err != nil {
		log.Error("Failed to connect to NATS", "error", err)
		os.Exit(1)
	}
	defer natsClient.Close()

	// Wait for NATS connection
	ctx, cancel := context.WithTimeout(context.Background(), 30)
	defer cancel()
	if err := natsClient.WaitForConnection(ctx, 30); err != nil {
		log.Error("Failed to wait for NATS connection", "error", err)
		os.Exit(1)
	}

	// Initialize controllers
	nodeController := controller.NewNodeController(natsClient, log)
	jobController := controller.NewJobController(natsClient, log)

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName:      "Lamda API Gateway",
		ServerHeader: "Lamda",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			log.Error("Request error", "error", err, "path", c.Path())
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal server error",
			})
		},
	})

	// Setup routes
	router.SetupRoutes(app, nodeController, jobController, log)

	// Start server in a goroutine
	go func() {
		addr := fmt.Sprintf(":%s", cfg.APIPort)
		log.Info("Starting server", "address", addr)
		if err := app.Listen(addr); err != nil {
			log.Error("Failed to start server", "error", err)
			os.Exit(1)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down API Gateway...")

	// Graceful shutdown
	if err := app.Shutdown(); err != nil {
		log.Error("Error during shutdown", "error", err)
	}

	log.Info("API Gateway stopped")
}
