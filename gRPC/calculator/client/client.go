package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	pb "github.com/ibilalkayy/Small-Projects/gRPC/calculator/proto"
)

func main() {
	serverAddress := "localhost:50051"

	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to the server: %v", err)
	}

	defer conn.Close()

	client := pb.NewCalculatorClient(conn)

	addResult, err := client.Add(context.Background(), &pb.AddRequest{Num1: 4.0, Num2: 6.0})
	if err != nil {
		log.Fatalf("Add RPC failed: %v", err)
	}
	fmt.Printf("Add Result: %.2f\n", addResult.Response)

	subtractResult, err := client.Subtract(context.Background(), &pb.SubtractRequest{Num1: 14.0, Num2: 6.0})
	if err != nil {
		log.Fatalf("Subtract RPC failed: %v", err)
	}
	fmt.Printf("Subtract Result: %.2f\n", subtractResult.Response)

	multiplyResult, err := client.Multiply(context.Background(), &pb.MultiplyRequest{Num1: 4.0, Num2: 6.0})
	if err != nil {
		log.Fatalf("Multiply RPC failed: %v", err)
	}
	fmt.Printf("Multiply Result: %.2f\n", multiplyResult.Response)

	divideResult, err := client.Divide(context.Background(), &pb.DivideRequest{Num1: 40.0, Num2: 5.0})
	if err != nil {
		log.Fatalf("Divide RPC failed: %v", err)
	}
	fmt.Printf("Divide Result: %.2f\n", divideResult.Response)
}
