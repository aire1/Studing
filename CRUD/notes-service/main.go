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

	log.Println("Подключился к Kafka")

	go func(reader *kafka.Reader) {
		var data shared.CreateNoteData
		for {
			ctx := context.Background()
			message, err := reader.FetchMessage(ctx)

			if err != nil {
				log.Fatal(err)
			}

			err = json.Unmarshal(message.Value, &data)
			if err != nil {
				log.Printf("failed to unmarshal message: %v", err)
				continue
			}

			go func(data shared.CreateNoteData) {
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
			}(data)

			log.Printf("Получено сообщение: %v", data)
			err = reader.CommitMessages(ctx, message)
			if err != nil {
				log.Printf("failed to commit message: %v", err)
			}
		}
	}(reader_create)

	go func(reader *kafka.Reader) {
		var data shared.GetNoteData
		for {
			message, err := reader.FetchMessage(context.Background())

			// err = reader.CommitMessages(ctx, message)
			// if err != nil {
			// 	log.Printf("failed to commit message: %v", err)
			// }

			// continue

			if err != nil {
				log.Fatal(err)
			}

			err = json.Unmarshal(message.Value, &data)
			if err != nil {
				log.Printf("failed to unmarshal message: %v", err)
				continue
			}

			go func(data shared.GetNoteData) {
				ctx, cancel := context.WithTimeout(context.Background(), time.Hour*24)

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

				cancel()
			}(data)

			log.Printf("Получено сообщение: %v", data)
			err = reader.CommitMessages(context.Background(), message)
			if err != nil {
				log.Printf("failed to commit message: %v", err)
			}
		}
	}(reader_get)

	select {}
}
