# gRPC Rate Limiter in Go with Redis

## Overview

This project demonstrates how to implement a **rate-limiting interceptor** in a **Go gRPC server**, utilizing **Redis** as the backend counter. The rate limiter restricts users from making more than a specified number of requests per second, enhancing API security and preventing abuse.

## Features

- **gRPC Server**: Built using Go and Protocol Buffers.
- **Rate Limiting**: Custom interceptor to limit requests per user.
- **Redis Backend**: Utilizes Redis for storing and checking request counts.
- **Interceptor Customization**: Easily adjustable rate limits and time windows.

## Prerequisites

- [Go 1.18+](https://go.dev/dl/)
- [Redis](https://redis.io/download)
- [Buf](https://buf.build/docs/installation)
- [Docker](https://www.docker.com/get-started) (optional, for Redis setup)

## Rate Limiting Logic
Current Implementation
We are using a fixed window counter approach with Redis.

How it works:

Each user has a Redis key like rate_limit:<user_token>.

On each request, the server increments the counter in Redis.

If the counter exceeds MaxHits within the time window (1 minute), requests are rejected.

The Redis key has a TTL of 1 minute, so the counter resets automatically after the window expires.

Key Points:

This method is strict: once the limit is hit, additional requests are blocked until the window resets.
The approach is simple and works well for basic per-user rate limiting.

Example Flow
MaxHits = 5 per minute

User sends 7 requests in one minute:
So first 7 req within 1 minute will be allowed and rest will be blocked until the counter resets to 0.
After 1 minute, the Redis key expires, and the counter resets to 0.

How It Works
Interceptor Integration:
The gRPC server integrates a UnaryServerInterceptor that wraps every incoming request.

Token Extraction:
The interceptor extracts the token from the gRPC request message.

Rate Limit Check:
The Limiter checks Redis for the request count and decides whether to allow or reject the request.

Request Handling:
If allowed: the request proceeds to the gRPC handler.
If rejected: a RateLimitResponse is returned with allowed: false and a descriptive message.

Redis as Backend:
Redis is used because it supports atomic operations and TTL-based expiry, making it ideal for distributed rate limiting across multiple gRPC server instances.
