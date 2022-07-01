package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "Hello",
	})
	return
}
