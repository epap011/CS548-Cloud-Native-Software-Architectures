package main

import (
	//This package provides context for request handling, cancellation, and timeouts.
	"context"
	//This package provides functions for logging.
	"log"
	//This package provides support for networking operations, like listening for incoming connections.
	"net"
	//This is the generated Go code from the Protocol Buffers file calculator.proto

	pb "example.com/grpc-calculator/proto"
	"google.golang.org/grpc"
	
// this method is a part of the server struct. It implements the Add RPC method defined in the gRPC service
func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	result := req.Operands.A + req.Operands.B
	println("Add method called with", req.Operands.A, "and", req.Operands.B)
	return &pb.AddResponse{Result: result}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051") // creates a tcp listener on port 50051

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterCalculatorServer(s, &server{})

	log.Println("Server started at :50051")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
