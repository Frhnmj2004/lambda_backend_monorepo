package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"lamda_backend/config"
	"lamda_backend/internal/reputation"
	"lamda_backend/pkg/blockchain"
	"lamda_backend/pkg/logger"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		os.Exit(1)
	}

	// Initialize logger
	log := logger.New("info").WithService("reputation-service")
	log.Info("Starting Lamda Reputation Service")

	// Connect to blockchain (BSC for JobManager contract)
	blockchainClient, err := blockchain.NewEVMClient(cfg.BSCRPCURL)
	if err != nil {
		log.Error("Failed to connect to blockchain", "error", err)
		os.Exit(1)
	}
	defer blockchainClient.Close()

	// Wait for blockchain connection
	ctx, cancel := context.WithTimeout(context.Background(), 30)
	defer cancel()
	if err := blockchainClient.WaitForConnection(ctx, 30); err != nil {
		log.Error("Failed to wait for blockchain connection", "error", err)
		os.Exit(1)
	}

	// Initialize reputation service
	reputationService := reputation.NewService(blockchainClient, log, cfg.NodeReputationContractAddress, cfg.AdminWalletPrivateKey)

	// Start the service
	if err := reputationService.Start(context.Background()); err != nil {
		log.Error("Failed to start reputation service", "error", err)
		os.Exit(1)
	}

	log.Info("Reputation service started successfully")

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down Reputation Service...")
	log.Info("Reputation Service stopped")
}
