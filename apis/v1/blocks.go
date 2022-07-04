package apis

import (
	model "eth-service-demo/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListBlocks(c *gin.Context) {

	var block model.Block

	limit, err := strconv.Atoi(c.Request.FormValue("limit"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "limit atoi error",
		})
		return
	}

	result, err := block.ListBlocks(limit)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": ",list block error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"blocks": result,
	})
}

func GetBlock(c *gin.Context) {

	var blockWithTranscations model.BlockWithTranscations

	num, err := strconv.ParseInt(c.Param("num"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "num atoi error",
		})
		return
	}

	result, err := blockWithTranscations.GetBlockDetail(num)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "GetBlockDetail error",
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
