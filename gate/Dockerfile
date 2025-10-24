# ----------------------------------------------------------------------------------
# Stage 1: Build
# ----------------------------------------------------------------------------------

FROM golang:1.24.5-alpine AS builder

WORKDIR /app
RUN apk add --no-cache git

COPY . .
RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gate .

# ----------------------------------------------------------------------------------
# Stage 2: Runtime
# ----------------------------------------------------------------------------------

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/gate .
EXPOSE 8081

ENTRYPOINT ["./gate"]
