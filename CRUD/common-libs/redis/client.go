package redis

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var Client *redis.Client

func Init() {
	Client = redis.NewClient(&redis.Options{
		//Addr:     "redis:6379",
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := Client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
}
