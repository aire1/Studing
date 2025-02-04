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

	rd "crud/common-libs/redis"

	shared "crud/common-libs/shared"
)

type AuthServer struct{}

func (s *AuthServer) Authorization(ctx context.Context, req *pb.AuthRequest) (*pb.TaskIdResponse, error) {
	log.Println("New auth request!")

	taskId := "getAuthorization_task:" + uuid.New().String()

	taskStatus := shared.AuthorizationCheckStatus{
		Result: "pending",
	}

	json_b, err := json.Marshal(taskStatus)
	if err != nil {
		return nil, err
	}

	if err := rd.Client.Set(ctx, taskId, json_b, time.Hour*1).Err(); err != nil {
		return nil, err
	}

	if err != nil {
		return &pb.TaskIdResponse{
			Message: "internal error",
		}, err
	}

	data := shared.AuthorizationGetData{
		Login:    req.Login,
		Passhash: req.Passhash,
		Taskid:   taskId,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Failed to marshal data: %v", err)
	}

	kp.Produce("get_authorizations", kafka.Message{
		Key:   []byte(req.Login),
		Value: jsonData,
	})

	return &pb.TaskIdResponse{
		Message: taskId,
	}, nil //возвращаем taskId задачи клиенту
}
