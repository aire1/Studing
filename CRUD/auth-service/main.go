package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"

	pb "crud/auth-service/proto"
	pg "crud/common-libs/postgres"

	//rd "crud/auth-service/common-libs/redis"

	rd "crud/common-libs/redis"

	authCheck "crud/auth-service/checkAuthorization"
	authGet "crud/auth-service/getAuthorization"
	reg "crud/auth-service/registration"

	shared "crud/common-libs/shared"

	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authserver struct {
	pb.UnimplementedAuthGateServer
}

func (s *authserver) CheckAuthorization(ctx context.Context, req *pb.AuthCheckRequest) (*pb.AuthCheckResponse, error) {
	username, err := authCheck.CheckAuthorization(ctx, req)
	if err != nil {
		return &pb.AuthCheckResponse{
			Status: false,
		}, status.Errorf(codes.Unauthenticated, "%s", err.Error())
	}

	return &pb.AuthCheckResponse{
		Status:   true,
		Username: username,
	}, nil
}

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
					BaseTaskStatus: shared.BaseTaskStatus{
						Result: "success",
						Info:   "",
					},
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
					BaseTaskStatus: shared.BaseTaskStatus{
						Result: "success",
						Info:   "",
					},
				}

				if token, err := authGet.Authorize(context.Background(), data); err != nil {
					log.Printf("failed to auth user: %v", err)

					status.Result = "fail"
					status.Info = err.Error()
				} else {
					status.Info = token
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

	//Создаем сервер gRPC auth-gate
	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Ошибка запуска auth-gate: %v", err)
		}
		s := grpc.NewServer()

		log.Printf("Поднял gRPC сервер на 50051")

		pb.RegisterAuthGateServer(s, &authserver{})
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Ошибка запуска auth-gate: %v", err)
		}
	}()

	select {}
}
