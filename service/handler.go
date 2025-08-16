package service

import "context"

type RateLimiterServer struct {
	pb.UnimplementedRateLimiterServer
}

func (s *RateLimiterServer) CheckRateLimit(ctx context.Context, req *pb.RateLimitRequest) (*pb.RateLimitResponse, error) {
	// If request reaches here, it passed the rate limiter
	return &pb.RateLimitResponse{
		Allowed: true,
		Message: "Request allowed",
	}, nil
}
