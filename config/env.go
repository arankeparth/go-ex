package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads the .env file into environment variables
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
}

// GetEnv retrieves a specific environment variable by key
func GetEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Printf("Warning: %s environment variable not set\n", key)
	}
	return value
}
