package config

import (
	"os"
)

type EthConfig struct {
	Endpoint      string
	BridgeAddress string
}

type NearConfig struct {
	Endpoint        string
	BridgeAccountId string
}

type Config struct {
	Eth  EthConfig
	Near NearConfig
}

func New() *Config {
	return &Config{
		Eth: EthConfig{
			Endpoint:      getEnv("ETH_ENDPOINT", ""),
			BridgeAddress: getEnv("ETH_BRIDGE_ADDRESS", ""),
		},
		Near: NearConfig{
			Endpoint:        getEnv("NEAR_ENDPOINT", ""),
			BridgeAccountId: getEnv("NEAR_BRIDGE_ACCOUNT_ID", ""),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}