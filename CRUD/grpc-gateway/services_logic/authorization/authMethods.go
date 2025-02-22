package authorization

import (
	"context"
	kp "crud/grpc-gateway/kafka_producer"
	pb_auth_gate "crud/grpc-gateway/proto/auth_gate"
	pb_main_gate "crud/grpc-gateway/proto/gate"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	rd "crud/common-libs/redis"
	shared "crud/common-libs/shared"
)

var (
	AuthClient pb_auth_gate.AuthGateClient
)

type AuthServer struct{}

func InitAuthGateClient() error {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	AuthClient = pb_auth_gate.NewAuthGateClient(conn)

	return nil
}

func createTaskGetAuth(ctx context.Context, req *pb_main_gate.AuthRequest, username string) (string, error) {
	taskId := fmt.Sprintf("%s:getAuthorization_task:%s", username, uuid.New().String())

	taskStatus := shared.AuthorizationCheckStatus{
		BaseTaskStatus: shared.BaseTaskStatus{
			Result: "pending",
		},
	}

	if err := rd.PushStatusIntoRedis(ctx, taskId, taskStatus, time.Hour); err != nil {
		return "", err
	}

	data := shared.AuthorizationGetData{
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

	kp.KafkaProducer.Produce("get_authorizations", kafka.Message{
		Key:   []byte(taskId),
		Value: jsonData,
	})

	return taskId, nil
}

func (s *AuthServer) GetAuthorization(ctx context.Context, req *pb_main_gate.AuthRequest) (*pb_main_gate.TaskResponse, error) {
	log.Println("New auth request!")

	taskId, err := createTaskGetAuth(ctx, req, req.Login)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal server error: %v", err)
	}

	return &pb_main_gate.TaskResponse{
		Status: "created",
		Info:   taskId,
	}, nil
}

func CheckAuthorization(ctx context.Context) (string, error) {
	log.Println("CheckAuthorization -> metadata start")
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "missing metadata")
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return "", status.Errorf(codes.Unauthenticated, "missing authorization token")
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := AuthClient.CheckAuthorization(
		ctx,
		&pb_auth_gate.AuthCheckRequest{
			JwtToken: strings.Join(authHeader, ""),
		},
	)
	if err != nil {
		return "", err
	}

	return res.Username, nil
}
