package tasks

import (
	"context"
	pb "crud/grpc-gateway/proto/gate"
	"fmt"
	"log"

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

func setTaskResponse(ctx context.Context, taskId string, v interface{}) (*pb.TaskResponse, error) {
	var pbResponse pb.TaskResponse

	err := shared.ParseKeyToContext(&ctx, taskId)
	if err != nil {
		return nil, err
	}

	var prefix = ctx.Value(shared.PrefixKey).(string)

	switch {
	case prefix == shared.RegistrationStatusPrefix:
		pbResponse.Status = v.(*shared.RegistrationStatus).Result
		pbResponse.Info = v.(*shared.RegistrationStatus).Info
	case prefix == shared.AuthorizationGetStatusPrefix:
		pbResponse.Status = v.(*shared.AuthorizationGetStatus).Result
		pbResponse.Info = v.(*shared.AuthorizationGetStatus).Info
	case prefix == shared.AuthorizationCheckStatusPrefix:
		pbResponse.Status = v.(*shared.AuthorizationCheckStatus).Result
		pbResponse.Info = v.(*shared.AuthorizationCheckStatus).Info
	case prefix == shared.CreateNoteStatusPrefix:
		pbResponse.Status = v.(*shared.CreateNoteStatus).Result
		pbResponse.Info = v.(*shared.CreateNoteStatus).Info
	case prefix == shared.GetNoteStatusPrefix:
		pbResponse.Status = v.(*shared.GetNoteStatus).Result
		pbResponse.Info = v.(*shared.GetNoteStatus).Info
		notes := v.(*shared.GetNoteStatus).Notes
		for _, note := range notes {
			if pbResponse.Data == nil {
				pbResponse.Data = &pb.TaskResponse_NoteResponse{
					NoteResponse: &pb.NoteResponse{
						Note: []*pb.Note{},
					},
				}
			}
			pbResponse.Data.(*pb.TaskResponse_NoteResponse).NoteResponse.Note = append(pbResponse.Data.(*pb.TaskResponse_NoteResponse).NoteResponse.Note, &pb.Note{
				Id:        fmt.Sprint(note.Id),
				Title:     note.Title,
				Content:   note.Content,
				CreatedAt: note.CreatedAt,
				UpdatedAt: note.UpdatedAt,
			})
		}

	default:
		return nil, fmt.Errorf("unknown task prefix")
	}

	return &pbResponse, nil
}

func (s *TasksServer) GetTaskStatus(ctx context.Context, req *pb.TaskRequest) (*pb.TaskResponse, error) {
	log.Println("New task status request!")

	v, err := shared.GetTaskFromRedis(rd.Client, ctx, req.Taskid)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}

	return setTaskResponse(ctx, req.Taskid, v)
}
