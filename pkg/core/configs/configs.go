package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func RootUrl() string {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	return os.Getenv("ROOT_URL")
}

func IdpMetadataUrl() string {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	return os.Getenv("IDP_METADATA_URL")
}
