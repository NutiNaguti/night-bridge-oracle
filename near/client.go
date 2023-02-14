package near

import (
	"context"
	"math/big"

	"github.com/NutiNaguti/night-bridge-oracle/http"
	near "github.com/aurora-is-near/near-api-go"
	"github.com/ethereum/go-ethereum/log"
	"github.com/near/borsh-go"
)

const TGAS = 100_000_000_000_000

type InsertBloomFilterRequest struct {
	BlockNumber uint64 `json:"block_number"`
	Logs        []byte `json:"logs"`
}

type ValidateTransferRequest struct {
	BlockNumber uint64 `json:"block_number"`
	Receiver    string `json:"receiver"`
	Proof       string `json:"proof"`
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

func LoadAccount(connection *near.Connection, cfg *near.Config, receiverId string) (*near.Account, error) {
	account, err := near.LoadAccount(connection, cfg, receiverId)
	if err != nil {
		log.Error("%v", err)
		return nil, err
	}
	return account, err
}

func InsertBloomFilter(ctx context.Context, account *near.Account, bridgeId string, req InsertBloomFilterRequest, indexerReq chan *http.IndexerRequest) error {
	serializedReq, err := borsh.Serialize(req)
	if err != nil {
		log.Error("%v", err)
		return err
	}

	_, err = account.FunctionCall(bridgeId, "insert_filter", serializedReq, TGAS, *big.NewInt(0))
	if err != nil {
		log.Error("%v", err)
		return err
	}

	// tx := res["transaction"].(map[string]interface{})
	// sender := tx["signer_id"].(string)
	// receiver := tx["receiver_id"].(string)
	// indexerReq <- &http.IndexerRequest{
	// 	Sender:   sender,
	// 	Receiver: receiver,
	// }
	return err
}

// func ValidateTransfer(account *near.Account, bridgeAccountId string, req ValidateTransferRequest, logsData chan http.IndexerRequest) {
// 	serealizedReq, err := borsh.Serialize(req)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	res, err := account.FunctionCall(bridgeAccountId, "validate_transfer", serealizedReq, TGAS, *big.NewInt(1_000_000_000_000_000_000))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	logsData <- http.IndexerRequest{
// 		Sender:    "",
// 		Receiver:  "",
// 		Amount:    "",
// 		Timestamp: "",
// 	}

// 	log.Printf("%s", res)
// }
