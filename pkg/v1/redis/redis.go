package redis

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	client *redis.Client
	once   sync.Once
)

func NewRedisClient() *redis.Client {
	once.Do(func() {
		client = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
		if err := client.Ping(context.Background()).Err(); err != nil {
			panic(fmt.Errorf("error while connecting to Redis: %v", err))
		}
	})
	return client
}

func SaveToken(ctx context.Context, userId uint, token string, duration time.Duration) error {
	client := NewRedisClient()
	redisKey := fmt.Sprintf("user:%d:refreshToken", userId)
	return client.Set(ctx, redisKey, token, duration).Err()

}

func DeleteTokenByUserId(ctx context.Context, userId uint) error {
	client := NewRedisClient()
	redisKey := fmt.Sprintf("user:%d:refreshToken", userId)

	_, err := client.Get(ctx, redisKey).Result()
	if err != nil {
		return err
	}

	return client.Del(ctx, redisKey).Err()
}
func GetByUserId(ctx context.Context, userId string) error {
	client := NewRedisClient()
	redisKey := fmt.Sprintf("user:%s:refreshToken", userId)

	_, err := client.Get(ctx, redisKey).Result()
	if err != nil {
		return err
	}
	return nil
}
