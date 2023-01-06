package main

import (
	"fmt"
	"log"

	"github.com/NutiNaguti/night-bridge-oracle/config"
	"github.com/NutiNaguti/night-bridge-oracle/ethereum"
	"github.com/NutiNaguti/night-bridge-oracle/near"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file")
	}
}

func main() {
	conf := config.New()

	// --------- Ethereum ---------
	log.Print("Ethereum connection started...")
	ethProvider := ethereum.New(conf.Eth.Endpoint)
	log.Print("Ethereum node connected")
	bridgeAddress := ethereum.HexToAddress(conf.Eth.BridgeAddress)

	// --------- NEAR ---------
	log.Print("NEAR connection started...")
	nearConnection := near.NewConnection(conf.Near.Endpoint)
	nearConfig := near.SetConfig(conf.Near.Endpoint, conf.Near.NetworkId, conf.Near.KeyPath)
	nearAccount := near.LoadAccount(nearConnection, nearConfig, conf.Near.ServiceAccountId)
	log.Print("NEAR node connected")

	log.Print("Subscribtion to event started...")
	req := make(chan near.InsertBloomFilterRequest)
	go ethProvider.SubscribeToEvents(bridgeAddress, req)

	log.Print("Sending data to NEAR started...")
	go near.InsertBloomFilter(nearAccount, conf.Near.BridgeAccountId, <-req)

	fmt.Scanln()
}
