package registration

import (
	"context"
	kp "crud/grpc-gateway/kafka_producer"
	pb "crud/grpc-gateway/proto/gate"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	rd "crud/common-libs/redis"
	shared "crud/common-libs/shared"
)

type RegistrationServer struct{}

func setTaskStatus(ctx context.Context, taskId string, taskStatus interface{}) error {
	json_b, err := json.Marshal(taskStatus)
	if err != nil {
		return err
	}

	if err := rd.Client.Set(ctx, taskId, json_b, time.Hour*1).Err(); err != nil {
		return err
	}

	return nil
}

func createTask(ctx context.Context, req *pb.RegisterRequest, username string) (string, error) {
	taskId := fmt.Sprintf("%s:registration_task:%s", username, uuid.New().String())

	taskStatus := shared.AuthorizationCheckStatus{
		BaseTaskStatus: shared.BaseTaskStatus{
			Result: "pending",
		},
	}

	if err := setTaskStatus(ctx, taskId, taskStatus); err != nil {
		return "", err
	}

	data := shared.RegistrationData{
		BaseTaskData: shared.BaseTaskData{
			Login:  req.Login,
			TaskId: taskId,
		},
		Passhash: req.Passhash,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Failed to marshal data: %v", err)
	}

	go func() {
		kp.KafkaProducer.Produce("registrations", kafka.Message{
			Key:   []byte(req.Login),
			Value: jsonData,
		})
	}()

	return taskId, nil
}

func (s *RegistrationServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.TaskResponse, error) {
	log.Println("New register request!")

	taskId, err := createTask(ctx, req, "anonymous")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal server error: %v", err)
	}

	return &pb.TaskResponse{
		Status: "created",
		Info:   taskId,
	}, nil
}
