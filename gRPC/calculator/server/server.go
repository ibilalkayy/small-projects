package main

import (
	"context"
	"log"
	"net"

	pb "github.com/ibilalkayy/Small-Projects/gRPC/calculator/proto"
	"google.golang.org/grpc"
)

type calculatorServer struct {
	pb.UnimplementedCalculatorServer
}

func (s *calculatorServer) Add(ctx context.Context, req *pb.AddRequest) (*pb.Response, error) {
	result := req.Num1 + req.Num2
	return &pb.Response{Response: result}, nil
}

func (s *calculatorServer) Subtract(ctx context.Context, req *pb.SubtractRequest) (*pb.Response, error) {
	result := req.Num1 - req.Num2
	return &pb.Response{Response: result}, nil
}

func (s *calculatorServer) Multiply(ctx context.Context, req *pb.MultiplyRequest) (*pb.Response, error) {
	result := req.Num1 * req.Num2
	return &pb.Response{Response: result}, nil
}

func (s *calculatorServer) Divide(ctx context.Context, req *pb.DivideRequest) (*pb.Response, error) {
	result := req.Num1 / req.Num2
	return &pb.Response{Response: result}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &calculatorServer{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
