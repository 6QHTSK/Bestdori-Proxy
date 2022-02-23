package Controller

import (
	"github.com/6QHTSK/ayachan-bestdoriAPI/models"
	"github.com/6QHTSK/ayachan-bestdoriAPI/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetMap(c *gin.Context) {
	strID := c.Param("chartID")
	chartID, err := strconv.Atoi(strID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":  false,
			"message": err.Error(),
		})
		return
	}
	diffStr := c.DefaultQuery("diff", "3")
	diff, err := strconv.Atoi(diffStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":  false,
			"message": err.Error(),
		})
		return
	}
	chartMap, errorCode, err := service.FetchMap(chartID, models.DiffType(diff))
	if err != nil {
		c.JSON(errorCode, gin.H{
			"result":  false,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"map":    chartMap,
	})
}
