package main

import (
	"context"
	"log"
	"net"

	pb "grpc_demo_client/proto"

	"google.golang.org/grpc"
)

type chatServer struct {
	pb.UnimplementedMessagerServer
}

func (s *gateServer) Auth(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	verdict := false
	var message string

	if req.Name == "Кирилл" {
		verdict = true
		message = "Здарова "
	} else {
		message = "Ну ты что, всем дурак? "
	}

	return &pb.HelloResponse{Message: message + req.Name, Pass: verdict}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGateServer(s, &gateServer{})
	log.Println("Сервер запущен на :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Ошибка: %v", err)
	}
}
