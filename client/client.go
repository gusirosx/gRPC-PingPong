// Package client is a CLI to make requests to the grpc-ping service.
package main

import (
	"context"
	"flag"
	pb "gRPC-Ping/proto"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
)

var (
	message      = flag.String("message", "Hi there", "The body of the content sent to server")
	sendUpstream = flag.Bool("relay", false, "Direct ping to relay the request to a ping-upstream service [false]")
)

func main() {
	// Set up a connection to the server.
	conn, err := Connection()
	if err != nil {
		log.Printf("failed to dial server %s: %v", *serverAddr, err)
	}
	defer conn.Close()

	// if err := streamTime(client, *duration); err != nil {
	// 	log.Fatal(err)
	// }

	client := pb.NewPingServiceClient(conn)
	send(client)
}

func send(client pb.PingServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	var resp *pb.Response
	var err error
	if *sendUpstream {
		resp, err = client.SendUpstream(ctx, &pb.Request{
			Message: *message,
		})
	} else {
		resp, err = client.Send(ctx, &pb.Request{
			Message: *message,
		})
	}

	if err != nil {
		log.Fatalf("Error while executing Send: %v", err)
	}

	respMessage := resp.Pong.GetMessage()
	timestamp := ptypes.TimestampString(resp.Pong.GetReceivedOn())
	log.Println("Unary Request/Unary Response")
	log.Printf("  Sent Ping: %s", *message)
	log.Printf("  Received:\n    Pong: %s\n    Server Time: %s", respMessage, timestamp)
}
