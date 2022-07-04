package apis

import (
	model "eth-service-demo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTransaction(c *gin.Context) {

	var transaction model.Transaction

	txHash := c.Param("txHash")

	result, err := transaction.GetTransaction(txHash)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "GetTransaction error",
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
