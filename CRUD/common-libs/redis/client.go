package redis

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var Client *redis.Client

func Init() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "redis:6379", // адрес вашего Redis сервера
		Password: "",           // пароль, если установлен
		DB:       0,            // используемая база данных
	})

	_, err := Client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
}
