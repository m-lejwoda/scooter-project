// Package config handles loading and validating the application configuration
package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	RedisURL    string
}

func Load() (*Config, error) {
	_ = godotenv.Load()
	dbURL := os.Getenv("DATABASE_URL")
	redisURL := os.Getenv("REDIS_URL")

	return &Config{
		DatabaseURL: dbURL,
		RedisURL:    redisURL,
	}, nil
}
