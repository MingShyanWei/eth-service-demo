package router

import (
	root "eth-service-demo/apis"

	v1 "eth-service-demo/apis/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", root.Hello)

	apiV1 := router.Group("/v1")

	apiV1.GET("blocks", v1.ListBlocks)
	apiV1.GET("blocks/:num", v1.GetBlock)
	apiV1.GET("transaction/:txHash", v1.GetTransaction)

	return router
}
