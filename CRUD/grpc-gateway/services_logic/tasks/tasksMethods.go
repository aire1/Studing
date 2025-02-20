package tasks

import (
	"context"
	pb "crud/grpc-gateway/proto/gate"
	"fmt"
	"log"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	rd "crud/common-libs/redis"
	shared "crud/common-libs/shared"
)

type TasksServer struct{}

type TaskData interface {
	SetTaskID(taskId string)
	SetLogin(username string)
}

func setTaskResponse(taskId string, v interface{}) (*pb.TaskResponse, error) {
	var pbResponse pb.TaskResponse

	switch {
	case strings.HasPrefix(taskId, shared.RegistrationStatusPrefix):
		pbResponse.Status = v.(*shared.RegistrationStatus).Result
		pbResponse.Info = v.(*shared.RegistrationStatus).Info
	case strings.HasPrefix(taskId, shared.AuthorizationGetStatusPrefix):
		pbResponse.Status = v.(*shared.AuthorizationGetStatus).Result
		pbResponse.Info = v.(*shared.AuthorizationGetStatus).Info
	case strings.HasPrefix(taskId, shared.AuthorizationCheckStatusPrefix):
		pbResponse.Status = v.(*shared.AuthorizationCheckStatus).Result
		pbResponse.Info = v.(*shared.AuthorizationCheckStatus).Info
	case strings.HasPrefix(taskId, shared.CreateNoteStatusPrefix):
		pbResponse.Status = v.(*shared.CreateNoteStatus).Result
		pbResponse.Info = v.(*shared.CreateNoteStatus).Info
	default:
		return nil, fmt.Errorf("unknown task prefix")
	}

	return &pbResponse, nil
}

func (s *TasksServer) GetTaskStatus(ctx context.Context, req *pb.TaskRequest, username string) (*pb.TaskResponse, error) {
	log.Println("New task status request!")

	v, err := shared.GetTaskFromRedis(rd.Client, ctx, req.Taskid)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}

	prefix, err := shared.GetStatusPrefix(req.Taskid)
	if err != nil {
		return nil, err
	} else if prefix == "" {
		return nil, fmt.Errorf("prefix is nil")
	}

	return setTaskResponse(prefix, v)
}
