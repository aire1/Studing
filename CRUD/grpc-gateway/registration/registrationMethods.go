package registration

import (
	"context"
	kp "crud/grpc-gateway/kafka"
	pb "crud/grpc-gateway/proto"
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	rd "crud/common-libs/redis"
	shared "crud/common-libs/shared"
)

type RegistrationServer struct{}

func createTask(ctx context.Context, req *pb.RegisterRequest) (string, error) {
	taskId := "registration_task:" + uuid.New().String()

	taskStatus := shared.AuthorizationCheckStatus{
		Result: "pending",
	}

	json_b, err := json.Marshal(taskStatus)
	if err != nil {
		return "", err
	}

	if err := rd.Client.Set(ctx, taskId, json_b, time.Hour*1).Err(); err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}

	data := shared.RegistrationData{
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

	return taskId, nil
}

func (s *RegistrationServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.TaskResponse, error) {
	log.Println("New register request!")

	taskId, err := createTask(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal server error: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	pbRes, err := shared.WaitForCompleteTask(rd.Client, taskId, ctx)
	if err != nil {
		if strings.Contains(err.Error(), "context deadline exceeded") {
			return &pb.TaskResponse{
				Status: "timeout",
				Info:   taskId,
			}, nil
		}
		return nil, err
	}
	return pbRes, nil
}
