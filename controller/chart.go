package controller

import (
	"Bestdori-Proxy/errors"
	"Bestdori-Proxy/models"
	"Bestdori-Proxy/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

var methodMap = map[string]models.WorkMethod{"": models.FULL, "info": models.INFO, "chart": models.CHART}

func GetChartInfo(c *gin.Context) {
	strID := c.Param("chartID")
	chartID, err := strconv.Atoi(strID)
	if err != nil {
		_ = c.Error(errors.ChartIDParseErr)
		return
	}
	diffStr := c.DefaultQuery("diff", "3")
	diff, err := strconv.Atoi(diffStr)
	if err != nil {
		_ = c.Error(errors.DiffParseErr)
		return
	}
	_, official := c.GetQuery("official")
	strMethod := c.Param("method")
	method, ok := methodMap[strMethod]
	if !ok {
		_ = c.Error(errors.MethodParseErr)
		return
	}
	chartInfo, err := service.FetchChartInfo(chartID, diff, official, method)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, chartInfo)
}

func GetChartList(c *gin.Context) {
	username := c.DefaultQuery("username", "")
	username = strings.ToLower(username)
	offsetStr := c.DefaultQuery("offset", "0")
	offset, err := strconv.ParseUint(offsetStr, 10, 32)
	if err != nil {
		_ = c.Error(errors.OffsetParseErr)
		return
	}
	limitStr := c.DefaultQuery("limit", "20")
	limit, err := strconv.ParseUint(limitStr, 10, 32)
	if err != nil {
		_ = c.Error(errors.LimitParseErr)
		return
	}
	count, list, err := service.FetchChartList(uint(offset), uint(limit), username)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"count": count,
		"list":  list,
	})
}
