package near

import (
	"log"

	near "github.com/aurora-is-near/near-api-go"
)

func NewConnection(nodeUrl string) *near.Connection {
	connection := near.NewConnection(nodeUrl)
	return connection
}

func SetConfig(networkID string, nodeUrl string, keyPath string) *near.Config {
	config := near.Config{
		NetworkID: networkID,
		NodeURL:   nodeUrl,
		KeyPath:   keyPath,
	}

	return &config
}

func LoadAccount(connection *near.Connection, cfg *near.Config, receiverId string) *near.Account {
	account, err := near.LoadAccount(connection, cfg, receiverId)
	if err != nil {
		log.Fatal(err)
	}

	return account
}
