package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"example.com/m/v2/greet/greetpb"
	"example.com/m/v2/repository"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type server struct {
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	err := saveToRedis(req)
	if err != nil {
		log.Fatalf("Failed to Save to Redis: %v", err)
	}
	res, err := makeResponse(req)
	if err != nil {
		log.Fatalf("Failed to make Response: %v", err)
	}
	return res, nil
}

func saveToRedis(req *greetpb.GreetRequest) error {
	r := repository.NewRedis()
	defer r.CloseRedis()

	// Set to Redis
	if err := r.Set("FirstName", []byte(req.GetGreeting().GetFirstName())); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func makeResponse(req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := greetpb.GreetResponse{
		Result: result,
	}
	return &res, nil
}

func main() {
	fmt.Println("Hello World")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})
	// Save to Redis

	// Send back
	fmt.Printf("Before Serve")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	fmt.Printf("After Serve")
}
