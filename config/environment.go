package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvironmentVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Print("Error loading .env file")
	}
}
