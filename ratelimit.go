package ratelimit

import (
	grpc_ratelimit "github.com/grpc-ecosystem/go-grpc-middleware/ratelimit"
	"go.uber.org/ratelimit"
)

type limiter struct {
	ratelimit.Limiter
}

func (l *limiter) Limit() bool {
	l.Take()
	return false
}

// NewLimiter return new go-grpc Limiter, specified the number of requests you want to limit as a counts per second.
func NewLimiter(count int) grpc_ratelimit.Limiter {
	return &limiter{
		Limiter: ratelimit.New(count),
	}
}
