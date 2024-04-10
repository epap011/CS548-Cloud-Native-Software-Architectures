package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	pb "example.com/grpc-calculator/proto"
	"google.golang.org/grpc"
)

func main() {
	var operand1, operand2 int
	flag.IntVar(&operand1, "a", 0, "First operand")
	flag.IntVar(&operand2, "b", 0, "Second operand")
	flag.Parse()

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCalculatorClient(conn)

	addRequest := &pb.AddRequest{
		Operands: &pb.Operands{
			A: int32(operand1),
			B: int32(operand2),
		},
	}

	res, err := client.Add(context.Background(), addRequest)
	if err != nil {
		log.Fatalf("Error calling Add: %v", err)
	}

	// Print the result of the addition operation
	fmt.Printf("%d + %d = %d\n", operand1, operand2, res.Result)
}
