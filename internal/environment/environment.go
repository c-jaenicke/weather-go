package environment

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

var Path string

// LoadEnv loads environment variables from the .env file at the given path
func LoadEnv(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return fmt.Errorf("failed to load .env file from path %s: %s", path, err)
	}
	return nil
}

// GetEnv retrieves the value assigned to the environment variable `key`
func GetEnv(key string) (string, error) {
	return os.Getenv(key), nil
}
