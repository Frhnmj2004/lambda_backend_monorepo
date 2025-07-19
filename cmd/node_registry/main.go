package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"lamda_backend/config"
	"lamda_backend/internal/node_registry"
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
	log := logger.New("info").WithService("node-registry")
	log.Info("Starting Lamda Node Registry")

	// Connect to database
	db, err := database.NewPostgresConnection(cfg.DatabaseURL, "info")
	if err != nil {
		log.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}

	// Auto-migrate database
	if err := database.AutoMigrate(db, &node_registry.Provider{}); err != nil {
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

	// Connect to blockchain (opBNB for NodeReputation contract)
	blockchainClient, err := blockchain.NewEVMClient(cfg.OpBNBRPCURL)
	if err != nil {
		log.Error("Failed to connect to blockchain", "error", err)
		os.Exit(1)
	}
	defer blockchainClient.Close()

	// Wait for blockchain connection
	if err := blockchainClient.WaitForConnection(ctx, 30); err != nil {
		log.Error("Failed to wait for blockchain connection", "error", err)
		os.Exit(1)
	}

	// Initialize node registry service
	nodeRegistryService := node_registry.NewService(db, natsClient, blockchainClient, log, cfg.NodeReputationContractAddress)

	// Start the service
	if err := nodeRegistryService.Start(context.Background()); err != nil {
		log.Error("Failed to start node registry service", "error", err)
		os.Exit(1)
	}

	log.Info("Node Registry service started successfully")

	// Start background task to mark offline providers
	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				if err := nodeRegistryService.MarkOfflineProviders(); err != nil {
					log.Error("Failed to mark offline providers", "error", err)
				}
			}
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down Node Registry...")
	log.Info("Node Registry stopped")
}
