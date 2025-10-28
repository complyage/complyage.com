# ----------------------------------------------------------------------------------
# Stage 1: Build
# ----------------------------------------------------------------------------------

FROM golang:1.24.5-alpine AS builder

WORKDIR /app
RUN apk add --no-cache git

# Copy all project files
COPY . .
RUN go mod tidy

# Build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o deployment .

# ----------------------------------------------------------------------------------
# Stage 2: Runtime
# ----------------------------------------------------------------------------------

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/deployment .
EXPOSE 9090

ENTRYPOINT ["./deployment"]
