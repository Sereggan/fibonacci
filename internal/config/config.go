package config

import (
	"os"
)

type Config struct {
	ServerAddress string
	RedisAddress  string
	GrpcPort      string
}

func New() *Config {
	return &Config{
		ServerAddress: getEnv("SERVER_ADDRESS", "localhost:8080"),
		RedisAddress:  getEnv("REDIS_ADDRESS", ""),
		GrpcPort:      getEnv("GRPC_SERVER_PORT", "5300"),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
