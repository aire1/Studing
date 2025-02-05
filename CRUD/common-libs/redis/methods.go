package redis

import (
	"context"
	"time"

	"github.com/pkg/errors"
)

func PushStatusIntoRedis(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	// if err := Client.Set(ctx, key, value, time.Hour*1).Err(); err != nil {
	// 	return errors.Errorf("can't push into reddis: %v", err)
	// }
	if err := Client.Publish(ctx, key, value); err != nil {
		return errors.Errorf("can't push into reddis: %v", err)
	}

	return nil
}
