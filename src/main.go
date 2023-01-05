package main

import (
	"fmt"
	"log"

	"github.com/NutiNaguti/night-bridge-oracle/config"
	"github.com/NutiNaguti/night-bridge-oracle/ethereum"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file")
	}
}

func main() {
	conf := config.New()
	ethProvider := ethereum.New(conf.Eth.Endpoint)
	balance := ethProvider.GetBalanceOf("0x8CAB5E96E1ab09e8678a8ffC75b5D818e73D4707")
	fmt.Print(balance)
}
