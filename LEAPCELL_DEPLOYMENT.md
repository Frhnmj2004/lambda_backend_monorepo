# Leapcell Deployment Guide for Lamda Backend

## 🚀 Quick Deploy to Leapcell

### 1. Connect Your Repository
1. Go to [Leapcell Dashboard](https://leapcell.com/dashboard)
2. Click "New Project" → "Deploy from GitHub"
3. Select your `lamda_backend` repository
4. Leapcell will automatically detect the Dockerfile and leapcell.json

### 2. Add PostgreSQL Database
1. In your Leapcell project, click "Add Service" → "PostgreSQL"
2. Leapcell will automatically provide `DATABASE_URL` environment variable

### 3. Configure Environment Variables
Add these environment variables in Leapcell dashboard:

```bash
# NATS Configuration (You'll need to set up NATS separately)
NATS_URL=nats://your-nats-service-url:4222

# Blockchain RPC URLs
BSC_RPC_URL=https://data-seed-prebsc-2-s1.binance.org:8545/
OPBNB_RPC_URL=https://opbnb-testnet-rpc.bnbchain.org

# Smart Contract Addresses
JOB_MANAGER_CONTRACT_ADDRESS=0xd9264B533dD53198C7aE345C6aFE8EF054303b53
NODE_REPUTATION_CONTRACT_ADDRESS=0x108f2c400C9828d8044a5F6985f0C9589B90758D

# Admin Wallet Private Key
ADMIN_WALLET_PRIVATE_KEY=2bc887ad1626908b26faa8be49182fb668d858986282baec90d81e9314ae47b3

# Environment
ENVIRONMENT=production
```

### 4. Deploy
Leapcell will automatically deploy your application. All services will run:
- ✅ API Gateway (port 8080)
- ✅ Job Dispatcher
- ✅ Node Registry  
- ✅ Reputation Service

## 🔧 NATS Setup Options

### Option 1: Leapcell NATS (Recommended)
1. In Leapcell dashboard, click "Add Service" → "NATS"
2. Copy the NATS URL from the service
3. Update `NATS_URL` environment variable

### Option 2: External NATS Service
Use any NATS service (Cloudflare, Upstash, etc.) and update `NATS_URL`

### Option 3: Self-hosted NATS
Deploy NATS separately and provide the URL

## 🌐 Node Agent Connection

After deployment, your node agent should connect to:

```bash
# NATS URL (from Leapcell NATS service)
NATS_URL=nats://your-leapcell-nats-url:4222

# API Gateway URL (from Leapcell deployment)
API_URL=https://your-leapcell-app-url.leapcell.com
```

## 📊 Monitoring

### Health Check
- **URL**: `https://your-app.leapcell.com/health`
- **Expected Response**: `{"service":"lamda-backend","status":"healthy"}`

### Logs
- View logs in Leapcell dashboard
- All services log to `/var/log/` directory

### Metrics
- CPU and memory usage available in dashboard
- Auto-scaling based on load

## 🔍 Troubleshooting

### Common Issues:

1. **Database Connection Failed**
   - Ensure PostgreSQL is added to Leapcell project
   - Check `DATABASE_URL` environment variable

2. **NATS Connection Failed**
   - Verify NATS service is running
   - Check `NATS_URL` environment variable

3. **Blockchain Connection Failed**
   - Verify RPC URLs are accessible
   - Check contract addresses are correct

4. **Port Issues**
   - Leapcell automatically handles port 8080
   - API Gateway runs on port 8080 by default

## 🎯 Production Checklist

- [ ] PostgreSQL database added
- [ ] NATS service configured
- [ ] All environment variables set
- [ ] Health check endpoint responding
- [ ] Node agent can connect to NATS
- [ ] Blockchain connections working
- [ ] Logs are being generated
- [ ] Auto-scaling configured

## 📝 Environment Variables Reference

| Variable | Required | Description |
|----------|----------|-------------|
| `DATABASE_URL` | ✅ | PostgreSQL connection string (auto-provided by Leapcell) |
| `NATS_URL` | ✅ | NATS server connection string |
| `BSC_RPC_URL` | ✅ | BSC testnet RPC endpoint |
| `OPBNB_RPC_URL` | ✅ | opBNB testnet RPC endpoint |
| `JOB_MANAGER_CONTRACT_ADDRESS` | ✅ | JobManager contract address |
| `NODE_REPUTATION_CONTRACT_ADDRESS` | ✅ | NodeReputation contract address |
| `ADMIN_WALLET_PRIVATE_KEY` | ✅ | Admin wallet private key |
| `API_PORT` | ❌ | Port for API Gateway (default: 8080) |
| `ENVIRONMENT` | ❌ | Set to "production" |

## 🚀 Leapcell Features

- **Auto-scaling**: Automatically scales based on load (1-3 instances)
- **Health checks**: Monitors service health every 30 seconds
- **Resource limits**: 0.5 CPU, 512Mi memory per instance
- **Zero-downtime deployments**: Rolling updates
- **Built-in monitoring**: CPU, memory, and network metrics

Your Lamda backend is now ready for Leapcell deployment! 🚀 