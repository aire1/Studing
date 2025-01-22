package main

import (
	"context"
	"fmt"
	pb "grpc_demo_client/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewMessagerClient(conn)

	md := metadata.Pairs("user", "client_1")
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	stream, err := client.Chat(ctx)
	if err != nil {
		log.Fatalf("Failed to create stream: %v", err)
	}

	go func() {
		for {
			msg, err := stream.Recv()
			if err != nil {
				log.Fatalf("Failed to receive message: %v", err)
			}
			log.Printf("%s said: %s", msg.User, msg.Text)
		}
	}()

	go func() {
		for {
			var input string
			_, err = fmt.Scan(&input)
			if err != nil {
				log.Fatalf("Failed to fetch user input: %v", err)
				continue
			}

			msg := &pb.ChatMessage{
				User: "client_1",
				Text: input,
			}

			if err := stream.Send(msg); err != nil {
				log.Fatalf("Failed to send message: %v", err)
			}
			log.Printf("Sent message: %s", msg.Text)
		}
	}()

	<-make(chan struct{})
}
