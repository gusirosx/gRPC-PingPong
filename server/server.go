// Sample grpc-ping acts as an intermediary to the ping service.
package main

import (
	"context"
	"log"
	"net"

	pb "gRPC-Ping/proto"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	log.Printf("grpc-ping: starting at %v", lis.Addr())

	srv := grpc.NewServer()
	pb.RegisterPingServiceServer(srv, &pingService{})
	// Register reflection service on gRPC server.
	reflection.Register(srv)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *pingService) Send(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	log.Print("sending ping response")
	return &pb.Response{
		Pong: &pb.Pong{
			Index:      1,
			Message:    req.GetMessage(),
			ReceivedOn: ptypes.TimestampNow(),
		},
	}, nil
}

// func (s *pingService) SendUpstream(ctx context.Context, req *pb.Request) (*pb.Response, error) {
// 	if conn == nil {
// 		return nil, fmt.Errorf("no upstream connection configured")
// 	}

// 	p := &pb.Request{
// 		Message: req.GetMessage() + " (relayed)",
// 	}

// 	hostWithoutPort := strings.Split(os.Getenv("GRPC_PING_HOST"), ":")[0]
// 	tokenAudience := "https://" + hostWithoutPort
// 	resp, err := PingRequest(conn, p, tokenAudience, os.Getenv("GRPC_PING_UNAUTHENTICATED") == "")
// 	if err != nil {
// 		log.Printf("PingRequest: %q", err)
// 		c := status.Code(err)
// 		return nil, status.Errorf(c, "Could not reach ping service: %s", status.Convert(err).Message())
// 	}

// 	log.Print("received upstream pong")
// 	return &pb.Response{
// 		Pong: resp.Pong,
// 	}, nil
// }

// // pingRequest sends a new gRPC ping request to the server configured in the connection.
// func pingRequest(conn *grpc.ClientConn, p *pb.Request) (*pb.Response, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
// 	defer cancel()

// 	client := pb.NewPingServiceClient(conn)
// 	return client.Send(ctx, p)
// }

// // [END run_grpc_request]
// // [END cloudrun_grpc_request]

// // PingRequest creates a new gRPC request to the upstream ping gRPC service.
// func PingRequest(conn *grpc.ClientConn, p *pb.Request, url string, authenticated bool) (*pb.Response, error) {
// 	if authenticated {
// 		return pingRequestWithAuth(conn, p, url)
// 	}
// 	return pingRequest(conn, p)
// }
