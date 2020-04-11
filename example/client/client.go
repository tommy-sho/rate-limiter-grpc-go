package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"

	ratelimit "github.com/tommy-sho/rate-limiter-grpc-go"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
	requestTime = 100
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(),
		grpc.WithUnaryInterceptor(ratelimit.UnaryClientInterceptor(ratelimit.NewLimiter(2))))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	for i := 0; i < requestTime; i++ {
		if err := sendRequest(c, name); err != nil {
			log.Fatal("sendRequest: %w", err)
		}
	}
}

func sendRequest(c pb.GreeterClient, name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		return fmt.Errorf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", r.GetMessage())

	return nil
}
