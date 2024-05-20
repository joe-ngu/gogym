package utils

import(
  "os"
	"log"
	"errors"

	"github.com/joho/godotenv"
)

func LoadDotenv() error {
	// Check if .env file exists
	_, err := os.Stat(".env")
	if os.IsNotExist(err) {
		return errors.New(".env file not found")
	}
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
		return err
	}
	return nil
}

func Getenv(env string) (string, error) {
	if value, found := os.LookupEnv(env); !found {
		return value, errors.New(env + " environment variable not set")
	} else {
		return value, nil
	}
}
