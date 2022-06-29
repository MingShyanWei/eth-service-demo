package apis

import (
	model "eth-service-demo/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTransaction(c *gin.Context) {

	var transaction model.Transaction

	txHash := c.Param("txHash")

	fmt.Println(txHash)

	result, err := transaction.GetTransaction(txHash)

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
