// go get github.com/joho/godotenv
package services

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
}

func GetEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Printf("Warning: %s environment variable not set\n", key)
	}
	return value
}
