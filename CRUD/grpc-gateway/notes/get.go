package notes

import (
	"context"
	"crud/common-libs/shared"
	kp "crud/grpc-gateway/kafka"
	pb "crud/grpc-gateway/proto/gate"
	"encoding/json"
	"fmt"
	"log"
	"time"

	rd "crud/common-libs/redis"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

func SelectNote(ctx context.Context, req *pb.NoteRequest, username *string) (string, error) {
	taskId := fmt.Sprintf("%s:createNote_task:%s", username, uuid.New().String())

	taskStatus := shared.GetNoteStatus{
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

	data := shared.GetNoteData{
		BaseTaskData: shared.BaseTaskData{
			Login:  *username,
			TaskId: taskId,
		},
		Offset: int(req.Offset),
		Count:  int(req.Count),
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Failed to marshal data: %v", err)
	}

	go func() {
		kp.KafkaProducer.Produce("get_note", kafka.Message{
			Key:   []byte(taskId),
			Value: jsonData,
		})
	}()

	return taskId, nil
}
