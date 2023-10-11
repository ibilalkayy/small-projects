package main

import (
	"context"
	"fmt"

	hello "github.com/ibilalkayy/Small-Projects/gRPC/hello-world/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	c := hello.NewGreeterClient(conn)
	resp, err := c.SayHello(context.Background(), &hello.HelloRequest{Name: "Bilal"})
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Message)
}
