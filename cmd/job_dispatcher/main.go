package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"lamda_backend/config"
	"lamda_backend/internal/job_dispatcher"
	"lamda_backend/pkg/blockchain"
	"lamda_backend/pkg/database"
	"lamda_backend/pkg/logger"
	"lamda_backend/pkg/nats"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		os.Exit(1)
	}

	// Initialize logger
	log := logger.New("info").WithService("job-dispatcher")
	log.Info("Starting Lamda Job Dispatcher")

	// Connect to database
	db, err := database.NewPostgresConnection(cfg.DatabaseURL, "info")
	if err != nil {
		log.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}

	// Auto-migrate database
	if err := database.AutoMigrate(db, &job_dispatcher.Job{}); err != nil {
		log.Error("Failed to migrate database", "error", err)
		os.Exit(1)
	}

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

	// Connect to blockchain (BSC for JobManager contract)
	blockchainClient, err := blockchain.NewEVMClient(cfg.BSCRPCURL)
	if err != nil {
		log.Error("Failed to connect to blockchain", "error", err)
		os.Exit(1)
	}
	defer blockchainClient.Close()

	// Wait for blockchain connection
	if err := blockchainClient.WaitForConnection(ctx, 30*time.Second); err != nil {
		log.Error("Failed to wait for blockchain connection", "error", err)
		os.Exit(1)
	}

	// Initialize job dispatcher service
	jobDispatcherService := job_dispatcher.NewService(db, natsClient, blockchainClient, log, cfg.JobManagerContractAddress)

	// Start the service
	if err := jobDispatcherService.Start(context.Background()); err != nil {
		log.Error("Failed to start job dispatcher service", "error", err)
		os.Exit(1)
	}

	log.Info("Job Dispatcher service started successfully")

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down Job Dispatcher...")
	log.Info("Job Dispatcher stopped")
}
