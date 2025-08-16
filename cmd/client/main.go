package main

import (
	"context"
	"fmt"
	"log"

	"github.com/abhishek-dev-2002/grpc-rateLimter-golang/pb/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewRateLimiterClient(conn)

	for i := 0; i < 7; i++ {
		resp, err := client.CheckRateLimit(context.Background(), &pb.RateLimitRequest{Token: "user123"})
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		fmt.Printf("[%d] Allowed: %v, Message: %s\n", i+1, resp.Allowed, resp.Message)
	}
}
