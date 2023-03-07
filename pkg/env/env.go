package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

var Path string

// GetEnv get value assigned to env value
func GetEnv(key string) (string, error) {
	err := godotenv.Load(Path)
	if err != nil {
		return "", fmt.Errorf("failed to get env variable! %s, Error:  %s, Path: %s", key, err.Error(), Path)
	}

	return os.Getenv(key), nil
}
