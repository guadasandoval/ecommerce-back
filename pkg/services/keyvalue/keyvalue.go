package keyvalue

import (
	"context"
	"ecommerce/pkg/libs/errors"
	"time"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

// SetValue function
func SetValue(key string, value interface{}) error {
	const op errors.Operation = "pkg.services.keyvalue.SetValue"
	status := redisClient.Set(context.Background(), key, value, 8*time.Hour)
	if status.Err() != nil {
		return errors.NewError(op, errors.KindUnexpected, status.Err())
	}

	return nil
}

// GetInt function
func GetInt(key string) (int, error) {
	const op errors.Operation = "pkg.services.keyvalue.GetInt"
	strCmd := redisClient.Get(context.Background(), key)
	if strCmd.Err() != nil {
		return 0, errors.NewError(op, errors.KindUnexpected, strCmd.Err())
	}

	value, err := strCmd.Int()
	if err != nil {
		return 0, errors.NewError(op, errors.KindUnexpected, err)
	}

	return value, nil
}

// HasValue function
func HasValue(key string) (bool, error) {
	const op errors.Operation = "pkg.services.keyvalue.HasValue"
	_, err := redisClient.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		return false, errors.NewError(op, errors.KindUnexpected, err)
	}

	return true, nil
}

// DeleteValue function
func DeleteValue(key string) error {
	const op errors.Operation = "services.keyvalue.DeleteValue"
	intCmd := redisClient.Del(context.Background(), key)
	if intCmd.Err() != nil {
		return errors.NewError(op, errors.KindUnexpected, intCmd.Err())
	}

	return nil
}
