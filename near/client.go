package near

import (
	"log"
	"math/big"

	"github.com/NutiNaguti/night-bridge-oracle/http"
	near "github.com/aurora-is-near/near-api-go"
	"github.com/near/borsh-go"
)

const TGAS = 100_000_000_000_000

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

func InsertBloomFilter(account *near.Account, bridgeId string, req InsertBloomFilterRequest, indexerReq chan *http.IndexerRequest) {
	serializedReq, err := borsh.Serialize(req)
	if err != nil {
		log.Fatal(err)
	}

	res, err := account.FunctionCall(bridgeId, "insert_filter", serializedReq, TGAS, *big.NewInt(0))
	if err != nil {
		log.Fatal(err)
	}

	tx := res["transaction"].(map[string]interface{})
	sender := tx["signer_id"].(string)
	receiver := tx["receiver_id"].(string)
	indexerReq <- &http.IndexerRequest{
		Sender:   sender,
		Receiver: receiver,
		// TODO: unhardcode
		Amount:    "100",
		Timestamp: "1674696071",
	}
}
