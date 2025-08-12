# syntax=docker/dockerfile:1
FROM golang:1.24-alpine

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    PATH="/go/bin:$PATH"

WORKDIR /app

# Install dependencies for building
RUN apk add --no-cache git build-base bash

# Install Air for hot reload
RUN go install github.com/air-verse/air@latest

# Copy module definitions (cache layer)
COPY api/go.mod api/go.sum ./
COPY base /base

# Download Go modules
RUN go mod download

# Copy application source
COPY api/ .

EXPOSE 8888

ENTRYPOINT ["air"]
