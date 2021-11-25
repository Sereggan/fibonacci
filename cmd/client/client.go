package main

import (
	"context"
	"fibonachi/internal/delivery/grpc/pb"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

const address = "127.0.0.1:5300"

// Test client for grpc
func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Println("error:", err)
	}

	defer conn.Close()

	client := pb.NewFibonacciClient(conn)

	req, err := client.Post(context.Background(), &pb.Request{
		X: 1,
		Y: 10,
	})

	fmt.Println(req)
}
