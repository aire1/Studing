package authorization

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

type AuthServer struct{}

type AuthData struct {
	Login    string `json:"login"`
	Passhash string `json:"passhash"`
	Taskid   string `json:"taskid"`
}

func (s *AuthServer) Authorization(ctx context.Context, req *pb.AuthRequest) (*pb.TaskIdResponse, error) {
	log.Println("New auth request!")

	taskId := uuid.New().String()
	err := rd.Client.Set(ctx, taskId, "pending", time.Hour*1).Err()
	if err != nil {
		return &pb.TaskIdResponse{
			Message: "internal error",
		}, err
	}

	data := AuthData{
		Login:    req.Login,
		Passhash: req.Passhash,
		Taskid:   taskId,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Failed to marshal data: %v", err)
	}

	kp.Produce("authorizations", kafka.Message{
		Key:   []byte(req.Login),
		Value: jsonData,
	})

	return &pb.TaskIdResponse{
		Message: taskId,
	}, nil //возвращаем taskId задачи клиенту
}
