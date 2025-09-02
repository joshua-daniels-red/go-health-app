package config

import (
	"os"
)

// Config holds all application configuration
type Config struct {
	Port string
}

// Load initializes configuration from environment variables
func Load() (*Config, error) {
	cfg := &Config{
		Port: getEnv("PORT", "8080"), // default to 8080 if not set
	}
	return cfg, nil
}

// helper to check environment variables with defaults
func getEnv(key, defaultVal string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}
	return defaultVal
}
