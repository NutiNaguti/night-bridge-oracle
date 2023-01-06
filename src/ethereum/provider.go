package ethereum

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

func HexToAddress(hex string) *common.Address {
	address := common.HexToAddress(hex)
	return &address
}

func setupClient(connectUrl string) *ethclient.Client {
	client, err := ethclient.Dial(connectUrl)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func (p *EthProvider) GetBalanceOf(address *common.Address) big.Int {
	balance, err := p.client.BalanceAt(context.Background(), *address, nil)
	if err != nil {
		log.Fatal(err)
	}

	return *balance
}

func (p *EthProvider) SubscribeToEvents(address *common.Address, logs chan types.Log) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{*address},
	}

	sub, err := p.client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			log.Print(vLog)
		}
	}
}
