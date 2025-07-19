# Lamda Network Backend

A comprehensive Go microservices backend for the Lamda Network, providing event-driven GPU marketplace orchestration with blockchain integration.

## Architecture Overview

The Lamda Network backend is built as a set of event-driven Go microservices that orchestrate the entire GPU marketplace ecosystem:

- **API Gateway**: Fiber-based HTTP API for frontend communication
- **Node Registry**: Manages GPU provider registration and heartbeats
- **Job Dispatcher**: Handles job creation and assignment to providers
- **Reputation Service**: Updates on-chain reputation scores
- **Event-Driven Communication**: All services communicate via NATS messaging

## System Integration

This backend serves as the central nervous system connecting:

1. **Frontend (React)**: User interface for renters and providers
2. **Blockchain Contracts**: 
   - `JobManager.sol` (BSC): High-value job creation and payment escrow
   - `NodeReputation.sol` (opBNB): Provider registration and reputation
3. **Node Agents**: Go executables running on provider GPU machines
4. **BNB Greenfield**: Decentralized data storage

## Project Structure

```
/lamda_backend
|-- api/
|   |-- controller/          # HTTP request handlers
|   |-- router/             # Route definitions
|-- cmd/
|   |-- api_gateway/        # API Gateway service entrypoint
|   |-- job_dispatcher/     # Job Dispatcher service entrypoint
|   |-- node_registry/      # Node Registry service entrypoint
|   |-- reputation_service/ # Reputation service entrypoint
|-- config/                 # Configuration management
|-- internal/
|   |-- auth/              # SIWE authentication
|   |-- job_dispatcher/    # Job dispatching business logic
|   |-- node_registry/     # Node registry business logic
|   |-- reputation/        # Reputation management logic
|-- migrations/            # Database migrations
|-- pkg/
|   |-- blockchain/        # Ethereum client wrapper
|   |-- database/          # Database connection wrapper
|   |-- logger/            # Structured logging wrapper
|   |-- nats/              # NATS messaging wrapper
|-- go.mod                 # Go module definition
|-- Dockerfile             # Multi-stage container build
|-- env.example            # Environment variables template
```

## Services

### API Gateway (`cmd/api_gateway/`)
- **Purpose**: HTTP API for frontend communication
- **Port**: 8080 (configurable)
- **Endpoints**:
  - `GET /health` - Health check
  - `GET /api/v1/nodes` - List active GPU providers
  - `GET /api/v1/nodes/:address` - Get specific provider
  - `GET /api/v1/nodes/stats` - Provider statistics
  - `GET /api/v1/jobs` - List jobs
  - `GET /api/v1/jobs/:id` - Get specific job
  - `GET /api/v1/jobs/renter/:address` - Jobs by renter
  - `GET /api/v1/jobs/provider/:address` - Jobs by provider
  - `GET /api/v1/jobs/stats` - Job statistics

### Node Registry (`cmd/node_registry/`)
- **Purpose**: Manages GPU provider registration and heartbeats
- **Blockchain**: Listens to `NodeReputation.sol` on opBNB
- **Events**: `NodeRegistered`, `NodeHeartbeat`
- **NATS**: Subscribes to `nodes.query` for provider queries
- **Database**: PostgreSQL for provider state

### Job Dispatcher (`cmd/job_dispatcher/`)
- **Purpose**: Handles job creation and assignment
- **Blockchain**: Listens to `JobManager.sol` on BSC
- **Events**: `JobCreated`
- **NATS**: Publishes to `jobs.dispatch.<provider_address>`
- **Database**: PostgreSQL for job tracking

### Reputation Service (`cmd/reputation_service/`)
- **Purpose**: Updates on-chain reputation scores
- **Blockchain**: Listens to `JobManager.sol` on BSC, calls `NodeReputation.sol` on opBNB
- **Events**: `JobConfirmed`
- **Function**: Calls `incrementJobs(address provider)` on-chain

## Configuration

Copy `env.example` to `.env` and configure the following variables:

```bash
# Database
DATABASE_URL=postgres://user:password@localhost:5432/lamda_db

# NATS
NATS_URL=nats://localhost:4222

# Blockchain RPC URLs
BSC_RPC_URL=https://bsc-dataseed1.binance.org/
OPBNB_RPC_URL=https://opbnb-mainnet-rpc.bnbchain.org

# Smart Contract Addresses
JOB_MANAGER_CONTRACT_ADDRESS=0x...
NODE_REPUTATION_CONTRACT_ADDRESS=0x...

# Admin Wallet (for reputation updates)
ADMIN_WALLET_PRIVATE_KEY=your_private_key_here

# API Gateway
API_PORT=8080
ENVIRONMENT=development
```

## Database Schema

### Providers Table
```sql
CREATE TABLE providers (
    id SERIAL PRIMARY KEY,
    wallet_address VARCHAR(42) NOT NULL UNIQUE,
    gpu_model VARCHAR(100) NOT NULL,
    vram INTEGER NOT NULL,
    last_seen TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    is_online BOOLEAN DEFAULT true,
    total_jobs_completed INTEGER DEFAULT 0,
    reputation_score INTEGER DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
```

## NATS Messaging

### Subjects
- `nodes.query` - Query active providers (request-reply)
- `jobs.query` - Query jobs (request-reply)
- `jobs.dispatch.<provider_address>` - Job assignment to specific provider

### Message Formats

#### Job Assignment
```json
{
  "jobId": "job_123",
  "dockerImage": "nvidia/cuda:11.8-base",
  "greenfieldInputUrl": "https://greenfield.bnbchain.org/bucket/input",
  "greenfieldOutputBucket": "output-bucket",
  "greenfieldOutputName": "result.tar.gz"
}
```

## Development

### Prerequisites
- Go 1.21+
- PostgreSQL
- NATS Server
- Docker (optional)

### Local Development
```bash
# Install dependencies
go mod download

# Run database migrations
psql -d lamda_db -f migrations/001_create_providers_table.sql

# Start NATS server
nats-server

# Run services (in separate terminals)
go run cmd/api_gateway/main.go
go run cmd/node_registry/main.go
go run cmd/job_dispatcher/main.go
go run cmd/reputation_service/main.go
```

### Testing
```bash
# Run all tests
go test ./...

# Run specific service tests
go test ./internal/node_registry/...
```

## Deployment

### Docker
```bash
# Build image
docker build -t lamda-backend .

# Run container
docker run -p 8080:8080 --env-file .env lamda-backend
```

### Railway Deployment
1. Connect your repository to Railway
2. Set environment variables in Railway dashboard
3. Deploy - Railway will automatically build and run the Dockerfile

### Production Considerations
- Use managed PostgreSQL (e.g., Railway Postgres)
- Use managed NATS (e.g., Railway NATS)
- Set up proper monitoring and logging
- Configure CORS for your frontend domain
- Use HTTPS in production
- Secure admin wallet private key

## API Documentation

### Authentication
The API uses SIWE (Sign-In With Ethereum) for wallet-based authentication.

### Rate Limiting
Consider implementing rate limiting for production use.

### Error Handling
All endpoints return consistent error responses:
```json
{
  "error": "Error message",
  "success": false
}
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## License

This project is licensed under the MIT License.

## Support

For support and questions, please open an issue in the repository. 