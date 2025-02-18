package notes

import (
	"context"
	"crud/common-libs/shared"
	kp "crud/grpc-gateway/kafka"
	pb "crud/grpc-gateway/proto/gate"
	"encoding/json"
	"log"
	"time"

	rd "crud/common-libs/redis"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type NotesServer struct{}

func createTask(ctx context.Context, req *pb.Note) (string, error) {
	taskId := "createNote_task:" + uuid.New().String()

	taskStatus := shared.CreateNoteStatus{
		BaseTaskStatus: shared.BaseTaskStatus{
			Result: "pending",
		},
	}

	json_b, err := json.Marshal(taskStatus)
	if err != nil {
		return "", err
	}

	if err := rd.Client.Set(ctx, taskId, json_b, time.Hour*1).Err(); err != nil {
		return "", err
	}

	data := shared.CreateNoteData{
		BaseTaskData: shared.BaseTaskData{
			Login:  req.Login,
			TaskId: taskId,
		},
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Failed to marshal data: %v", err)
	}

	go func() {
		kp.KafkaProducer.Produce("create_note", kafka.Message{
			Key:   []byte(req.Login),
			Value: jsonData,
		})
	}()
}

func Create(ctx context.Context, req *pb.Note) {

}
