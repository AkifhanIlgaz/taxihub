FROM golang:1.25.3-alpine AS builder

WORKDIR /app

# Install dependencies
RUN apk add --no-cache git

# Copy workspace and root module files first
COPY go.work go.work.sum* go.mod ./

# Copy all module files for workspace
COPY common/go.mod common/go.sum* ./common/
COPY driver-service/go.mod driver-service/go.sum ./driver-service/
COPY api-gateway/go.mod api-gateway/go.sum* ./api-gateway/

# Download dependencies for driver-service
WORKDIR /app/driver-service
RUN go mod download

# Copy source code
WORKDIR /app
COPY common ./common
COPY driver-service ./driver-service

# Build the application
WORKDIR /app/driver-service
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates curl

WORKDIR /root/

# Copy the binary and config from builder
COPY --from=builder /app/driver-service/main .
COPY --from=builder /app/driver-service/internal/config ./internal/config

EXPOSE 50052

CMD ["./main"]
