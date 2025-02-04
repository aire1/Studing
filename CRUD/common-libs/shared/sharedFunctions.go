package shared

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "crud/grpc-gateway/proto"
)

const (
	RegistrationStatusPrefix       = "registration_task:"
	AuthorizationGetStatusPrefix   = "getAuthorization_task:"
	AuthorizationCheckStatusPrefix = "authorization_check_status:"
)

type RedisSerializable interface {
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}

func (r *RegistrationData) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func (r *RegistrationData) Unmarshal(data []byte) error {
	return json.Unmarshal(data, r)
}

func (r *RegistrationStatus) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func (r *RegistrationStatus) Unmarshal(data []byte) error {
	return json.Unmarshal(data, r)
}

func (r *AuthorizationGetStatus) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func (r *AuthorizationGetStatus) Unmarshal(data []byte) error {
	return json.Unmarshal(data, r)
}

func (r *AuthorizationCheckStatus) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func (r *AuthorizationCheckStatus) Unmarshal(data []byte) error {
	return json.Unmarshal(data, r)
}

func GetTaskFromRedis(client *redis.Client, ctx context.Context, key string) (any, error) {
	var v RedisSerializable

	switch {
	case strings.HasPrefix(key, RegistrationStatusPrefix):
		v = &RegistrationStatus{}
	case strings.HasPrefix(key, AuthorizationGetStatusPrefix):
		v = &AuthorizationGetStatus{}
	case strings.HasPrefix(key, AuthorizationCheckStatusPrefix):
		v = &AuthorizationCheckStatus{}
	}

	if err := UnmarshalFromRedis(client, ctx, key, v); err != nil {
		return nil, err
	}

	return v, nil
}

func UnmarshalFromRedis(client *redis.Client, ctx context.Context, key string, v RedisSerializable) error {
	json_b, err := client.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	if err := v.Unmarshal([]byte(json_b)); err != nil {
		return err
	}

	return nil
}

func WaitForCompleteTask(client *redis.Client, taskId string, ctx context.Context) (*pb.TaskResponse, error) {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return &pb.TaskResponse{Status: "timeout", Info: taskId}, nil
		case <-ticker.C:
			v, err := GetTaskFromRedis(client, ctx, taskId)
			if err != nil {
				if err == redis.Nil {
					continue
				}
				return nil, status.Errorf(codes.Internal, "error getting task from redis: %v", err)
			}

			var pbResponse pb.TaskResponse
			switch task := v.(type) {
			case *RegistrationStatus:
				pbResponse.Status = task.Result
				pbResponse.Info = task.Info
			case *AuthorizationGetStatus:
				pbResponse.Status = task.Result
				pbResponse.Info = task.Info
			case *AuthorizationCheckStatus:
				pbResponse.Status = task.Result
				pbResponse.Info = task.Info
			default:
				return nil, status.Errorf(codes.Internal, "unknown task type")
			}

			if pbResponse.Status == "pending" {
				continue
			}

			return &pbResponse, nil
		}
	}
}

// Добавьте другие структуры и их методы Marshal/Unmarshal
