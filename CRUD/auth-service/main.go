package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	pg "crud/auth-service/common-libs/postgres"
	rd "crud/auth-service/common-libs/redis"

	auth "crud/auth-service/authorization"
	reg "crud/auth-service/registration"

	"github.com/segmentio/kafka-go"
)

func pushRedis(ctx context.Context, key string, value interface{}, expiration time.Duration) {
	err := rd.Client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		log.Printf("can't push result into Redis: %v\n", err)
	}
}

func main() {
	rd.Init()
	defer rd.Client.Close()
	pg.CreatePool()
	defer pg.Pool.Close()

	reader_reg := kafka.NewReader(kafka.ReaderConfig{
		//Brokers:        []string{"kafka1:9092", "kafka2:9092", "kafka3:9092"},
		Brokers:        []string{"localhost:19092", "localhost:19094", "localhost:19096"},
		Topic:          "registrations",
		GroupID:        "Auth-service",
		MinBytes:       10e3,
		MaxBytes:       10e6,
		MaxWait:        500 * time.Millisecond,
		CommitInterval: time.Second,
	})
	defer reader_reg.Close()

	reader_auth := kafka.NewReader(kafka.ReaderConfig{
		//Brokers:        []string{"kafka1:9092", "kafka2:9092", "kafka3:9092"},
		Brokers:        []string{"localhost:19092", "localhost:19094", "localhost:19096"},
		Topic:          "authorizations",
		GroupID:        "Auth-service",
		MinBytes:       10e3,
		MaxBytes:       10e6,
		MaxWait:        500 * time.Millisecond,
		CommitInterval: time.Second,
	})
	defer reader_auth.Close()

	log.Println("Подключился к Kafka")

	go func(reader *kafka.Reader) {
		var data auth.AuthData
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

			go func(data auth.AuthData) {
				err = reg.Register(context.Background(), data)

				if err != nil {
					log.Printf("failed to register user: %v", err)
					pushRedis(context.Background(), data.Taskid, "fail", 1*time.Hour)
				} else {
					pushRedis(context.Background(), data.Taskid, "success", 1*time.Hour)
				}
			}(data)

			fmt.Printf("Получено сообщение: %v\n", data)

			err = reader.CommitMessages(context.Background(), message)
			if err != nil {
				log.Printf("failed to commit message: %v", err)
			}
		}
	}(reader_reg)

	go func(reader *kafka.Reader) {
		var data auth.AuthData
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

			go func(data auth.AuthData) {
				res, err = auth.Authorize(context.Background(), data)

				if err != nil {
					log.Printf("failed to auth user: %v", err)
					pushRedis(context.Background(), data.Taskid, "fail", 1*time.Hour)
					return
				}

				//добавить пуш токена в Redis

				pushRedis(context.Background(), data.Taskid, "success", 1*time.Hour)
			}(data)

			fmt.Printf("Получено сообщение: %v\n", data)

			err = reader.CommitMessages(context.Background(), message)
			if err != nil {
				log.Printf("failed to commit message: %v", err)
			}
		}
	}(reader_auth)

	select {}
}
