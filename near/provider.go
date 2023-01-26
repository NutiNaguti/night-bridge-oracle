package near

import (
	"log"
	"math/big"

	near "github.com/aurora-is-near/near-api-go"
	"github.com/near/borsh-go"
)

type InsertBloomFilterRequest struct {
	BlockNumber uint64 `json:"block_number"`
	Logs        string `json:"logs"`
}

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

func InsertBloomFilter(account *near.Account, receiverId string, req InsertBloomFilterRequest) {
	serializedReq, err := borsh.Serialize(req)
	if err != nil {
		log.Fatal(err)
	}
	actions := []near.Action{{FunctionCall: near.FunctionCall{MethodName: "insert_filter", Args: serializedReq, Gas: 100_000_000_000_000, Deposit: *big.NewInt(1)}}}
	res, err := account.SignAndSendTransaction(receiverId, actions)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(res)
}
