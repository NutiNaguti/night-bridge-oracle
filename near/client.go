package near

import (
	"encoding/json"
	"log"
	"math/big"

	"github.com/NutiNaguti/night-bridge-oracle/http"
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

func InsertBloomFilter(account *near.Account, receiverId string, req InsertBloomFilterRequest, indexerReq chan *http.IndexerRequest) {
	serializedReq, err := borsh.Serialize(req)
	if err != nil {
		log.Fatal(err)
	}
	res, err := account.FunctionCall(receiverId, "insert_filter", serializedReq, 100_000_000_000_000, *big.NewInt(0))
	if err != nil {
		log.Fatal(err)
	}
	b, err := json.Marshal(res)
	if err != nil {
		log.Print(err)
	}
	log.Print(string(b))

	// tx := res["transaction"].(map[string]interface{})
	// sender := tx["signer"].(string)
	// receiver := tx["receiver_id"].(string)
	indexerReq <- &http.IndexerRequest{
		Sender:    "nutinaguti.testnet",
		Receiver:  "0x8CAB5E96E1ab09e8678a8ffC75b5D818e73D4707",
		Amount:    "100",
		Timestamp: "1674696071",
	}
}
