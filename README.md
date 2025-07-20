# Lamda Network Backend

A production-ready Go backend for the Lamda Network, providing blockchain event listening, job dispatching, and reputation management services with IPFS integration.

## Architecture

The backend consists of four main services:

1. **Node Registry Service** - Listens to NodeReputation contract events on opBNB and maintains provider database
2. **Job Dispatcher Service** - Listens to JobManager contract events on BSC and dispatches jobs to providers via IPFS
3. **Reputation Service** - Listens to JobManager events on BSC and updates reputation on opBNB
4. **API Gateway** - Provides REST API endpoints for external integrations

## Prerequisites

- Go 1.21+
- PostgreSQL 13+
- NATS Server
- Access to BSC and opBNB RPC endpoints

## Smart Contracts

The backend integrates with two deployed smart contracts:

- **JobManager.sol** (BSC Testnet): `0xd9264B533dD53198C7aE345C6aFE8EF054303b53`
- **NodeReputation.sol** (opBNB Testnet): `0x108f2c400C9828d8044a5F6985f0C9589B90758D`

## IPFS Integration

The backend now uses IPFS for file storage instead of BNB Greenfield:

- **Input Files**: Frontend uploads input files to IPFS (via Pinata or similar) and provides the CID
- **Job Assignment**: Backend sends IPFS CIDs to node agents via NATS
- **Output Files**: Node agents upload results to IPFS and return the output CID
- **Storage**: All file references use IPFS Content Identifiers (CIDs)

## Quick Start

### 1. Clone and Setup

```bash
git clone <repository-url>
cd lamda_backend
go mod download
```

### 2. Environment Configuration

Copy the example environment file and configure it:

```bash
cp env.example .env
```

Edit `.env` with your configuration:

```env
# Database Configuration
DATABASE_URL=postgres://username:password@localhost:5432/lamda_db

# NATS Configuration
NATS_URL=nats://localhost:4222

# Blockchain RPC URLs
BSC_RPC_URL=https://bsc-dataseed1.binance.org/
OPBNB_RPC_URL=https://opbnb-mainnet-rpc.bnbchain.org

# Smart Contract Addresses
JOB_MANAGER_CONTRACT_ADDRESS=0xd9264B533dD53198C7aE345C6aFE8EF054303b53
NODE_REPUTATION_CONTRACT_ADDRESS=0x108f2c400C9828d8044a5F6985f0C9589B90758D

# Admin Wallet Private Key (for reputation updates)
ADMIN_WALLET_PRIVATE_KEY=your_admin_wallet_private_key_here

# API Gateway Configuration
API_PORT=8080

# Environment
ENVIRONMENT=development
```

### 3. Database Setup

The backend uses GORM auto-migration, so the database schema will be automatically created when you start the services. Just create the database:

```bash
# Create database
createdb lamda_db
```

The tables will be automatically created when you run the services for the first time.

### 4. Start Services

#### Option A: Individual Services

```bash
# Terminal 1: Node Registry Service
go run cmd/node_registry/main.go

# Terminal 2: Job Dispatcher Service
go run cmd/job_dispatcher/main.go

# Terminal 3: Reputation Service
go run cmd/reputation_service/main.go

# Terminal 4: API Gateway
go run cmd/api_gateway/main.go
```

#### Option B: Using Makefile

```bash
# Start all services
make start

# Start individual services
make start-node-registry
make start-job-dispatcher
make start-reputation-service
make start-api-gateway
```

## Production Deployment

### Docker Deployment

```bash
# Build the Docker image
docker build -t lamda-backend .

# Run with docker-compose
docker-compose up -d
```

### Systemd Service Files

Create systemd service files for each service:

```ini
# /etc/systemd/system/lamda-node-registry.service
[Unit]
Description=Lamda Node Registry Service
After=network.target

[Service]
Type=simple
User=lamda
WorkingDirectory=/opt/lamda-backend
ExecStart=/opt/lamda-backend/node-registry
Restart=always
EnvironmentFile=/opt/lamda-backend/.env

[Install]
WantedBy=multi-user.target
```

## API Endpoints

### Node Registry API

- `GET /api/nodes` - List active nodes
- `GET /api/nodes/{address}` - Get node details
- `POST /api/nodes/query` - Query nodes with filters

### Job Management API

- `GET /api/jobs` - List jobs
- `GET /api/jobs/{id}` - Get job details
- `POST /api/jobs/query` - Query jobs with filters
- `PUT /api/jobs/{id}/status` - Update job status

### Reputation API

- `GET /api/reputation/{address}` - Get provider reputation
- `GET /api/reputation/{address}/jobs` - Get completed jobs count

## Event Flow

1. **Provider Registration**: Node agents call `registerNode()` on NodeReputation contract
2. **Job Creation**: Frontend uploads input file to IPFS and calls `createJob()` on JobManager contract with the CID
3. **Job Dispatch**: Backend listens to JobCreated events and dispatches jobs with IPFS CIDs to providers via NATS
4. **Job Completion**: Providers complete jobs, upload results to IPFS, and call `confirmResult()` on JobManager
5. **Reputation Update**: Backend listens to JobConfirmed events and calls `incrementJobs()` on NodeReputation

## NATS Message Format

### Job Assignment Message

```json
{
  "jobId": "0x1234567890abcdef...",
  "dockerImage": "nvidia/cuda:11.8-base",
  "inputFileCID": "QmX...abc123"
}
```

## Monitoring and Logging

The backend includes comprehensive logging and monitoring:

- Structured logging with service identification
- Error tracking and reporting
- Health check endpoints
- Transaction monitoring

## Security Considerations

- Admin private key should be stored securely
- Use environment variables for sensitive configuration
- Implement rate limiting on API endpoints
- Validate all input data
- Use HTTPS in production

## Troubleshooting

### Common Issues

1. **Blockchain Connection Failed**
   - Check RPC URLs are accessible
   - Verify network connectivity
   - Check contract addresses are correct

2. **Database Connection Issues**
   - Verify PostgreSQL is running
   - Check database credentials
   - Ensure migrations have been run

3. **NATS Connection Issues**
   - Verify NATS server is running
   - Check NATS URL configuration
   - Ensure network connectivity

### Logs

Check service logs for detailed error information:

```bash
# View logs for a specific service
journalctl -u lamda-node-registry -f
```

## Development

### Running Tests

```bash
go test ./...
```

### Code Generation

Contract bindings are pre-generated. To regenerate:

```bash
# Install abigen
go install github.com/ethereum/go-ethereum/cmd/abigen

# Generate bindings
abigen --abi=contracts/JobManager.abi --pkg=contracts --out=pkg/contracts/jobmanager.go
abigen --abi=contracts/NodeReputation.abi --pkg=contracts --out=pkg/contracts/nodereputation.go
```

## License

[License information] 