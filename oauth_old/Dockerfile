# syntax=docker/dockerfile:1
FROM golang:1.24-alpine

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    PATH="/go/bin:$PATH"

WORKDIR /app

RUN apk add --no-cache git build-base bash
RUN go install github.com/air-verse/air@latest

COPY oauth/go.mod oauth/go.sum ./
COPY base /base
RUN go mod download

COPY oauth/ .

EXPOSE 8888
ENTRYPOINT ["air"]
