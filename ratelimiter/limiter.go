package ratelimiter

import (
	"context"
	"fmt"
	"time"

	"github.com/abhishek-dev-2002/grpc-rateLimter-golang/pb/pb"
	"github.com/abhishek-dev-2002/grpc-rateLimter-golang/pkg/redisclient"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Limiter struct {
	MaxHits int
}

func NewLimiter(maxHits int) *Limiter {
	return &Limiter{MaxHits: maxHits}
}

func (l *Limiter) Limit(ctx context.Context, token string) error {
	client := redisclient.GetClient()
	key := fmt.Sprintf("rate_limit:%s", token)

	count, err := client.Get(ctx, key).Int64()
	if err != nil && err.Error() != "redis: nil" {
		return status.Errorf(codes.Internal, "could not get rate limit count: %v", err)
	}

	if count >= int64(l.MaxHits) {
		return status.Errorf(codes.ResourceExhausted, "rate limit exceeded")
	}

	_, err = redisclient.Increment(ctx, key)
	if err != nil {
		return status.Errorf(codes.Internal, "could not increment rate limit count: %v", err)
	}

	client.Expire(ctx, key, time.Minute)
	return nil
}

// gRPC interceptor
func RateLimitInterceptor(limiter *Limiter) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		token := req.(*pb.RateLimitRequest).Token
		if err := limiter.Limit(ctx, token); err != nil {
			return &pb.RateLimitResponse{
				Allowed: false,
				Message: err.Error(),
			}, nil
		}
		return handler(ctx, req)
	}
}
