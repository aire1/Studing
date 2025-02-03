package tasks

import (
	"context"
	pb "crud/grpc-gateway/proto"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	rd "crud/grpc-gateway/common-libs/redis"
)

type TasksServer struct{}

type RegistrationData struct {
	Login    string `json:"login"`
	Passhash string `json:"passhash"`
	Taskid   string `json:"taskid"`
}

func (s *TasksServer) GetTaskStatus(ctx context.Context, req *pb.TaskRequest) (*pb.TaskResponse, error) {
	log.Println("New task status request!")

	res, err := rd.Client.Get(ctx, req.Taskid).Result()
	if err != nil {
		return &pb.TaskResponse{}, status.Errorf(codes.Internal, "error getting task from redis")
	}

	return &pb.TaskResponse{
		Status: res,
	}, nil

}
