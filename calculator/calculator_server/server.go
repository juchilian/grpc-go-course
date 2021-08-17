package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"example.com/m/v2/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct {
}

func (*server) Calculate(ctx context.Context, req *calculatorpb.CalculatorRequest) (*calculatorpb.CalculatorResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	n1 := req.GetCalculator().GetN1()
	n2 := req.GetCalculator().GetN2()
	res := &calculatorpb.CalculatorResponse{
		Result: n1 + n2,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello World")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
