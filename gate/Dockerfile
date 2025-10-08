# syntax=docker/dockerfile:1
FROM golang:1.24-alpine

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    PATH="/go/bin:$PATH"

WORKDIR /app

RUN apk add --no-cache git build-base bash
RUN go install github.com/air-verse/air@latest

COPY gate/go.mod gate/go.sum ./
RUN go mod download

COPY gate/ .

EXPOSE 8888
ENTRYPOINT ["air"]
