package main

import (
	_ "eth-service-demo/database"
	"eth-service-demo/router"

	db "eth-service-demo/database"
)

func main() {
	db.Init()

	router := router.InitRouter()
	router.Run(":8000")
}
