package router

import (
	v1 "api/apis/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	apiV1 := router.Group("/v1")

	apiV1.GET("blocks", v1.ListBlocks)
	apiV1.GET("blocks/:num", v1.GetBlock)

	return router
}
