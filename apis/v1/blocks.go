package apis

import (
	cache "eth-service-demo/cache"
	model "eth-service-demo/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListBlocks(c *gin.Context) {

	var block model.Block
	var result []model.Block

	limit, err := strconv.Atoi(c.Request.FormValue("limit"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "limit atoi error",
		})
		return
	}

	log.Println(len(cache.Blocks))

	if len(cache.Blocks) == 0 {
		result, err = block.ListBlocks(limit)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    -1,
				"message": ",list block error",
			})
			return
		}
	} else {
		result = cache.Blocks[0 : limit-1]
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
