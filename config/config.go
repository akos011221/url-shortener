package config

import (
	"errors"
	"os"
)

type Config struct {
	Env           string
	ServerAddress string
	DatabaseURL   string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		Env:           os.Getenv("ENV"),
		ServerAddress: os.Getenv("SERVER_ADDRESS"),
		DatabaseURL:   os.Getenv("DATABASE_URL"),
	}

	// Set default
	if cfg.Env == "" {
		cfg.Env = "development"
	}
	if cfg.ServerAddress == "" {
		cfg.ServerAddress = ":8080"
	}

	// Validate required fields
	if cfg.DatabaseURL == "" && cfg.Env == "production" {
		return nil, errors.New("DATABASE_URL is required in production")
	}

	return cfg, nil
}
