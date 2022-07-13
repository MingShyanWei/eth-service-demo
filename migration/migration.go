package main

import (
	db "eth-service-demo/database"
	model "eth-service-demo/models"
	"os"
)

func main() {
	dsn := os.Getenv("DB_CONNECTION")
	db.Init(dsn)

	db.GetDb().AutoMigrate(&model.Block{}, &model.Transaction{}, &model.TransactionLog{})
}
