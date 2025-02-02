package auth

import (
	"context"
	kp "crud/grpc-gateway/kafka"
	pb "crud/grpc-gateway/proto"
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"

	rd "crud/grpc-gateway/common-libs/redis"
)

type GateServer struct {
	pb.UnimplementedGateServer
	//mu sync.Mutex
}

type RegistrationData struct {
	Login    string `json:"login"`
	Passhash string `json:"passhash"`
	Taskid   string `json:"taskid"`
}

func (s *GateServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	log.Println("New req!")

	taskId := uuid.New().String()
	err := rd.Client.Set(ctx, taskId, "pending", time.Hour*1).Err()
	if err != nil {
		return &pb.RegisterResponse{
			Message: "internal error",
		}, err
	}

	data := RegistrationData{
		Login:    req.Login,
		Passhash: req.Passhash,
		Taskid:   taskId,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Failed to marshal data: %v", err)
	}

	kp.Produce("registrations", kafka.Message{
		Key:   []byte(req.Login),
		Value: jsonData,
	})

	return &pb.RegisterResponse{
		Message: taskId,
	}, nil //возвращаем taskId задачи клиенту
}
