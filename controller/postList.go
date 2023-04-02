package controller

import (
	"Bestdori-Proxy/models"
	"Bestdori-Proxy/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var postListHandlerMap = map[int]gin.HandlerFunc{
	models.ServerBandori:  bandoriPostList,
	models.ServerBestdori: bestdoriPostList,
	models.ServerLLSif:    unsupported,
}

func PostListHandler(ctx *gin.Context) {
	postListHandlerMap[ctx.GetInt("server")](ctx)
}

func bandoriPostList(ctx *gin.Context) {
	offset := ctx.GetUint64("offset")
	limit := ctx.GetUint64("limit")
	count, list, err := service.FetchBandoriPostList(offset, limit)
	if abortWhenErr(ctx, err, err) {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"count": count,
		"list":  list,
	})
}

func bestdoriPostList(ctx *gin.Context) {
	offset := ctx.GetUint64("offset")
	limit := ctx.GetUint64("limit")
	username := ctx.GetString("username")
	count, list, err := service.FetchBestdoriPostList(offset, limit, username)
	if abortWhenErr(ctx, err, err) {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"count": count,
		"list":  list,
	})
}
