package main

import (
	"fmt"
	"log"

	"github.com/NutiNaguti/night-bridge-oracle/config"
	"github.com/NutiNaguti/night-bridge-oracle/ethereum"
	"github.com/NutiNaguti/night-bridge-oracle/near"
	"github.com/ethereum/go-ethereum/core/types"
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
	_ = nearConnection
	log.Print("NEAR node connected")

	log.Print("Subscribtion to event started...")
	logs := make(chan types.Log)
	go ethProvider.SubscribeToEvents(bridgeAddress, logs)

	log.Print("Sending data to NEAR...")

	fmt.Scanln()
}
