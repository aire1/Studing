package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	rd "crud/common-libs/redis"
	auth "crud/grpc-gateway/authorization"
	kafka "crud/grpc-gateway/kafka"
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
	res, err := auth.CheckAuthorization(ctx)
	if err != nil {
		return nil, err
	} else if !res {
		return nil, status.Errorf(codes.Unauthenticated, "wrong jwt token")
	}

	return s.TasksServer.GetTaskStatus(ctx, req)
}

func (s *GateServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.TaskResponse, error) {
	return s.RegistrationServer.Register(ctx, req)
}

func (s *GateServer) GetAuthorization(ctx context.Context, req *pb.AuthRequest) (*pb.TaskResponse, error) {
	return s.AuthServer.GetAuthorization(ctx, req)
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)

	rd.Init()
	defer rd.Client.Close()

	kafka.Init()
	defer kafka.KafkaProducer.Close()

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
