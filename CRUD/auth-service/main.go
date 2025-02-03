package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	pg "crud/auth-service/common-libs/postgres"
	rd "crud/auth-service/common-libs/redis"

	authCheck "crud/auth-service/checkAuthorization"
	authGet "crud/auth-service/getAuthorization"
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

	reader_auth_checks := kafka.NewReader(kafka.ReaderConfig{
		//Brokers:        []string{"kafka1:9092", "kafka2:9092", "kafka3:9092"},
		Brokers:        []string{"localhost:19092", "localhost:19094", "localhost:19096"},
		Topic:          "authorizations",
		GroupID:        "Auth-service",
		MinBytes:       10e3,
		MaxBytes:       10e6,
		MaxWait:        500 * time.Millisecond,
		CommitInterval: time.Second,
	})
	defer reader_auth_checks.Close()

	log.Println("Подключился к Kafka")

	go func(reader *kafka.Reader) {
		var data reg.RegData
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

			go func(data reg.RegData) {
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
		var data authGet.AuthData
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

			go func(data authGet.AuthData) {
				if err = authGet.Authorize(context.Background(), data); err != nil {
					log.Printf("failed to auth user: %v", err)
					pushRedis(context.Background(), data.Taskid, "fail", 1*time.Hour)
					return
				}

				pushRedis(context.Background(), data.Taskid, "success", 1*time.Hour)
			}(data)

			fmt.Printf("Получено сообщение: %v\n", data)

			err = reader.CommitMessages(context.Background(), message)
			if err != nil {
				log.Printf("failed to commit message: %v", err)
			}
		}
	}(reader_auth)

	go func(reader *kafka.Reader) {
		var data authCheck.AuthData
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

			go func(data authCheck.AuthData) {
				if res, err := authCheck.CheckAuthorization(context.Background(), data); err != nil {
					log.Printf("failed to auth user: %v", err)
					pushRedis(context.Background(), data.Taskid, "fail", 1*time.Hour)
				} else if !res {
					log.Printf("denied to auth user %s", data.Username)
					pushRedis(context.Background(), data.Taskid, "denied", 1*time.Hour)
				} else {
					log.Printf("auth user success %s", data.Username)
					pushRedis(context.Background(), data.Taskid, "success", 1*time.Hour)
				}
			}(data)

			fmt.Printf("Получено сообщение: %v\n", data)

			err = reader.CommitMessages(context.Background(), message)
			if err != nil {
				log.Printf("failed to commit message: %v", err)
			}
		}
	}(reader_auth_checks)

	select {}
}
