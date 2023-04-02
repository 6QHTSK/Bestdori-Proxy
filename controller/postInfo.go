package controller

import (
	"Bestdori-Proxy/models"
	"Bestdori-Proxy/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	bandoriEasy = 0
	// bandoriNormal
	// bandoriHard
	bandoriExpert  = 3
	bandoriSpecial = 4
)

var postInfoHandlerMap = map[int]gin.HandlerFunc{
	models.ServerBandori:  bandoriPostInfo,
	models.ServerBestdori: bestdoriPostInfo,
	models.ServerLLSif:    unsupported,
}

func PostInfoHandler(ctx *gin.Context) {
	postInfoHandlerMap[ctx.GetInt("server")](ctx)
}

func bandoriPostInfo(ctx *gin.Context) {
	postID := ctx.GetInt("postID")
	method := ctx.GetInt("method")
	diff := ctx.GetInt("diff")
	if diff <= bandoriEasy || diff >= bandoriSpecial {
		diff = bandoriExpert
	}
	PostInfo, err := service.FetchBandoriPost(postID, diff, method)
	if abortWhenErr(ctx, err, err) {
		return
	}
	ctx.JSON(http.StatusOK, PostInfo)
}

func bestdoriPostInfo(ctx *gin.Context) {
	postID := ctx.GetInt("postID")
	method := ctx.GetInt("method")
	PostInfo, err := service.FetchBestdoriPost(postID, method)
	if abortWhenErr(ctx, err, err) {
		return
	}
	ctx.JSON(http.StatusOK, PostInfo)
}
