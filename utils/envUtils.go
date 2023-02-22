package utils

import (
	"github.com/joho/godotenv"
	"os"
)

// GetEnvVariable
// This function is used to get environment variable, either from .env file or from OS
func GetEnvVariable(key string) string {

	// Load .env file if any
	_ = godotenv.Load(".env")

	// Return value of environment variable
	return os.Getenv(key)
}
