package store

import (
	"context"
	"fibonachi/internal/store/redisclient"
	"github.com/go-redis/redis/v8"
	"time"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type RedisClient interface {
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	Get(ctx context.Context, key string) (value string, err error)
}

type Store struct {
	RedisClient
}

func NewStore(redis *redis.Client) *Store {
	return &Store{
		RedisClient: redisclient.NewClient(redis),
	}
}
