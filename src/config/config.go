package config

import (
	"os"
)

type DbConfig struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
}

type EthConfig struct {
	Endpoint      string
	BridgeAddress string
}

type NearConfig struct {
	Endpoint          string
	NetworkId         string
	KeyPath           string
	BridgeAccountId   string
	LiteNodeAccountId string
	ServiceAccountId  string
}

type Config struct {
	Db   DbConfig
	Eth  EthConfig
	Near NearConfig
}

func New() *Config {
	return &Config{
		Db: DbConfig{
			Host:     getEnv("DB_HOST", ""),
			Port:     getEnv("DB_PORT", ""),
			Database: getEnv("DB_DATABASE", ""),
			User:     getEnv("DB_USER", ""),
			Password: getEnv("DB_PASSWORD", ""),
		},
		Eth: EthConfig{
			Endpoint:      getEnv("ETH_ENDPOINT", ""),
			BridgeAddress: getEnv("ETH_BRIDGE_ADDRESS", ""),
		},
		Near: NearConfig{
			Endpoint:          getEnv("NEAR_ENDPOINT", ""),
			NetworkId:         getEnv("NEAR_NETWORK_ID", ""),
			KeyPath:           getEnv("NEAR_KEY_PATH", ""),
			BridgeAccountId:   getEnv("NEAR_BRIDGE_ACCOUNT_ID", ""),
			LiteNodeAccountId: getEnv("NEAR_LITE_NODE_ACCOUNT_ID", ""),
			ServiceAccountId:  getEnv("NEAR_SERVICE_ACCOUNT_ID", ""),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
