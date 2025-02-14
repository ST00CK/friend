package config

import (
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv() {
	if godotenv.Load() != nil {
		GetEnv("NEO4J_URI")
		GetEnv("NEO4J_USERNAME")
		GetEnv("NEO4J_PASSWORD")
		GetEnv("APP_PORT")
	} else {
		err := godotenv.Load()
		if err != nil {
			println("Error loading .env file")
		}
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
