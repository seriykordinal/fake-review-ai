package config

import (
	"os"
)

type Config struct {
	ServerPort   string
	PostgresDSN  string
	MLServiceURL string
}

func Load() *Config {

	return &Config{
		ServerPort:   os.Getenv("SERVER_PORT"),
		PostgresDSN:  os.Getenv("POSTGRES_DSN"),
		MLServiceURL: os.Getenv("ML_SERVICE_URL"),
	}
}
