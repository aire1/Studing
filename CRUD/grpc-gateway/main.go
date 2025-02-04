package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	rd "crud/common-libs/redis"
	auth "crud/grpc-gateway/authorization"
	pb "crud/grpc-gateway/proto"
	reg "crud/grpc-gateway/registration"
	tasks "crud/grpc-gateway/tasks"
)

type GateServer struct {
	pb.UnimplementedGateServer
	reg.RegistrationServer
	auth.AuthServer
	tasks.TasksServer
}

func (s *GateServer) GetTaskStatus(ctx context.Context, req *pb.TaskRequest) (*pb.TaskResponse, error) {
	return s.TasksServer.GetTaskStatus(ctx, req)
}

func (s *GateServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.TaskResponse, error) {
	return s.RegistrationServer.Register(ctx, req)
}

func (s *GateServer) Authorization(ctx context.Context, req *pb.AuthRequest) (*pb.TaskIdResponse, error) {
	return s.AuthServer.Authorization(ctx, req)
}

func main() {
	rd.Init()
	defer rd.Client.Close()

	lis, err := net.Listen("tcp", ":50050")
	if err != nil {
		log.Fatalf("Ошибка запуска grpc-gate: %v", err)
	}
	s := grpc.NewServer()

	log.Println("grpc-gate запущен на :50050")
	pb.RegisterGateServer(s, &GateServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Ошибка запуска grpc-gate: %v", err)
	}
}
