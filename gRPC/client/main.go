package main

import (
	"context"
	pb "grpc_demo_client/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Не удалось подключиться: %v", err)
	}
	defer conn.Close()

	client := pb.NewGateClient(conn)
	resp, err := client.Auth(context.Background(), &pb.HelloRequest{Name: "Кирилл"})

	if err != nil {
		log.Fatalf("Ошибка вызова RPC: %v", err)
	}

	log.Printf("Ответ сервера: %s", resp.Message)

}
