package shared

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/go-redis/redis/v8"
)

const (
	RegistrationStatusPrefix       = "registration_task:"
	AuthorizationGetStatusPrefix   = "getAuthorization_task:"
	AuthorizationCheckStatusPrefix = "checkAuthorization_task:"
)

type RedisSerializable interface {
}

func GetTaskFromRedis(client *redis.Client, ctx context.Context, key string) (any, error) {
	var v RedisSerializable

	switch {
	case strings.HasPrefix(key, RegistrationStatusPrefix):
		v = &RegistrationStatus{}
	case strings.HasPrefix(key, AuthorizationGetStatusPrefix):
		v = &AuthorizationGetStatus{}
	case strings.HasPrefix(key, AuthorizationCheckStatusPrefix):
		v = &AuthorizationCheckStatus{}
	}

	if err := UnmarshalFromRedis(client, ctx, key, v); err != nil {
		return nil, err
	}

	return v, nil
}

func UnmarshalFromRedis(client *redis.Client, ctx context.Context, key string, v RedisSerializable) error {
	json_b, err := client.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(json_b), v); err != nil {
		return err
	}

	return nil
}
