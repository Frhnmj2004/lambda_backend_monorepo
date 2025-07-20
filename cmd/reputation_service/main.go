package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	// Create context for blockchain connection checks
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Connect to BSC blockchain
	bscClient, err := blockchain.NewEVMClient(cfg.BSCRPCURL)
	if err != nil {
		log.Error("Failed to connect to BSC blockchain", "error", err)
		os.Exit(1)
	}
	defer bscClient.Close()

	// Wait for BSC connection
	if err := bscClient.WaitForConnection(ctx, 30*time.Second); err != nil {
		log.Error("Failed to wait for BSC connection", "error", err)
		os.Exit(1)
	}

	// Connect to opBNB blockchain
	opBNBClient, err := blockchain.NewEVMClient(cfg.OpBNBRPCURL)
	if err != nil {
		log.Error("Failed to connect to opBNB blockchain", "error", err)
		os.Exit(1)
	}
	defer opBNBClient.Close()

	// Wait for opBNB connection
	if err := opBNBClient.WaitForConnection(ctx, 30*time.Second); err != nil {
		log.Error("Failed to wait for opBNB connection", "error", err)
		os.Exit(1)
	}

	// Initialize reputation service
	reputationService := reputation.NewService(
		bscClient,
		opBNBClient,
		log,
		cfg.JobManagerContractAddress,
		cfg.NodeReputationContractAddress,
		cfg.AdminWalletPrivateKey,
	)

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
