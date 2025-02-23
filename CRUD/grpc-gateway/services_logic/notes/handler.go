package notes

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	rd "crud/common-libs/redis"
	shared "crud/common-libs/shared"
	kp "crud/grpc-gateway/kafka_producer"
	pb "crud/grpc-gateway/proto/gate"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NotesServer struct{}

func (s *NotesServer) CreateNote(ctx context.Context, req *pb.Note) (*pb.TaskResponse, error) {
	username := ctx.Value("username").(string)
	uid := ctx.Value("uid").(string)

	data := shared.CreateNoteData{}

	if req.Title == "" {
		return nil, status.Error(codes.InvalidArgument, "title is empty")
	} else if req.Content == "" {
		return nil, status.Error(codes.InvalidArgument, "content is empty")
	}

	taskId := fmt.Sprintf("%s:createNote_task:%s", username, uuid.New().String())
	taskStatus := shared.CreateNoteStatus{
		BaseTaskStatus: shared.BaseTaskStatus{
			Result: "pending",
		},
	}

	if err := rd.PushStatusIntoRedis(ctx, taskId, taskStatus, time.Hour); err != nil {
		return nil, err
	}

	data.Title = req.Title
	data.Content = req.Content
	data.Login = username
	data.TaskId = taskId
	data.UserId = uid

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Failed to marshal data: %v", err)
	}

	kp.KafkaProducer.Produce("create_note", kafka.Message{
		Key:   []byte(taskId),
		Value: jsonData,
	})

	return &pb.TaskResponse{
		Status: "created",
		Info:   taskId,
	}, nil
}

func (s *NotesServer) GetNotes(ctx context.Context, req *pb.NoteRequest) (*pb.TaskResponse, error) {
	username := ctx.Value("username").(string)

	data := shared.GetNoteData{}

	if req.Count < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "Count can't be less than 1")
	}

	taskId := fmt.Sprintf("%s:getNote_task:%s", username, uuid.New().String())
	taskStatus := shared.GetNoteStatus{
		BaseTaskStatus: shared.BaseTaskStatus{
			Result: "pending",
		},
	}

	if err := rd.PushStatusIntoRedis(ctx, taskId, taskStatus, time.Hour); err != nil {
		return nil, err
	}

	data.Login = username
	data.TaskId = taskId
	data.Count = int(req.Count)
	data.Offset = int(req.Offset)

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Failed to marshal data: %v", err)
	}

	kp.KafkaProducer.Produce("get_note", kafka.Message{
		Key:   []byte(taskId),
		Value: jsonData,
	})

	return &pb.TaskResponse{
		Status: "created",
		Info:   taskId,
	}, nil
}
