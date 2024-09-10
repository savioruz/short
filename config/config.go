package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds the configuration settings for the application
type Config struct {
	Server ServerConfig
	Redis  RedisConfig
}

// ServerConfig holds the configuration settings for the server
type ServerConfig struct {
	Host string
	Port string
}

// RedisConfig holds the configuration settings for Redis
type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	// Load environment variables from .env file if present
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	// Get server configuration
	host := getEnv("APP_HOST", "localhost")
	port := getEnv("APP_PORT", "3000")

	// Get Redis configuration
	redisHost := getEnv("REDIS_HOST", "localhost")
	redisPort, err := strconv.Atoi(getEnv("REDIS_PORT", "6379"))
	if err != nil {
		return nil, fmt.Errorf("invalid REDIS_PORT value: %v", err)
	}
	redisPassword := getEnv("REDIS_PASSWORD", "")
	redisDB, err := strconv.Atoi(getEnv("REDIS_DB", "0"))
	if err != nil {
		return nil, fmt.Errorf("invalid REDIS_DB value: %v", err)
	}

	return &Config{
		Server: ServerConfig{
			Host: host,
			Port: port,
		},
		Redis: RedisConfig{
			Addr:     fmt.Sprintf("%s:%d", redisHost, redisPort),
			Password: redisPassword,
			DB:       redisDB,
		},
	}, nil
}

// getEnv retrieves the value of an environment variable or returns a default value if not set
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
