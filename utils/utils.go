package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	godotenv.Load(".env")

	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("FATAL: %s environment variable is not set", key)
	}

	return value
}
