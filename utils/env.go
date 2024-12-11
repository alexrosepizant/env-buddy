package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnvVariable loads a specific key from a .env file and sets it as an environment variable
func LoadEnvVariable(filename, key string) error {
	// Load the .env file
	envMap, err := godotenv.Read(filename)
	if err != nil {
		return fmt.Errorf("error reading .env file: %w", err)
	}

	// Get the value for the specified key
	value, exists := envMap[key]
	if !exists {
		return fmt.Errorf("key %s not found in .env file", key)
	}

	// Set the environment variable
	err = os.Setenv(key, value)
	if err != nil {
		return fmt.Errorf("error setting environment variable: %w", err)
	}

	return nil
} 