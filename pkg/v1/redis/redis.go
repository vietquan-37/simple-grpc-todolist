package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		panic(fmt.Errorf("error while ping redis: %v", err))
	}

	return client
}

func SaveTokenToBlackList(ctx context.Context, token string, duration time.Duration) error {
	client := NewRedisClient()
	err := client.Set(ctx, token, "revoked", duration).Err()
	if err != nil {
		return err
	}
	return nil
}
func IsTokenInBlackList(ctx context.Context, token string) (bool, error) {
	client := NewRedisClient()
	result, err := client.Get(ctx, token).Result()
	if err == redis.Nil {
		return false, nil
	}
	if err != nil {

		return false, err
	}

	if result == "revoked" {
		return true, nil
	}
	return false, nil

}
