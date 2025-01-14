package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(".env 에러: %v", err)
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
