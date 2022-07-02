package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	db "eth-service-demo/database"
	model "eth-service-demo/models"
)

var ctx = context.TODO()

type Indexer struct {
	client *ethclient.Client
}

func main() {
	// db.Init()

	client, err := ethclient.Dial("https://data-seed-prebsc-2-s3.binance.org:8545/")
	if err != nil {
		log.Fatal(err)
	}
	// create an indexer
	indexer := &Indexer{
		client,
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Last block num:", header.Number.String()) // 5671744

	// start the indexer
	indexer.Start(20683000, 20683100)

	// ERROR: 20683001 0xdae86a289a4bee8c9ff24e300e264db7fb65435d7f8a145857a7987a1b89a06a hasn't to!!!
}

func (indexer *Indexer) Start(from int64, to int64) error {

	if from <= to {
		for i := from; i <= to; i++ {
			log.Println("i:", i)
			err := indexer.GetBlock(i)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (indexer *Indexer) GetBlock(num int64) error {

	// -- 1. Check block was existed in db --

	// -- 2. Get block from RPC --

	// get block by number
	block, err := indexer.client.BlockByNumber(ctx, big.NewInt(num))
	if err != nil {
		return err
	}

	log.Println("block num:", block.Number().Uint64()) // bk_num 20597939
	// log.Println(block.Time())             // bk_time 1656437007
	// log.Println(block.Hash().Hex())       // bk_hash 0x4a5aad19c3b7b375852c787c66f1eb86a6ba3e359164966782a2456c99bc794d
	// log.Println(block.ParentHash().Hex()) // bk_parenthash 0xe9f64982c3fd8c1caa6cb8c68b8cf75cd5baa72223009fd426879b3e7fa3ed3b

	var newBlock = model.Block{
		Num:        block.Number().Uint64(),
		Hash:       block.Hash().String(),
		ParentHash: block.ParentHash().String(),
		Time:       block.Time(),
	}
	// db.GetDb().Create(&newBlock)

	// log.Println(block.Transactions())

	var newTransactions []model.Transaction
	var newTransactionLogs []model.TransactionLog

	for _, tx := range block.Transactions() {
		log.Println(block.Number().Uint64(), "/", tx.Hash().Hex()) // tx_hash 0xd7d9c32699fabd278e9d5f1119d7bfcee07f778c1314940f511b5385e6b30c12

		chainID, err := indexer.client.NetworkID(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		var msg types.Message
		if msg, err = tx.AsMessage(types.NewEIP155Signer(chainID), big.NewInt(1)); err == nil {
			// log.Println(msg.From().Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
		}

		var to string
		if tx.To() != nil {
			to = tx.To().String()
			// log.Println(tx.To().Hex()) // tx_to 0x55fE59D8Ad77035154dDd0AD0388D09Dd4047A8e
		} else {
			to = ""
		}

		// log.Println(tx.Nonce())                            // tx_nonce 110644
		// log.Println(string(hex.EncodeToString(tx.Data()))) // tx_data []
		// log.Println(tx.Value().String())                   // tx_value 200000000000000000

		var newTransaction = model.Transaction{
			TxHash: tx.Hash().String(),
			Num:    block.Number().Uint64(),
			From:   msg.From().String(),
			To:     to,
			Nonce:  tx.Nonce(),
			Data:   "0x" + string(hex.EncodeToString(tx.Data())),
			Value:  tx.Value().String(),
		}
		newTransactions = append(newTransactions, newTransaction)
		// db.GetDb().Create(&newTransaction)

		receipt, err := indexer.client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		if receipt.Status == 1 {
			for _, receiptLog := range receipt.Logs {
				// log.Println(receiptLog.Index)
				// log.Println(receiptLog.Data)
				var newTransactionLog = model.TransactionLog{
					TxHash: tx.Hash().String(),
					Index:  receiptLog.Index,
					Data:   "0x" + string(hex.EncodeToString(receiptLog.Data)),
				}

				newTransactionLogs = append(newTransactionLogs, newTransactionLog)
				// db.GetDb().Create(&newLog)

			}
		}

	}

	// -- 3. Save block into db --
	log.Println("save block num:", block.Number().Uint64())
	db.GetDb().Create(&newBlock)
	db.GetDb().Create(&newTransactions)
	db.GetDb().Create(&newTransactionLogs)

	return nil
}
