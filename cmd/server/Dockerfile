# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the server binary
RUN CGO_ENABLED=0 GOOS=linux go build -o grpc-server ./cmd/server

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the server binary
COPY --from=builder /app/grpc-server .

# Expose gRPC port
EXPOSE 50051

# Run the server
CMD ["./grpc-server"]
