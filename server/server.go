// Sample grpc-ping acts as an intermediary to the ping service.
package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"net"
// 	"os"
// 	"strings"
// 	"time"

// 	pb "github.com/GoogleCloudPlatform/golang-samples/run/grpc-ping/pkg/api/v1"
// 	"github.com/golang/protobuf/ptypes"
// 	"google.golang.org/grpc/status"

// 	"google.golang.org/grpc"
// )

// // [START cloudrun_grpc_server]
// // [START run_grpc_server]
// func main() {
// 	log.Printf("grpc-ping: starting server...")

// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "8080"
// 		log.Printf("Defaulting to port %s", port)
// 	}

// 	listener, err := net.Listen("tcp", ":"+port)
// 	if err != nil {
// 		log.Fatalf("net.Listen: %v", err)
// 	}

// 	grpcServer := grpc.NewServer()
// 	pb.RegisterPingServiceServer(grpcServer, &pingService{})
// 	if err = grpcServer.Serve(listener); err != nil {
// 		log.Fatal(err)
// 	}
// }

// // [END run_grpc_server]
// // [END cloudrun_grpc_server]

// // conn holds an open connection to the ping service.
// var conn *grpc.ClientConn

// func init() {
// 	if os.Getenv("GRPC_PING_HOST") != "" {
// 		var err error
// 		conn, err = NewConn(os.Getenv("GRPC_PING_HOST"), os.Getenv("GRPC_PING_INSECURE") != "")
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	} else {
// 		log.Println("Starting without support for SendUpstream: configure with 'GRPC_PING_HOST' environment variable. E.g., example.com:443")
// 	}
// }

// // NewConn creates a new gRPC connection.
// // host should be of the form domain:port, e.g., example.com:443
// func NewConn(host string, insecure bool) (*grpc.ClientConn, error) {
// 	var opts []grpc.DialOption

// 	return grpc.Dial(host, opts...)
// }

// type pingService struct {
// 	pb.UnimplementedPingServiceServer
// }

// func (s *pingService) Send(ctx context.Context, req *pb.Request) (*pb.Response, error) {
// 	log.Print("sending ping response")
// 	return &pb.Response{
// 		Pong: &pb.Pong{
// 			Index:      1,
// 			Message:    req.GetMessage(),
// 			ReceivedOn: ptypes.TimestampNow(),
// 		},
// 	}, nil
// }

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
