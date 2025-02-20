package notes

import (
	"context"
	"crud/common-libs/shared"
	kp "crud/grpc-gateway/kafka_producer"
	pb "crud/grpc-gateway/proto/gate"
	"encoding/json"
	"fmt"
	"log"
	"time"

	rd "crud/common-libs/redis"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

func CreateNote(ctx context.Context, req *pb.Note, username *string) (string, error) {
	taskId := fmt.Sprintf("%s:createNote_task:%s", *username, uuid.New().String())

	taskStatus := shared.CreateNoteStatus{
		BaseTaskStatus: shared.BaseTaskStatus{
			Result: "pending",
		},
	}

	if err := rd.PushStatusIntoRedis(ctx, taskId, taskStatus, time.Hour); err != nil {
		return "", err
	}

	data := shared.CreateNoteData{
		BaseTaskData: shared.BaseTaskData{
			Login:  *username,
			TaskId: taskId,
		},
		Note: shared.Note{
			Title:   req.Title,
			Content: req.Content,
		},
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Failed to marshal data: %v", err)
	}

	go func() {
		kp.KafkaProducer.Produce("create_note", kafka.Message{
			Key:   []byte(taskId),
			Value: jsonData,
		})
	}()

	return taskId, nil
}
