package main

import (
	"log"
	"net"

	"github.com/abhishek-dev-2002/grpc-rateLimter-golang/pkg/redisclient"
	"github.com/abhishek-dev-2002/grpc-rateLimter-golang/service"
	"google.golang.org/grpc"
)

func main() {
	// Initialize Redis
	redisclient.NewClient("localhost:6379")

	// Listen on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create limiter (5 requests per minute per token)
	limiter := ratelimiter.NewLimiter(5)

	// Setup gRPC server with interceptor
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(ratelimiter.RateLimitInterceptor(limiter)),
	)

	// Register the service
	pb.RegisterRateLimiterServer(grpcServer, &service.RateLimiterServer{})

	log.Println("🚀 gRPC server running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
