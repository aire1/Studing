package authorization

import (
	"context"
	kp "crud/grpc-gateway/kafka"
	pb "crud/grpc-gateway/proto"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	rd "crud/common-libs/redis"

	shared "crud/common-libs/shared"
)

type AuthServer struct{}

func createTask(ctx context.Context, req *pb.AuthRequest) (string, error) {
	taskId := "getAuthorization_task:" + uuid.New().String()

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

	return taskId, nil
}

func (s *AuthServer) Authorization(ctx context.Context, req *pb.AuthRequest) (*pb.TaskResponse, error) {
	log.Println("New auth request!")

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

	if pbRes.Status == "success" {
		authResId := fmt.Sprintf("session:%s", req.Login)
		jwtToken, err := rd.Client.Get(ctx, authResId).Result()
		if err != nil {
			return pbRes, nil
		}

		md := metadata.Pairs("authorization", jwtToken)
		grpc.SetHeader(ctx, md)
	}

	return pbRes, nil
}
