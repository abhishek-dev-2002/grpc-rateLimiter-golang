package redisclient

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func NewClient(addr string) {
	client = redis.NewClient(&redis.Options{
		Addr: addr,
	})
}

func GetClient() *redis.Client {
	if client == nil {
		NewClient("localhost:6379")
	}
	return client
}

func Increment(ctx context.Context, key string) (int64, error) {
	return client.Incr(ctx, key).Result()
}
