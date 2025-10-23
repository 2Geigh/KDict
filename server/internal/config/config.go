package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	ApiUrlWithoutKey = "https://krdict.korean.go.kr/api/search?key="
)

var (
	ApiKey        string
	ApiUrlWithKey string
)

// LoadConfig loads the environment variables and constructs the API URL
func LoadConfig() {

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Read the API key from environment variables
	ApiKey = os.Getenv("API_KEY")
	if ApiKey == "" {
		log.Fatal("API_KEY environment variable is not set")
	}

	// Construct the full API URL
	ApiUrlWithKey = ApiUrlWithoutKey + ApiKey

}
