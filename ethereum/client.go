package ethereum

import (
	"context"
	"log"
	"math/big"

	"github.com/NutiNaguti/night-bridge-oracle/near"
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

func (p *EthProvider) SubscribeToEvents(ctx context.Context, address *common.Address, req chan near.InsertBloomFilterRequest, proof chan common.Hash) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{*address},
	}

	subscribtionCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	logs := make(chan types.Log)
	sub, err := p.client.SubscribeFilterLogs(subscribtionCtx, query, logs)
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			blockNumber := big.NewInt(int64(vLog.BlockNumber))
			blockHeader := p.getBlockHeader(*blockNumber)
			proof <- vLog.Topics[2]
			req <- near.InsertBloomFilterRequest{
				BlockNumber: vLog.BlockNumber,
				Logs:        blockHeader.Bloom.Bytes(),
			}
		}
	}
}

func (p *EthProvider) getBlockHeader(blockNumber big.Int) *types.Header {
	blockHeader, err := p.client.HeaderByNumber(context.Background(), &blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	return blockHeader
}
