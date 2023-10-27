package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        int
	DatabaseURL string
	// Add other configuration variables as needed
}

var CONFIG Config

func LoadConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	CONFIG = Config{
		Port:        getEnvAsInt("PORT", 8080), // Default to port 8080 if PORT is not set
		DatabaseURL: getEnvAsString("DATABASE_URL", ""),
		// Add other configuration variables here
	}

	return CONFIG
}

func GetConfig() Config {
	return CONFIG
}

func getEnvAsString(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value, ok := os.LookupEnv(key); ok {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}

	return defaultValue
}
