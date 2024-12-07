package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadConfig loads environment variables from the .env file
func LoadConfig() {
	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Example: Accessing environment variables
	// dbConnection := os.Getenv("DB_CONNECTION_STRING")
	// if dbConnection == "" {
	//     log.Fatal("DB_CONNECTION_STRING is not set in the .env file")
	// }
}

// GetEnv returns the value of an environment variable or a default value if not set
func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
