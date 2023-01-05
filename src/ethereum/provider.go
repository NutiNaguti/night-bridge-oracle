package ethereum

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthProvider struct {
	client *ethclient.Client
}

func New(connectUrl string) *EthProvider {
	var client = setupClient(connectUrl)

	return &EthProvider{
		client,
	}

}

func setupClient(connectUrl string) *ethclient.Client {
	client, err := ethclient.Dial(connectUrl)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func (p *EthProvider) GetBalanceOf(hexAddress string) big.Int {
	account := common.HexToAddress(hexAddress)
	balance, err := p.client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	return *balance
}
