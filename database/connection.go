package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)


//Loading the database URI for connection
func DatabaseURIEnvVariable(URI string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(URI)
}