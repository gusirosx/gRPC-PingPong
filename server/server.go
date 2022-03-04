// Sample grpc-ping acts as an intermediary to the ping service.
package main

import (
	"context"
	"log"
	"net"

	pb "gRPC-Ping/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	port = ":50050"
)

// pingService is used to implement PingServiceServer.
type pingService struct {
	pb.UnimplementedPingServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen on port %s: %v", port, err)
	}
	log.Printf("grpc-ping: starting on port %s", port)

	srv := grpc.NewServer()
	pb.RegisterPingServiceServer(srv, &pingService{})
	// Register reflection service on gRPC server.
	reflection.Register(srv)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *pingService) Send(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	log.Println("Received :", req.Message)
	response := &pb.Response{
		Index:      1,
		Message:    "Hello there", //req.GetMessage(),
		ReceivedOn: timestamppb.Now(),
	}

	return response, nil
}
