# syntax=docker/dockerfile:1
FROM golang:1.24-alpine

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    PATH="/go/bin:$PATH"

WORKDIR /app

# Install basic tools for development
RUN apk add --no-cache git build-base bash

# Install air
RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8888

# Leave it running for attach/debug (air can be manually run)
ENTRYPOINT ["air"]
