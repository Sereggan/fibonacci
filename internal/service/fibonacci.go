package service

import (
	"context"
	"encoding/json"
	"fibonachi/internal/store"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

type FibonacciService struct {
	redisClient store.RedisClient
}

func NewFibonacciService(store store.RedisClient) *FibonacciService {
	return &FibonacciService{
		redisClient: store,
	}
}

func (f *FibonacciService) CalculateResult(ctx context.Context, from uint64, to uint64) ([]uint64, error) {
	key := fmt.Sprintf("%d_%d", from, to)
	data, err := f.redisClient.Get(ctx, key)
	if err != nil {
		logrus.Errorf("Failed to get value by key: %s, err: %s", key, err.Error())
	}
	var values []uint64
	err = json.Unmarshal([]byte(data), &values)

	if len(values) != 0 {
		logrus.Infof("Found key in Redis, key: %s, values: %v", key, values)
		return values, nil
	}

	fibonacciValues := calculateNumbers(from, to)

	err = f.redisClient.Set(ctx, key, fibonacciValues, time.Minute*3600)

	if err != nil {
		logrus.Errorf("Failed to save value by key: %s, err: %s", key, err.Error())
	}

	return fibonacciValues, nil
}

func calculateNumbers(from uint64, to uint64) []uint64 {
	var values []uint64

	var x1 uint64 = 1
	var x2 uint64 = 1
	var counter uint64 = 0

	for counter < from {
		x1, x2 = x2, x1+x2
		counter++
	}

	for counter <= to {
		values = append(values, x1)
		x1, x2 = x2, x1+x2
		counter++
	}

	return values
}
