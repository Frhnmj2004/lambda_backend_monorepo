# Lamda Backend Makefile

.PHONY: help build clean test start start-node-registry start-job-dispatcher start-reputation-service start-api-gateway docker-build docker-run

# Default target
help:
	@echo "Available targets:"
	@echo "  build                    - Build all services"
	@echo "  clean                    - Clean build artifacts"
	@echo "  test                     - Run tests"
	@echo "  start                    - Start all services"
	@echo "  start-node-registry      - Start node registry service"
	@echo "  start-job-dispatcher     - Start job dispatcher service"
	@echo "  start-reputation-service - Start reputation service"
	@echo "  start-api-gateway        - Start API gateway"
	@echo "  docker-build             - Build Docker image"
	@echo "  docker-run               - Run with Docker Compose"

# Build all services
build:
	@echo "Building all services..."
	go build -o bin/node-registry cmd/node_registry/main.go
	go build -o bin/job-dispatcher cmd/job_dispatcher/main.go
	go build -o bin/reputation-service cmd/reputation_service/main.go
	go build -o bin/api-gateway cmd/api_gateway/main.go
	@echo "Build complete!"

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin/
	@echo "Clean complete!"

# Run tests
test:
	@echo "Running tests..."
	go test ./...
	@echo "Tests complete!"

# Start all services (requires multiple terminals)
start:
	@echo "Starting all services..."
	@echo "Please run each service in a separate terminal:"
	@echo "Terminal 1: make start-node-registry"
	@echo "Terminal 2: make start-job-dispatcher"
	@echo "Terminal 3: make start-reputation-service"
	@echo "Terminal 4: make start-api-gateway"

# Start individual services
start-node-registry:
	@echo "Starting Node Registry Service..."
	go run cmd/node_registry/main.go

start-job-dispatcher:
	@echo "Starting Job Dispatcher Service..."
	go run cmd/job_dispatcher/main.go

start-reputation-service:
	@echo "Starting Reputation Service..."
	go run cmd/reputation_service/main.go

start-api-gateway:
	@echo "Starting API Gateway..."
	go run cmd/api_gateway/main.go

# Docker targets
docker-build:
	@echo "Building Docker image..."
	docker build -t lamda-backend .

docker-run:
	@echo "Running with Docker Compose..."
	docker-compose up -d

# Development helpers
dev-setup:
	@echo "Setting up development environment..."
	go mod download
	@echo "Development setup complete!"

# Linting and formatting
lint:
	@echo "Running linter..."
	golangci-lint run

format:
	@echo "Formatting code..."
	go fmt ./...

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod download
	go install github.com/ethereum/go-ethereum/cmd/abigen@latest
	@echo "Dependencies installed!" 