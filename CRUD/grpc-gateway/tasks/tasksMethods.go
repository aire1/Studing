package tasks

import (
	"context"
	pb "crud/grpc-gateway/proto/gate"
	"log"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	rd "crud/common-libs/redis"
	shared "crud/common-libs/shared"
)

type TasksServer struct{}

func (s *TasksServer) GetTaskStatus(ctx context.Context, req *pb.TaskRequest) (*pb.TaskResponse, error) {
	log.Println("New task status request!")

	v, err := shared.GetTaskFromRedis(rd.Client, ctx, req.Taskid)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}

	var pbResponse pb.TaskResponse

	switch {
	case strings.HasPrefix(req.Taskid, shared.RegistrationStatusPrefix):
		pbResponse.Status = v.(*shared.RegistrationStatus).Result
		pbResponse.Info = v.(*shared.RegistrationStatus).Info
	case strings.HasPrefix(req.Taskid, shared.AuthorizationGetStatusPrefix):
		pbResponse.Status = v.(*shared.AuthorizationGetStatus).Result
		pbResponse.Info = v.(*shared.AuthorizationGetStatus).Info
	case strings.HasPrefix(req.Taskid, shared.AuthorizationCheckStatusPrefix):
		pbResponse.Status = v.(*shared.AuthorizationCheckStatus).Result
		pbResponse.Info = v.(*shared.AuthorizationCheckStatus).Info
	}

	return &pbResponse, nil
}
