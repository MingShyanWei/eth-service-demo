package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"

	db "eth-service-demo/database"
	model "eth-service-demo/models"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var ctx = context.TODO()

type Indexer struct {
	client *ethclient.Client
}

var indexer *Indexer

type Job interface {
	Process(*ethclient.Client)
}

type Worker struct {
	done             sync.WaitGroup
	readyPool        chan chan Job
	assignedJobQueue chan Job

	quit chan bool
}

type JobQueue struct {
	internalQueue     chan Job
	readyPool         chan chan Job
	workers           []*Worker
	dispatcherStopped sync.WaitGroup
	workersStopped    sync.WaitGroup
	quit              chan bool
}

func NewJobQueue(maxWorkers int) *JobQueue {
	workersStopped := sync.WaitGroup{}
	readyPool := make(chan chan Job, maxWorkers)
	workers := make([]*Worker, maxWorkers, maxWorkers)
	for i := 0; i < maxWorkers; i++ {
		workers[i] = NewWorker(readyPool, workersStopped)
	}
	return &JobQueue{
		internalQueue:     make(chan Job),
		readyPool:         readyPool,
		workers:           workers,
		dispatcherStopped: sync.WaitGroup{},
		workersStopped:    workersStopped,
		quit:              make(chan bool),
	}
}

func (q *JobQueue) Start() {
	for i := 0; i < len(q.workers); i++ {
		q.workers[i].Start()
	}
	go q.dispatch()
}

func (q *JobQueue) Stop() {
	q.quit <- true
	q.dispatcherStopped.Wait()
}

func (q *JobQueue) dispatch() {
	q.dispatcherStopped.Add(1)
	for {
		select {
		case job := <-q.internalQueue:
			workerChannel := <-q.readyPool
			workerChannel <- job
		case <-q.quit:
			for i := 0; i < len(q.workers); i++ {
				q.workers[i].Stop()
			}
			q.workersStopped.Wait()
			q.dispatcherStopped.Done()
			return
		}
	}
}

func (q *JobQueue) Submit(job Job) {
	q.internalQueue <- job
}

func NewWorker(readyPool chan chan Job, done sync.WaitGroup) *Worker {
	return &Worker{
		done:             done,
		readyPool:        readyPool,
		assignedJobQueue: make(chan Job),
		quit:             make(chan bool),
	}
}

func (w *Worker) Start() {

	client, err := ethclient.Dial("https://data-seed-prebsc-2-s3.binance.org:8545/")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		w.done.Add(1)
		for {
			w.readyPool <- w.assignedJobQueue
			select {
			case job := <-w.assignedJobQueue:
				job.Process(client)
			case <-w.quit:
				w.done.Done()
				return
			}
		}
	}()
}

// Stop - stops the worker
func (w *Worker) Stop() {
	w.quit <- true
}

//////////////// Example //////////////////

// GetAndSaveBlock Job
type RPCJob struct {
	num int64
}

// Process - GetAndSaveBlock process function
func (t *RPCJob) Process(client *ethclient.Client) {
	fmt.Printf("Processing job '%s'\n", t.num)
	err := GetAndSaveBlock(client, t.num)
	if err != nil {
		log.Panic(err)
	}
	time.Sleep(1 * time.Second)
}

func GetAndSaveBlock(client *ethclient.Client, num int64) error {

	// -- 1. Check block was existed in db --
	var exists bool
	err := db.GetDb().Model(&model.Block{}).
		Select("count(*) > 0").
		Where("num = ?", num).
		Find(&exists).
		Error
	if err != nil {
		return err
	}

	if exists == true {
		log.Println("block ", num, "was existed.")
		return nil
	}

	// -- 2. Get block from RPC --

	// get block by number
	block, err := client.BlockByNumber(ctx, big.NewInt(num))
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

		chainID, err := client.NetworkID(context.Background())
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

		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
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

func GetLastBlock(client *ethclient.Client) (int64, error) {
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return 0, err
	}

	num := header.Number.Int64()
	fmt.Println("Last block num:", num) // 5671744

	return num, nil
}

func main() {
	db.Init()

	workNum, err := strconv.Atoi(os.Getenv("WORKER_NUM"))
	if err != nil {
		// log.Fatal("WORKER_NUM ", err)
		workNum = runtime.NumCPU()
	}

	fromBlockNum, err := strconv.ParseInt(os.Getenv("FROM_BLOCK_NUM"), 10, 64)
	if err != nil {
		log.Fatal("FROM_BLOCK_NUM ", err)
	}

	rawurl := os.Getenv("RPC_URL")

	client, err := ethclient.Dial(rawurl)
	if err != nil {
		log.Fatal(err)
	}
	// create an indexer
	indexer = &Indexer{
		client,
	}

	// ---

	// queue := NewJobQueue(runtime.NumCPU())
	queue := NewJobQueue(workNum)
	queue.Start()
	defer queue.Stop()

	var i int64
	for {
		m, err := GetLastBlock(client)
		if err != nil {
			log.Panic(err)
		}

		for i = fromBlockNum; i < m; i++ {
			queue.Submit(&RPCJob{i})
		}

		log.Println("Wait 60 seconds...")
		time.Sleep(60 * time.Second)
	}

}
