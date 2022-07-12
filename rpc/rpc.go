package database

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

var client *ethclient.Client

func Init(rpcUrl string) {
	var err error

	// client, err = ethclient.Dial("https://data-seed-prebsc-2-s3.binance.org:8545/")
	client, err = ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatal(err)
	}

	lastBlock, err := GetLastBlock(client)
	if err != nil {
		log.Panic(err)
	}

	log.Println("lastBlock:", lastBlock)
}

func GetClient() *ethclient.Client {
	return client
}

func GetLastBlock(client *ethclient.Client) (int64, error) {
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return 0, err
	}

	num := header.Number.Int64()
	fmt.Println("Last block num:", num) // 5671744

	return num, nil
}
