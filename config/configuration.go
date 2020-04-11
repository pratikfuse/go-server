package config

import (
	"os"

	"github.com/joho/godotenv"
)

func Load() error {
	environment := os.Getenv("APP_ENV")
	if environment == "" {
		environment = "local"
	}
	err := godotenv.Load(".env." + environment)
	return err
}
