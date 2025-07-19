-- Migration: Create providers table
-- This table stores information about GPU providers in the Lamda network

CREATE TABLE IF NOT EXISTS providers (
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

-- Create index on wallet_address for fast lookups
CREATE INDEX IF NOT EXISTS idx_providers_wallet_address ON providers(wallet_address);

-- Create index on last_seen for finding active providers
CREATE INDEX IF NOT EXISTS idx_providers_last_seen ON providers(last_seen);

-- Create index on is_online for filtering active providers
CREATE INDEX IF NOT EXISTS idx_providers_is_online ON providers(is_online);

-- Create index on reputation_score for sorting by reputation
CREATE INDEX IF NOT EXISTS idx_providers_reputation_score ON providers(reputation_score);

-- Add comment to table
COMMENT ON TABLE providers IS 'Stores information about GPU providers in the Lamda network';

-- Add comments to columns
COMMENT ON COLUMN providers.wallet_address IS 'Ethereum wallet address of the provider';
COMMENT ON COLUMN providers.gpu_model IS 'Model of the GPU (e.g., RTX 4090, A100)';
COMMENT ON COLUMN providers.vram IS 'VRAM capacity in GB';
COMMENT ON COLUMN providers.last_seen IS 'Timestamp of last heartbeat from provider';
COMMENT ON COLUMN providers.is_online IS 'Whether the provider is currently online';
COMMENT ON COLUMN providers.total_jobs_completed IS 'Total number of jobs completed by this provider';
COMMENT ON COLUMN providers.reputation_score IS 'On-chain reputation score'; 