package database

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient interface {
	Set(ctx context.Context, key string, value interface{}) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, key string) error
	Exists(ctx context.Context, key string) (bool, error)
	Close() error
}

type redisClient struct {
	client *redis.Client
}

func NewRedisClient() RedisClient {
	redis := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return &redisClient{
		client: redis,
	}
}

func (r *redisClient) Set(ctx context.Context, key string, value interface{}) error {
	return r.client.Set(ctx, key, value, 5*time.Minute).Err()

}

func (r *redisClient) Get(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()

	if err != nil {
		return "", err
	}
	return val, nil

}

func (r *redisClient) Del(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

func (r *redisClient) Exists(ctx context.Context, key string) (bool, error) {
	result, err := r.client.Exists(ctx, key).Result()
	return result > 0, err
}

func (r *redisClient) Close() error {
	return r.client.Close()
}
