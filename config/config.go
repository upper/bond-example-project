package config

import (
	"os"
)

type config struct {
	Database database
}

func env(name string, defaultValue string) string {
	if value := os.Getenv(name); value != "" {
		return value
	}
	return defaultValue
}

var Config = config{
	Database: database{
		Host:     env("DATABASE_HOST", "127.0.0.1"),
		Name:     env("DATABASE_NAME", "bond"),
		User:     env("DATABASE_USER", "postgres"),
		Password: env("DATABASE_PASSWORD", "postgres"),
	},
}
