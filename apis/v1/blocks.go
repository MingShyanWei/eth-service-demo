package apis

import (
	model "eth-service-demo/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListBlocks(c *gin.Context) {

	var block model.Block

	limit, err := strconv.Atoi(c.Request.FormValue("limit"))
	if err != nil {
		// handle the error in some way
	}

	fmt.Println(limit)

	result, err := block.ListBlocks(limit)

	fmt.Println(result)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "error",
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
		panic(err)
	}

	fmt.Println(num)

	result, err := blockWithTranscations.GetBlockDetail(num)

	fmt.Println(result)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "error",
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
