package initalizers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	// load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environment variables")
	}
}
