package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	db "eth-service-demo/database"
	model "eth-service-demo/models"
)

var ctx = context.TODO()

type indexer struct {
	client *ethclient.Client
}

func main() {
	db.Init()

	client, err := ethclient.Dial("https://data-seed-prebsc-2-s3.binance.org:8545/")
	if err != nil {
		panic(err)
	}
	// create an indexer
	indexer := &indexer{
		client,
	}
	// start the indexer
	indexer.Start(20597920, 20597939)
}

func (indexer *indexer) Start(from int64, to int64) error {

	start := big.NewInt(from)
	end := big.NewInt(to)

	for i := new(big.Int).Set(start); i.Cmp(end) <= 0; i.Add(i, big.NewInt(1)) {
		// get block by number
		block, err := indexer.client.BlockByNumber(ctx, i)
		if err != nil {
			return err
		}

		fmt.Println(block.Number().Uint64())  // bk_num 20597939
		fmt.Println(block.Time())             // bk_time 1656437007
		fmt.Println(block.Hash().Hex())       // bk_hash 0x4a5aad19c3b7b375852c787c66f1eb86a6ba3e359164966782a2456c99bc794d
		fmt.Println(block.ParentHash().Hex()) // bk_parenthash 0xe9f64982c3fd8c1caa6cb8c68b8cf75cd5baa72223009fd426879b3e7fa3ed3b

		var newBlock = model.Block{
			Num:        block.Number().Uint64(),
			Hash:       block.Hash().String(),
			ParentHash: block.ParentHash().String(),
			Time:       block.Time(),
		}
		db.GetDb().Create(&newBlock)

		fmt.Println(block.Transactions())

		for _, tx := range block.Transactions() {
			fmt.Println(tx.Hash().Hex()) // tx_hash 0xd7d9c32699fabd278e9d5f1119d7bfcee07f778c1314940f511b5385e6b30c12

			chainID, err := indexer.client.NetworkID(context.Background())
			if err != nil {
				return err
			}

			var msg types.Message
			if msg, err = tx.AsMessage(types.NewEIP155Signer(chainID), big.NewInt(1)); err == nil {
				fmt.Println(msg.From().Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
			}

			fmt.Println(tx.To().Hex()) // tx_to 0x55fE59D8Ad77035154dDd0AD0388D09Dd4047A8e

			fmt.Println(tx.Nonce())                            // tx_nonce 110644
			fmt.Println(string(hex.EncodeToString(tx.Data()))) // tx_data []
			fmt.Println(tx.Value().String())                   // tx_value 200000000000000000

			var newTransaction = model.Transaction{
				TxHash: tx.Hash().String(),
				Num:    block.Number().Uint64(),
				From:   msg.From().String(),
				To:     tx.To().String(),
				Nonce:  tx.Nonce(),
				Data:   "0x" + string(hex.EncodeToString(tx.Data())),
				Value:  tx.Value().String(),
			}
			db.GetDb().Create(&newTransaction)

		}

	}
	return nil
}
