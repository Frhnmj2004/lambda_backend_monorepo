package config

import (
	"fmt"
	"os"
	"strconv"

	"lamda_backend/pkg/blockchain"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the Lamda backend
type Config struct {
	// Database
	DatabaseURL string

	// NATS
	NATSURL string

	// Blockchain RPC URLs
	BSCRPCURL   string
	OpBNBRPCURL string

	// Contract Addresses
	JobManagerContractAddress     string
	NodeReputationContractAddress string

	// Admin wallet for reputation updates
	AdminWalletPrivateKey string

	// API Gateway
	APIPort string

	// Environment
	Environment string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	// Load .env file if it exists (only in development)
	if os.Getenv("ENVIRONMENT") != "production" {
		if err := godotenv.Load(); err != nil {
			// Don't fail if .env doesn't exist
			fmt.Println("No .env file found, using environment variables")
		}
	}

	// Get port from Leapcell PORT or default to API_PORT
	port := getEnv("PORT", getEnv("API_PORT", "8080"))

	config := &Config{
		DatabaseURL:                   getEnv("DATABASE_URL", "postgres://lamda:lamda123@localhost:5432/lamda_db"),
		NATSURL:                       getEnv("NATS_URL", "nats://localhost:4222"),
		BSCRPCURL:                     getEnv("BSC_RPC_URL", "https://data-seed-prebsc-2-s1.binance.org:8545/"),
		OpBNBRPCURL:                   getEnv("OPBNB_RPC_URL", "https://opbnb-testnet-rpc.bnbchain.org"),
		JobManagerContractAddress:     getEnv("JOB_MANAGER_CONTRACT_ADDRESS", ""),
		NodeReputationContractAddress: getEnv("NODE_REPUTATION_CONTRACT_ADDRESS", ""),
		AdminWalletPrivateKey:         getEnv("ADMIN_WALLET_PRIVATE_KEY", ""),
		APIPort:                       port,
		Environment:                   getEnv("ENVIRONMENT", "development"),
	}

	// Validate required fields
	if config.JobManagerContractAddress == "" {
		return nil, fmt.Errorf("JOB_MANAGER_CONTRACT_ADDRESS is required")
	}
	if config.NodeReputationContractAddress == "" {
		return nil, fmt.Errorf("NODE_REPUTATION_CONTRACT_ADDRESS is required")
	}
	if config.AdminWalletPrivateKey == "" {
		return nil, fmt.Errorf("ADMIN_WALLET_PRIVATE_KEY is required")
	}

	return config, nil
}

// CreateBSCClient creates a blockchain client for BSC network
func (c *Config) CreateBSCClient() (*blockchain.EVMClient, error) {
	return blockchain.NewEVMClient(c.BSCRPCURL)
}

// CreateOpBNBClient creates a blockchain client for opBNB network
func (c *Config) CreateOpBNBClient() (*blockchain.EVMClient, error) {
	return blockchain.NewEVMClient(c.OpBNBRPCURL)
}

// getEnv gets an environment variable with a fallback default value
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

// getEnvInt gets an environment variable as integer with a fallback default value
func getEnvInt(key string, fallback int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return fallback
}
