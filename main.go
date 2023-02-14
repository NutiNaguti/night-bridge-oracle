package main

import (
	"context"
	"log"

	"github.com/NutiNaguti/night-bridge-oracle/config"
	"github.com/NutiNaguti/night-bridge-oracle/ethereum"
	"github.com/NutiNaguti/night-bridge-oracle/http"
	nearClient "github.com/NutiNaguti/night-bridge-oracle/near"
	near "github.com/aurora-is-near/near-api-go"
	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
)

var nearAccount *near.Account

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	conf := config.New()
	http.SetBaseURI(conf.HttpClient.BaseURI)
	ethProvider, bridgeAddress := ethSetup(conf)
	err := nearSetup(conf)
	if err != nil {
		log.Printf("%v", err)
	}

	// --------- Listen events ----------
	log.Print("Subscribtion to event started...")
	subscriptionCtx, subscriptionCancel := context.WithCancel(context.Background())
	defer subscriptionCancel()

	eventData := make(chan nearClient.InsertBloomFilterRequest)
	proof := make(chan common.Hash, 100)
	go ethProvider.SubscribeToEvents(subscriptionCtx, bridgeAddress, eventData, proof)

	indexerReq := make(chan *http.IndexerRequest, 100)
	for {
		select {
		case <-subscriptionCtx.Done():
			log.Printf("%s", subscriptionCtx.Err())
			return
		case insertBloomFilterReq := <-eventData:
			log.Print("Sending data to NEAR started...")

			insertCtx, insertCancel := context.WithCancel(context.Background())
			go nearClient.InsertBloomFilter(insertCtx, nearAccount, conf.Near.BridgeAccountId, insertBloomFilterReq, indexerReq)
			insertCancel()

			addTxCtx, addTxCancel := context.WithCancel(context.Background())
			go http.AddNewTransaction(addTxCtx, <-indexerReq)
			addTxCancel()
		}
	}
}

func ethSetup(conf *config.Config) (ethProvider *ethereum.EthProvider, bridgeAddress *common.Address) {
	log.Print("Ethereum connection started...")
	ethProvider = ethereum.New(conf.Eth.Endpoint)
	log.Print("Ethereum node connected")
	bridgeAddress = ethereum.HexToAddress(conf.Eth.BridgeAddress)
	return
}

func nearSetup(conf *config.Config) error {
	log.Print("NEAR connection started...")
	var err error
	nearConnection := nearClient.NewConnection(conf.Near.Endpoint)
	nearConfig := nearClient.SetConfig(conf.Near.Endpoint, conf.Near.NetworkId, conf.Near.KeyPath)
	nearAccount, err = nearClient.LoadAccount(nearConnection, nearConfig, conf.Near.ServiceAccountId)
	if err != nil {
		return err
	}
	log.Print("NEAR node connected")
	return err
}
