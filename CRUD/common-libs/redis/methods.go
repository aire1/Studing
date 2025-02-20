package redis

import (
	"context"
	"encoding/json"
	"time"
)

func PushStatusIntoRedis(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	json_b, err := json.Marshal(value)
	if err != nil {
		return err
	}

	if err := Client.Set(ctx, key, json_b, time.Hour*1).Err(); err != nil {
		return err
	}

	return nil
}
