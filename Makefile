# Lamda Network Backend Makefile

.PHONY: help build run test clean docker-build docker-run migrate dev

# Default target
help:
	@echo "Lamda Network Backend - Available Commands:"
	@echo ""
	@echo "Development:"
	@echo "  make dev          - Run all services in development mode"
	@echo "  make build        - Build all service binaries"
	@echo "  make test         - Run all tests"
	@echo "  make clean        - Clean build artifacts"
	@echo ""
	@echo "Database:"
	@echo "  make migrate      - Run database migrations"
	@echo ""
	@echo "Docker:"
	@echo "  make docker-build - Build Docker image"
	@echo "  make docker-run   - Run Docker container"
	@echo ""
	@echo "Individual Services:"
	@echo "  make api-gateway  - Run API Gateway service"
	@echo "  make node-registry - Run Node Registry service"
	@echo "  make job-dispatcher - Run Job Dispatcher service"
	@echo "  make reputation-service - Run Reputation Service"

# Build all services
build:
	@echo "Building all services..."
	go build -o bin/api_gateway ./cmd/api_gateway
	go build -o bin/job_dispatcher ./cmd/job_dispatcher
	go build -o bin/node_registry ./cmd/node_registry
	go build -o bin/reputation_service ./cmd/reputation_service
	@echo "Build complete!"

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin/
	go clean

# Database migrations
migrate:
	@echo "Running database migrations..."
	@if [ -f migrations/001_create_providers_table.sql ]; then \
		echo "Migration file found. Please run manually:"; \
		echo "psql -d lamda_db -f migrations/001_create_providers_table.sql"; \
	else \
		echo "No migration files found."; \
	fi

# Docker commands
docker-build:
	@echo "Building Docker image..."
	docker build -t lamda-backend .

docker-run:
	@echo "Running Docker container..."
	docker run -p 8080:8080 --env-file .env lamda-backend

# Development mode - run all services
dev:
	@echo "Starting development environment..."
	@echo "Make sure you have PostgreSQL and NATS running!"
	@echo ""
	@echo "Starting services in background..."
	@echo "API Gateway will be available at http://localhost:8080"
	@echo "Press Ctrl+C to stop all services"
	@echo ""
	@trap 'kill %1 %2 %3 %4' SIGINT; \
	go run cmd/api_gateway/main.go & \
	go run cmd/node_registry/main.go & \
	go run cmd/job_dispatcher/main.go & \
	go run cmd/reputation_service/main.go & \
	wait

# Individual service commands
api-gateway:
	@echo "Starting API Gateway..."
	go run cmd/api_gateway/main.go

node-registry:
	@echo "Starting Node Registry..."
	go run cmd/node_registry/main.go

job-dispatcher:
	@echo "Starting Job Dispatcher..."
	go run cmd/job_dispatcher/main.go

reputation-service:
	@echo "Starting Reputation Service..."
	go run cmd/reputation_service/main.go

# Dependencies
deps:
	@echo "Installing dependencies..."
	go mod download
	go mod tidy

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Lint code
lint:
	@echo "Linting code..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not found. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

# Generate documentation
docs:
	@echo "Generating documentation..."
	@if command -v swag >/dev/null 2>&1; then \
		swag init -g cmd/api_gateway/main.go; \
	else \
		echo "swag not found. Install with: go install github.com/swaggo/swag/cmd/swag@latest"; \
	fi

# Setup development environment
setup:
	@echo "Setting up development environment..."
	@echo "1. Installing dependencies..."
	go mod download
	@echo "2. Creating bin directory..."
	mkdir -p bin
	@echo "3. Setup complete!"
	@echo ""
	@echo "Next steps:"
	@echo "1. Copy env.example to .env and configure"
	@echo "2. Start PostgreSQL and NATS"
	@echo "3. Run 'make migrate' to setup database"
	@echo "4. Run 'make dev' to start all services"

# Production build
prod-build:
	@echo "Building for production..."
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/api_gateway ./cmd/api_gateway
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/job_dispatcher ./cmd/job_dispatcher
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/node_registry ./cmd/node_registry
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/reputation_service ./cmd/reputation_service
	@echo "Production build complete!"

# Health check
health:
	@echo "Checking service health..."
	@curl -f http://localhost:8080/health || echo "API Gateway not responding" 