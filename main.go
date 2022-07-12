package main

import (
	_ "eth-service-demo/database"
	"eth-service-demo/router"
	"os"

	db "eth-service-demo/database"
)

func main() {
	dsn := os.Getenv("DB_CONNECTION")
	db.Init(dsn)

	router := router.InitRouter()
	router.Run(":8000")
}
