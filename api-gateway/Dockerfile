FROM golang:1.25.3-alpine AS builder

WORKDIR /app

# Install dependencies
RUN apk add --no-cache git openssl

# Copy workspace and root module files first
COPY go.work go.work.sum* go.mod ./

# Copy all module files for workspace
COPY common/go.mod common/go.sum* ./common/
COPY api-gateway/go.mod api-gateway/go.sum ./api-gateway/
COPY driver-service/go.mod driver-service/go.sum* ./driver-service/

# Download dependencies for api-gateway
WORKDIR /app/api-gateway
RUN go mod download

# Copy source code
WORKDIR /app
COPY common ./common
COPY api-gateway ./api-gateway

# Build the application
WORKDIR /app/api-gateway
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates curl

WORKDIR /root/

# Copy the binary and config from builder
COPY --from=builder /app/api-gateway/main .
COPY --from=builder /app/api-gateway/internal/config ./internal/config
COPY --from=builder /app/api-gateway/pkg/token/keys ./pkg/token/keys

EXPOSE 8080

CMD ["./main"]
