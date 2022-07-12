package main

import (
	_ "eth-service-demo/database"
	"eth-service-demo/router"
	"os"

	cache "eth-service-demo/cache"
	db "eth-service-demo/database"
)

func main() {
	dsn := os.Getenv("DB_CONNECTION")
	db.Init(dsn)

	go cache.UpdateBlocksCashe()

	router := router.InitRouter()
	router.Run(":8000")
}
