package Controller

import (
	"github.com/6QHTSK/ayachan-bestdoriAPI/models"
	"github.com/6QHTSK/ayachan-bestdoriAPI/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetChartInfo(c *gin.Context) {
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
	chartInfo, errorCode, err := service.FetchChartInfo(chartID, models.DiffType(diff))
	if err != nil {
		c.JSON(errorCode, gin.H{
			"result":  false,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"info":   chartInfo,
	})
}

func GetChartList(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "0")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":  false,
			"message": err.Error(),
		})
		return
	}
	limitStr := c.DefaultQuery("limit", "20")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":  false,
			"message": err.Error(),
		})
		return
	}
	count, list, errorCode, err := service.FetchChartList(page, limit)
	if err != nil {
		c.JSON(errorCode, gin.H{
			"result":  false,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"count":  count,
		"list":   list,
	})
}
