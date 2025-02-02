package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "crud/grpc-gateway/proto"

	auth "crud/grpc-gateway/auth"

	rd "crud/grpc-gateway/common-libs/redis"
)

func main() {
	rd.Init()
	defer rd.Client.Close()

	lis, err := net.Listen("tcp", ":50050")
	if err != nil {
		log.Fatalf("Ошибка запуска grpc-gate: %v", err)
	}
	s := grpc.NewServer()

	log.Println("grpc-gate запущен на :50050")
	pb.RegisterGateServer(s, &auth.GateServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Ошибка запуска grpc-gate: %v", err)
	}
}
