package main

import (
	db "eth-service-demo/database"
	model "eth-service-demo/models"
)

func main() {
	db.Init()

	db.GetDb().AutoMigrate(&model.Block{}, &model.Transaction{})
}
