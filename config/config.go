package config

import "os"

type Config struct {
	Env	string
	ServerAddress string
	DatabaseURL string
}

func LoadConfig() (*Config, error) {
	return &Config{
		Env:	os.Getenv("ENV"),
		ServerAddress: os.Getenv("SERVER_ADDRESS"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}, nil
}
