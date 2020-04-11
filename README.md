# rate-limiter-grpc-go
grpc-go interceptor that can limit grpc requests both client and server.
`Limiter` allow requests only if the number of requests per second is less than the specified number.

## Usage

- client

```go
// allow two requests per seconds
conn, err := grpc.Dial(address,grpc.WithUnaryInterceptor(ratelimit.UnaryClientInterceptor(ratelimit.NewLimiter(2))))
...
client := pb.NewService(conn)
```

- server
```bash
receive request at 1586597159
receive request at 1586597159
receive request at 1586597160
receive request at 1586597160
receive request at 1586597161
receive request at 1586597161
receive request at 1586597162
receive request at 1586597162
receive request at 1586597163
receive request at 1586597163
```

## LICENSE
MIT