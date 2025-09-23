package config

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

// Get retrieves the value of the environment variable named by the key.
// It returns the value, which will be empty if the variable is not present.
func Get(Key string) string {
	return os.Getenv(Key)
}

func GetRequired(Key string) (string, error) {
	value := os.Getenv(Key)
	if strings.TrimSpace(value) == "" {
		return "", errors.New("missing required env: " + Key)
	}
	return value, nil
}

func GetOrDefault(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}

func GetIntOrDefault(Key string, defaultValue int) (int, error) {
	valueStr := os.Getenv(Key)
	if valueStr == "" {
		return defaultValue, nil
	}
	i, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func MongoUser() (string, error)   { return GetRequired("MONGO_USER") }
func MongoPass() (string, error)   { return GetRequired("MONGO_PASS") }
func MongoHost() (string, error)   { return GetRequired("MONGO_HOST") }
func MongoPort() (string, error)   { return GetRequired("MONGO_PORT") }
func MongoDBName() (string, error) { return GetRequired("MONGODB_DB") }
func HTTPPort() string             { return GetOrDefault("PORT", "8080") }
