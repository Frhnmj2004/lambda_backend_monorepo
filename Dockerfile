# Multi-stage build for Lamda Backend
FROM golang:1.21-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git ca-certificates tzdata

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build all services
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api_gateway ./cmd/api_gateway
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o job_dispatcher ./cmd/job_dispatcher
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o node_registry ./cmd/node_registry
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o reputation_service ./cmd/reputation_service

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests and supervisor
RUN apk --no-cache add ca-certificates tzdata supervisor

# Create non-root user
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

# Set working directory
WORKDIR /app

# Copy binaries from builder stage
COPY --from=builder /app/api_gateway .
COPY --from=builder /app/job_dispatcher .
COPY --from=builder /app/node_registry .
COPY --from=builder /app/reputation_service .

# Create supervisor configuration
RUN mkdir -p /etc/supervisor.d
COPY <<EOF /etc/supervisor.d/lamda.ini
[supervisord]
nodaemon=true
user=appuser

[program:api_gateway]
command=./api_gateway
directory=/app
autostart=true
autorestart=true
stderr_logfile=/var/log/api_gateway.err.log
stdout_logfile=/var/log/api_gateway.out.log

[program:job_dispatcher]
command=./job_dispatcher
directory=/app
autostart=true
autorestart=true
stderr_logfile=/var/log/job_dispatcher.err.log
stdout_logfile=/var/log/job_dispatcher.out.log

[program:node_registry]
command=./node_registry
directory=/app
autostart=true
autorestart=true
stderr_logfile=/var/log/node_registry.err.log
stdout_logfile=/var/log/node_registry.out.log

[program:reputation_service]
command=./reputation_service
directory=/app
autostart=true
autorestart=true
stderr_logfile=/var/log/reputation_service.err.log
stdout_logfile=/var/log/reputation_service.out.log
EOF

# Create log directory
RUN mkdir -p /var/log && chown -R appuser:appgroup /var/log

# Change ownership to non-root user
RUN chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

# Expose port for API Gateway
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# Start supervisor to manage all services
CMD ["supervisord", "-c", "/etc/supervisor.d/lamda.ini"] 