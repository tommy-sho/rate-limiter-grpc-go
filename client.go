package ratelimit

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	grpc_ratelimit "github.com/grpc-ecosystem/go-grpc-middleware/ratelimit"
)

// UnaryServerInterceptor return server unary interceptor that limit requests.
func UnaryClientInterceptor(limiter grpc_ratelimit.Limiter) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		if limiter.Limit() {
			return status.Errorf(codes.ResourceExhausted, "%s have been rejected by rate limiting.", method)
		}

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

// StreamServerInterceptor return stream server unary interceptor that limit requests.
func StreamClientInterceptor(limiter grpc_ratelimit.Limiter) grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		if limiter.Limit() {
			return nil, status.Errorf(codes.ResourceExhausted, "%s have been rejected by rate limiting.", method)
		}

		return streamer(ctx, desc, cc, method, opts...)
	}
}
