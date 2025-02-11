package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	pg "crud/common-libs/postgres"
	//rd "crud/auth-service/common-libs/redis"

	rd "crud/common-libs/redis"

	authCheck "crud/auth-service/checkAuthorization"
	authGet "crud/auth-service/getAuthorization"
	reg "crud/auth-service/registration"

	shared "crud/common-libs/shared"

	"github.com/segmentio/kafka-go"
)

func main() {
	rd.Init()
	defer rd.Client.Close()
	pg.CreatePool()
	defer pg.Pool.Close()

	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)

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
		Topic:          "get_authorizations",
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
		Topic:          "check_authorizations",
		GroupID:        "Auth-service",
		MinBytes:       10e3,
		MaxBytes:       10e6,
		MaxWait:        500 * time.Millisecond,
		CommitInterval: time.Second,
	})
	defer reader_auth_checks.Close()

	log.Println("Подключился к Kafka")

	go func(reader *kafka.Reader) {
		var data shared.RegistrationData
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

			go func(data shared.RegistrationData) {
				status := shared.RegistrationStatus{
					Result: "success",
					Info:   "",
				}

				err = reg.Register(context.Background(), data)

				if err != nil {
					log.Printf("failed to register user: %v", err)

					status.Result = "fail"
					status.Info = err.Error()
				}

				if json_data, err := json.Marshal(status); err == nil {
					rd.PushStatusIntoRedis(context.Background(), data.TaskId, json_data, time.Hour)
				} else {
					log.Printf("failed to marshal status info")
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
		var data shared.AuthorizationGetData
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

			go func(data shared.AuthorizationGetData) {
				status := shared.AuthorizationGetStatus{
					Result: "success",
					Info:   "",
				}

				if err = authGet.Authorize(context.Background(), data); err != nil {
					log.Printf("failed to auth user: %v", err)

					status.Result = "fail"
					status.Info = err.Error()
				}

				if json_data, err := json.Marshal(status); err == nil {
					rd.PushStatusIntoRedis(context.Background(), data.TaskId, json_data, time.Hour)
				} else {
					log.Printf("failed to marshal status info")
				}
			}(data)

			fmt.Printf("Получено сообщение: %v\n", data)

			err = reader.CommitMessages(context.Background(), message)
			if err != nil {
				log.Printf("failed to commit message: %v", err)
			}
		}
	}(reader_auth)

	go func(reader *kafka.Reader) {
		var data shared.AuthorizationCheckData
		for {
			message, err := reader.FetchMessage(context.Background())
			if err != nil {
				log.Fatal(err)
			}

			log.Printf("Получил сообщение")

			err = json.Unmarshal(message.Value, &data)
			if err != nil {
				log.Printf("failed to unmarshal message: %v", err)
				continue
			}

			go func(data shared.AuthorizationCheckData) {
				log.Printf("Начал выполнять задание")

				status := shared.AuthorizationCheckStatus{
					Result: "success",
					Info:   "",
				}

				if err := authCheck.CheckAuthorization(context.Background(), data); err != nil {
					log.Printf("denied to auth user: %v", err)
					status.Result = "fail"
					status.Info = err.Error()
				}

				log.Printf("Задание выполнено")

				if json_data, err := json.Marshal(status); err == nil {
					rd.PushStatusIntoRedis(context.Background(), data.TaskId, json_data, time.Hour)
				} else {
					log.Printf("failed to marshal status info")
				}

				log.Printf("Ответ в Redis отправлен")
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
