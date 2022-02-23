package Controller

import (
	"github.com/6QHTSK/ayachan-bestdoriAPI/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetEngine(c *gin.Context) {
	result, errorCode, err := service.FetchEngine()
	if err != nil {
		c.JSON(errorCode, gin.H{
			"result": false,
			"error":  err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"engine": result,
	})
}
