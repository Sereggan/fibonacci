package redisclient

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisClient struct {
	c *redis.Client
}

func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	p, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.c.Set(ctx, key, p, ttl).Err()
}

func (r *RedisClient) Get(ctx context.Context, key string) (value string, err error) {
	return r.c.Get(ctx, key).Result()
}

func NewClient(c *redis.Client) *RedisClient {
	return &RedisClient{c}
}
