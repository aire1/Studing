package main

import (
	"context"
	pg "crud/common-libs/postgres"
	rd "crud/common-libs/redis"
	"crud/common-libs/shared"
	methods "crud/notes-service/methods"
	"encoding/json"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	rd.Init()
	defer rd.Client.Close()
	pg.CreatePool()
	defer pg.Pool.Close()

	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)

	reader_create := kafka.NewReader(kafka.ReaderConfig{
		//Brokers:        []string{"kafka1:9092", "kafka2:9092", "kafka3:9092"},
		Brokers:        []string{"localhost:19092", "localhost:19094", "localhost:19096"},
		Topic:          "create_note",
		GroupID:        "Notes-service",
		MinBytes:       10e3,
		MaxBytes:       10e6,
		MaxWait:        500 * time.Millisecond,
		CommitInterval: time.Second,
	})
	defer reader_create.Close()

	reader_get := kafka.NewReader(kafka.ReaderConfig{
		//Brokers:        []string{"kafka1:9092", "kafka2:9092", "kafka3:9092"},
		Brokers:        []string{"localhost:19092", "localhost:19094", "localhost:19096"},
		Topic:          "get_note",
		GroupID:        "Notes-service",
		MinBytes:       10e3,
		MaxBytes:       10e6,
		MaxWait:        500 * time.Millisecond,
		CommitInterval: time.Second,
	})
	defer reader_get.Close()

	reader_update := kafka.NewReader(kafka.ReaderConfig{
		//Brokers:        []string{"kafka1:9092", "kafka2:9092", "kafka3:9092"},
		Brokers:        []string{"localhost:19092", "localhost:19094", "localhost:19096"},
		Topic:          "update_note",
		GroupID:        "Notes-service",
		MinBytes:       10e3,
		MaxBytes:       10e6,
		MaxWait:        500 * time.Millisecond,
		CommitInterval: time.Second,
	})
	defer reader_update.Close()

	delete_update := kafka.NewReader(kafka.ReaderConfig{
		//Brokers:        []string{"kafka1:9092", "kafka2:9092", "kafka3:9092"},
		Brokers:        []string{"localhost:19092", "localhost:19094", "localhost:19096"},
		Topic:          "delete_note",
		GroupID:        "Notes-service",
		MinBytes:       10e3,
		MaxBytes:       10e6,
		MaxWait:        500 * time.Millisecond,
		CommitInterval: time.Second,
	})
	defer delete_update.Close()

	log.Println("Подключился к Kafka")

	go processMessages(reader_create, handleCreateNote)
	go processMessages(reader_get, handleGetNote)
	go processMessages(reader_update, handleUpdateNote)
	go processMessages(delete_update, handleDeleteNote)

	select {}
}

func processMessages(reader *kafka.Reader, handler func(context.Context, []byte) error) {
	for {
		message, err := reader.FetchMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		err = handler(context.Background(), message.Value)
		if err != nil {
			log.Printf("failed to handle message: %v", err)
			continue
		}

		err = reader.CommitMessages(context.Background(), message)
		if err != nil {
			log.Printf("failed to commit message: %v", err)
		}
	}
}

func handleCreateNote(ctx context.Context, message []byte) error {
	var data shared.CreateNoteData
	err := json.Unmarshal(message, &data)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, time.Hour*24)
	defer cancel()

	status := shared.CreateNoteStatus{
		BaseTaskStatus: shared.BaseTaskStatus{
			Result: "success",
			Info:   "",
		},
	}

	id, err := methods.Create(ctx, data)
	if err != nil {
		status.Result = "error"
		status.Info = err.Error()
	} else {
		status.Result = "success"
		status.Info = id
	}

	err = rd.PushStatusIntoRedis(ctx, data.TaskId, status, time.Hour)
	if err != nil {
		log.Printf("failed to push status into redis: %v", err)
	}

	return nil
}

func handleGetNote(ctx context.Context, message []byte) error {
	var data shared.GetNoteData
	err := json.Unmarshal(message, &data)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, time.Hour*24)
	defer cancel()

	status := shared.GetNoteStatus{
		BaseTaskStatus: shared.BaseTaskStatus{
			Result: "success",
			Info:   "",
		},
	}

	notes, err := methods.Get(ctx, data)
	if err != nil {
		status.Result = "error"
		status.Info = err.Error()
	} else if notes == nil {
		status.Info = "no rows in result set"
	} else {
		status.Result = "success"
		status.Notes = *notes
	}

	err = rd.PushStatusIntoRedis(ctx, data.TaskId, status, time.Hour)
	if err != nil {
		log.Printf("failed to push status into redis: %v", err)
	}

	return nil
}

func handleDeleteNote(ctx context.Context, message []byte) error {
	var data shared.DeleteNoteData
	err := json.Unmarshal(message, &data)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, time.Hour*24)
	defer cancel()

	status := shared.DeleteNoteStatus{
		BaseTaskStatus: shared.BaseTaskStatus{
			Result: "success",
			Info:   "",
		},
	}

	err = methods.Delete(ctx, data)
	if err != nil {
		status.Result = "error"
		status.Info = err.Error()
	} else {
		status.Result = "success"
	}

	err = rd.PushStatusIntoRedis(ctx, data.TaskId, status, time.Hour)
	if err != nil {
		log.Printf("failed to push status into redis: %v", err)
	}

	return nil
}

func handleUpdateNote(ctx context.Context, message []byte) error {
	var data shared.UpdateNoteData
	err := json.Unmarshal(message, &data)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, time.Hour*24)
	defer cancel()

	status := shared.UpdateNoteStatus{
		BaseTaskStatus: shared.BaseTaskStatus{
			Result: "success",
			Info:   "",
		},
	}

	err = methods.Update(ctx, data)
	if err != nil {
		status.Result = "error"
		status.Info = err.Error()
	} else {
		status.Result = "success"
	}

	err = rd.PushStatusIntoRedis(ctx, data.TaskId, status, time.Hour)
	if err != nil {
		log.Printf("failed to push status into redis: %v", err)
	}

	return nil
}
