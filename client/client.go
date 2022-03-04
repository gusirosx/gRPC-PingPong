// Package client is a CLI to make requests to the grpc-ping service.
package main

import (
	"fmt"
	pb "gRPC-Ping/proto"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set up a connection to the server.
	conn, err := Connection()
	if err != nil {
		log.Printf("failed to dial server %s: %v", *serverAddr, err)
	}
	defer conn.Close()
	client := pb.NewPingServiceClient(conn)
	//send(client)

	// Set up a http server.
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		fmt.Fprintln(ctx.Writer, "Up and running...")
	})

	router.GET("/ping", func(ctx *gin.Context) {
		message := "Ping" //The body of the content sent to server
		// Contact the server and print out its response.
		response, err := client.Send(ctx, &pb.Request{Message: message})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"response": Response{
				Send:         message,
				SendTime:     time.Now(),
				Received:     response.Message,
				ReceivedTime: response.ReceivedOn.AsTime()},
		})
	})
	// Run http server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}

// REST response struct
type Response struct {
	Send         string    `json:"send"`
	SendTime     time.Time `json:"send_on"`
	Received     string    `json:"received"`
	ReceivedTime time.Time `json:"received_on"`
}
