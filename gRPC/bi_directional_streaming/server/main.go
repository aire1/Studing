package main

import (
	"log"
	"net"
	"sync"

	pb "grpc_demo_client/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type chatServer struct {
	pb.UnimplementedMessagerServer
	clients map[string]pb.Messager_ChatServer
	mu      sync.Mutex
}

func (s *chatServer) Chat(stream pb.Messager_ChatServer) error {
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return status.Error(codes.Unauthenticated, "metadata not found")
	}
	userValues := md.Get("user")
	if len(userValues) == 0 {
		return status.Error(codes.Unauthenticated, "user not found")
	}

	s.mu.Lock()
	s.clients[userValues[0]] = stream
	s.mu.Unlock()

	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Printf("Client disconnected: %v", err)
			s.mu.Lock()
			delete(s.clients, msg.User)
			s.mu.Unlock()
			return err
		}
		log.Printf("Received message from %s: %s", msg.User, msg.Text)

		s.mu.Lock()
		for user, client := range s.clients {
			if user != msg.User {
				if err := client.Send(msg); err != nil {
					log.Printf("Failed to send message to %s: %v", user, err)
				}
			}
		}
		s.mu.Unlock()
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterMessagerServer(grpcServer, &chatServer{clients: make(map[string]pb.Messager_ChatServer)})
	log.Println("Server is running on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
