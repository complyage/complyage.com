# ----------------------------------------------------------------------------------
# Stage 1: Build
# ----------------------------------------------------------------------------------

FROM golang:1.24.5-alpine AS builder

WORKDIR /app
RUN apk add --no-cache git

COPY ../base /base
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ui .

# ----------------------------------------------------------------------------------
# Stage 2: Runtime
# ----------------------------------------------------------------------------------
      
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/ui .
EXPOSE 8080

ENTRYPOINT ["./ui"]
