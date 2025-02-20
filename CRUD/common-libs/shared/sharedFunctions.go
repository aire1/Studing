package shared

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-redis/redis/v8"
)

const (
	RegistrationStatusPrefix       = "registration_task"
	AuthorizationGetStatusPrefix   = "getAuthorization_task"
	AuthorizationCheckStatusPrefix = "checkAuthorization_task"
	CreateNoteStatusPrefix         = "createNote_task"
	GetNoteStatusPrefix            = "getNote_task"
)

var typeMap = map[string]reflect.Type{
	RegistrationStatusPrefix:       reflect.TypeOf(RegistrationStatus{}),
	AuthorizationGetStatusPrefix:   reflect.TypeOf(AuthorizationGetStatus{}),
	AuthorizationCheckStatusPrefix: reflect.TypeOf(AuthorizationCheckStatus{}),
	CreateNoteStatusPrefix:         reflect.TypeOf(CreateNoteStatus{}),
	GetNoteStatusPrefix:            reflect.TypeOf(GetNoteStatus{}),
}

func GetStatusPrefix(key string) (string, error) {
	parts := strings.Split(key, ":")
	if len(parts) < 2 {
		return "", fmt.Errorf("invalid task id: %s", key)
	}

	return parts[1], nil
}

func GetTaskFromRedis(client *redis.Client, ctx context.Context, key string) (TaskStatus, error) {
	prefix, err := GetStatusPrefix(key)
	if err != nil {
		return nil, err
	} else if prefix == "" {
		return nil, fmt.Errorf("prefix is nil")
	}

	var v TaskStatus
	if typ, ok := typeMap[prefix]; ok {
		v = reflect.New(typ).Interface().(TaskStatus)
	} else {
		return nil, fmt.Errorf("unknown prefix: %s", prefix)
	}

	if err := UnmarshalFromRedis(client, ctx, key, v); err != nil {
		return nil, err
	}

	return v, nil
}

func UnmarshalFromRedis(client *redis.Client, ctx context.Context, key string, v TaskStatus) error {
	json_b, err := client.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(json_b), v); err != nil {
		return err
	}

	return nil
}
