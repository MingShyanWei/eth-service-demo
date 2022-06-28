package main

import (
	_ "api/database"
	"api/router"
)

func main() {

	router := router.InitRouter()
	router.Run(":8000")
}
