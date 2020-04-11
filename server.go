package ratelimit

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	grpc_ratelimit "github.com/grpc-ecosystem/go-grpc-middleware/ratelimit"
)

// UnaryServerInterceptor return server unary interceptor that limit requests.
func UnaryServerInterceptor(limiter grpc_ratelimit.Limiter) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if limiter.Limit() {
			return nil, status.Errorf(codes.ResourceExhausted, "%s have been rejected by rate limiting.", info.FullMethod)
		}

		return handler(ctx, req)
	}
}

// StreamServerInterceptor return stream server unary interceptor that limit requests.
func StreamServerInterceptor(limiter grpc_ratelimit.Limiter) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if limiter.Limit() {
			return status.Errorf(codes.ResourceExhausted, "%s have been rejected by rate limiting.", info.FullMethod)
		}

		return handler(srv, stream)
	}
}
