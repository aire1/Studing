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
	UpdateNoteStatusPrefix         = "updateNote_task"
	DeleteNoteStatusPrefix         = "deleteNote_task"
)

var typeMap = map[string]reflect.Type{
	RegistrationStatusPrefix:       reflect.TypeOf(RegistrationStatus{}),
	AuthorizationGetStatusPrefix:   reflect.TypeOf(AuthorizationGetStatus{}),
	AuthorizationCheckStatusPrefix: reflect.TypeOf(AuthorizationCheckStatus{}),
	CreateNoteStatusPrefix:         reflect.TypeOf(CreateNoteStatus{}),
	GetNoteStatusPrefix:            reflect.TypeOf(GetNoteStatus{}),
}

// Парсит ключ задачи Redis вида username:prefix:uuid в контекст
type contextKey string

const (
	UsernameRedisKey contextKey = "usernameRedis"
	PrefixKey        contextKey = "prefix"
	UuidKey          contextKey = "uuid"
	UsernameKey      contextKey = "username"
)

func ParseKeyToContext(ctx *context.Context, key string) error {
	//Ключ задачи в Redis имеет вид: username:prefix:uuid
	parts := strings.Split(key, ":")
	if len(parts) < 3 {
		return fmt.Errorf("invalid task id: %s", key)
	}

	username, prefix, uuid := parts[0], parts[1], parts[2]

	*ctx = context.WithValue(*ctx, UsernameRedisKey, username)
	*ctx = context.WithValue(*ctx, PrefixKey, prefix)
	*ctx = context.WithValue(*ctx, UuidKey, uuid)

	return nil
}

func GetTaskFromRedis(client *redis.Client, ctx context.Context, key string) (TaskStatus, error) {
	var v TaskStatus
	if typ, ok := typeMap[ctx.Value(PrefixKey).(string)]; ok {
		v = reflect.New(typ).Interface().(TaskStatus)
	} else {
		return nil, fmt.Errorf("unknown prefix: %s", ctx.Value(PrefixKey).(string))
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
