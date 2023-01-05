package config

import (
	"os"
)

type EthConfig struct {
	Endpoint string
}

type Config struct {
	Eth EthConfig
}

func New() *Config {
	return &Config{
		Eth: EthConfig{
			Endpoint: getEnv("ETH_ENDPOINT", ""),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
