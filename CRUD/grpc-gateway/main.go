package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	rd "crud/common-libs/redis"
	shared "crud/common-libs/shared"
	kafka "crud/grpc-gateway/kafka_producer"
	pb "crud/grpc-gateway/proto/gate"
	auth "crud/grpc-gateway/services_logic/authorization"
	notes "crud/grpc-gateway/services_logic/notes"
	reg "crud/grpc-gateway/services_logic/registration"
	tasks "crud/grpc-gateway/services_logic/tasks"
)

type GateServer struct {
	pb.UnimplementedGateServer
	reg.RegistrationServer
	auth.AuthServer
	tasks.TasksServer
	notes.NotesServer
}

func (s *GateServer) GetTaskStatus(ctx context.Context, req *pb.TaskRequest) (*pb.TaskResponse, error) {
	err := shared.ParseKeyToContext(&ctx, req.Taskid)
	if err != nil {
		return nil, err
	}

	if ctx.Value(shared.PrefixKey).(string) == "getAuthorization_task" {
		err = auth.CheckAuthorization(ctx) checkAuth(ctx)
		if err != nil {
			return nil, err
		}
	}

	return s.TasksServer.GetTaskStatus(ctx, req, "anonymous")
}

func (s *GateServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.TaskResponse, error) {
	return s.RegistrationServer.Register(ctx, req)
}

func (s *GateServer) GetAuthorization(ctx context.Context, req *pb.AuthRequest) (*pb.TaskResponse, error) {
	return s.AuthServer.GetAuthorization(ctx, req)
}

func (s *GateServer) CreateNote(ctx context.Context, req *pb.Note) (*pb.TaskResponse, error) {
	err := auth.CheckAuthorization(&ctx)
	if err != nil {
		return nil, err
	}

	return s.NotesServer.CreateNote(ctx, req)
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)

	rd.Init()
	defer rd.Client.Close()

	kafka.Init()
	defer kafka.KafkaProducer.Close()

	if err := auth.InitAuthGateClient(); err != nil {
		log.Fatalf("Ошибка запуска grpc-client: %v", err)
	}

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
