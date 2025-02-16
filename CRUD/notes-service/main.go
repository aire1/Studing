package main

import (
	"context"
	pg "crud/common-libs/postgres"
	rd "crud/common-libs/redis"
	"crud/common-libs/shared"
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
		GroupID:        "Auth-service",
		MinBytes:       10e3,
		MaxBytes:       10e6,
		MaxWait:        500 * time.Millisecond,
		CommitInterval: time.Second,
	})
	defer reader_create.Close()

	go func(reader *kafka.Reader) {
		var data shared.CreateTaskData
		for {
			message, err := reader.FetchMessage(context.Background())
			if err != nil {
				log.Fatal(err)
			}

			err = json.Unmarshal(message.Value, &data)
			if err != nil {
				log.Printf("failed to unmarshal message: %v", err)
				continue
			}

			go func(data shared.CreateTaskData) {
				status := shared.CreateTaskStatus{
					BaseTaskStatus: shared.BaseTaskStatus{
						Result: "success",
						Info:   "",
					},
				}

				//todo: реализовать функцию добавления заметки в БД

			}(data)
		}
	}(reader_create)

}
