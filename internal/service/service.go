package service

import (
	"context"
	"fibonachi/internal/store"
)

type Fibonacci interface {
	CalculateResult(context.Context, uint64, uint64) ([]uint64, error)
}

type Service struct {
	Fibonacci
}

func NewService(stores *store.Store) *Service {
	return &Service{
		Fibonacci: NewFibonacciService(stores.RedisClient),
	}
}
