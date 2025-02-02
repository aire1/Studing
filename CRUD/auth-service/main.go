package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	rd "common-libs/redis"

	"crud/auth-service/registration"

	"github.com/segmentio/kafka-go"
)

func main() {
	rd.Init()
	defer rd.Client.Close()

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{"localhost:9092", "localhost:9094", "localhost:9096"},
		Topic:          "registrations",
		GroupID:        "Auth-service",
		MinBytes:       10e3,
		MaxBytes:       10e6,
		MaxWait:        500 * time.Millisecond,
		CommitInterval: time.Second,
	})
	defer reader.Close()

	log.Println("Подключился к Kafka")

	var data registration.RegistrationData
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

		go func(data registration.RegistrationData) {
			err = registration.Register(context.Background(), data)
			if err != nil {
				log.Printf("failed to register user: %v", err)
			}
		}(data)

		fmt.Printf("Получено сообщение: %v\n", data)

		err = reader.CommitMessages(context.Background(), message)
		if err != nil {
			log.Printf("failed to commit message: %v", err)
		}
	}
}
