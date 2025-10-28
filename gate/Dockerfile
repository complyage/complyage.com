# ----------------------------------------------------------------------------------
# Stage 1: Build
# ----------------------------------------------------------------------------------

FROM golang:1.24.5-alpine AS builder

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
EXPOSE 8082

ENTRYPOINT ["./gate"]
