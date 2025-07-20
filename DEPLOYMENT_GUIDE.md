# Lamda Backend Deployment Guide

## Quick Setup (No Migration Scripts!)

The backend now uses **GORM auto-migration**, which means the database schema is automatically created when you start the services. No manual migration scripts needed!

## Prerequisites

- Go 1.21+
- PostgreSQL 13+
- NATS Server
- Access to BSC and opBNB RPC endpoints

## Deployment Steps

### 1. Environment Setup

```bash
# Copy and configure environment
cp env.example .env
# Edit .env with your configuration
```

### 2. Database Setup

```bash
# Just create the database - that's it!
createdb lamda_db
```

### 3. Start Services

The database tables will be automatically created when services start.

```bash
# Option A: Individual services
go run cmd/node_registry/main.go
go run cmd/job_dispatcher/main.go
go run cmd/reputation_service/main.go
go run cmd/api_gateway/main.go

# Option B: Using Makefile
make start-node-registry
make start-job-dispatcher
make start-reputation-service
make start-api-gateway
```

## What Happens Automatically

When you start each service:

1. **Node Registry Service**: Creates `providers` table with IPFS-compatible schema
2. **Job Dispatcher Service**: Creates `jobs` table with IPFS CID columns
3. **Reputation Service**: No database tables needed
4. **API Gateway**: No database tables needed

## Database Schema

The following tables are automatically created:

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

### Jobs Table
```sql
CREATE TABLE jobs (
    id VARCHAR(66) PRIMARY KEY,
    renter_address VARCHAR(42) NOT NULL,
    provider_address VARCHAR(42) NOT NULL,
    docker_image VARCHAR(255),
    input_file_cid TEXT,           -- IPFS CID for input file
    output_file_cid TEXT,          -- IPFS CID for output file
    payment_amount VARCHAR(78) NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'created',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    assigned_at TIMESTAMP WITH TIME ZONE,
    completed_at TIMESTAMP WITH TIME ZONE,
    failed_at TIMESTAMP WITH TIME ZONE,
    error_message TEXT,
    block_number BIGINT,
    transaction_hash VARCHAR(66)
);
```

## Benefits of Auto-Migration

✅ **Zero Migration Scripts**: No SQL files to manage
✅ **Automatic Schema Creation**: Tables created on first run
✅ **Version Control**: Schema changes tracked in Go code
✅ **Deployment Simplicity**: Just start the services
✅ **Consistency**: Schema always matches the code

## Troubleshooting

### Database Connection Issues
```bash
# Check if PostgreSQL is running
pg_isready -h localhost -p 5432

# Check if database exists
psql -l | grep lamda_db
```

### Service Startup Issues
```bash
# Check logs for auto-migration errors
# Look for "AutoMigrate" in the logs
```

## Production Deployment

For production, the process is the same:

1. Create the database
2. Set environment variables
3. Start the services

The schema will be automatically created and maintained by GORM.

## IPFS Integration

The backend is fully configured for IPFS:
- Input files referenced by IPFS CIDs
- Output files stored as IPFS CIDs
- NATS messages use IPFS CIDs
- Database stores IPFS CIDs

No additional configuration needed for IPFS integration! 