# ----------------------------------------------------------------------------------
# Stage 1: Build
# ----------------------------------------------------------------------------------
FROM golang:1.24.5-alpine AS builder

# Set working directory
WORKDIR /app

# Install git for Go modules
RUN apk add --no-cache git

# ----------------------------------------------------------------------------------
# Copy production go.mod and go.sum first for dependency caching
# ----------------------------------------------------------------------------------

COPY . .

# Build static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o oauth .

# ----------------------------------------------------------------------------------
# Stage 2: Runtime
# ----------------------------------------------------------------------------------
FROM alpine:latest

WORKDIR /app

# Copy the built binary from builder
COPY --from=builder /app/oauth .

# Expose port
EXPOSE 8083

# Run the binary
ENTRYPOINT ["./oauth"]
