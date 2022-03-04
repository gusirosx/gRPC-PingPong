// Package client is a CLI to make requests to the grpc-ping service.
package main

import (
	"context"
	"flag"
	pb "gRPC-Ping/proto"
	"log"
	"time"
)

var (
	message = flag.String("message", "Hi there", "The body of the content sent to server")
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

	resp, err := client.Send(ctx, &pb.Request{Message: *message})

	if err != nil {
		log.Fatalf("Error while executing Send: %v", err)
	}

	respMessage := resp.Pong.GetMessage()
	//timestamp := ptypes.TimestampString(resp.Pong.GetReceivedOn())
	timestamp := resp.Pong.GetReceivedOn().AsTime().Format(time.RFC3339)
	//ptypes.TimestampString is deprecated: Call the ts.AsTime method instead, followed by a call to the Format method on the time.Time value.
	log.Println("Unary Request/Unary Response")
	log.Printf("  Sent Ping: %s", *message)
	log.Printf("  Received:\n    Pong: %s\n    Server Time: %s", respMessage, timestamp)
}
