# ----------------------------------------------------------------------------------
# Stage 1: Build
# ----------------------------------------------------------------------------------

FROM golang:1.24.5-alpine AS builder

WORKDIR /app
RUN apk add --no-cache git

# ----------------------------------------------------------------------------------
# Remove Dev go mod and update to the github
# ----------------------------------------------------------------------------------

RUN rm -f go.mod go.sum \
    && go mod init api \
    && go get github.com/complyage/base@latest \
    && go get github.com/ralphferrara/aria@latest \
    && go mod tidy

# ----------------------------------------------------------------------------------
# Copy and Regenerate go.mod
# ----------------------------------------------------------------------------------

COPY . .
RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api .

# ----------------------------------------------------------------------------------
# Stage 2: Runtime
# ----------------------------------------------------------------------------------

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/api .
EXPOSE 8081

ENTRYPOINT ["./api"]
