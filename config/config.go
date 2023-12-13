// config.go
package config

import (
	"os"
	"sync"
)

type AppConfig struct {
	WeatherAPIKey string
}

var (
	once     sync.Once
	instance *AppConfig
)

// GetConfig returns the singleton instance of AppConfig.
func GetConfig() *AppConfig {
	once.Do(func() {
		instance = loadConfig()
	})
	return instance
}

func loadConfig() *AppConfig {
	return &AppConfig{
		WeatherAPIKey: getEnv("WEATHER_API_KEY", ""),
	}
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
