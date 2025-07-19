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

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates tzdata

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

# Change ownership to non-root user
RUN chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

# Expose port for API Gateway
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# Default command (can be overridden)
CMD ["./api_gateway"] 