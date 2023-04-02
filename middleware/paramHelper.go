package middleware

import (
	"Bestdori-Proxy/errors"
	"Bestdori-Proxy/models"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

var methodMap = map[string]int{"": models.MethodFull, "full": models.MethodFull, "info": models.MethodInfo, "chart": models.MethodChart}
var serverMap = map[string]int{"bandori": models.ServerBandori, "bestdori": models.ServerBestdori, "llsif": models.ServerLLSif}

func abortWhenErr(ctx *gin.Context, err error, myError *errors.ProxyError) bool {
	if err != nil {
		_ = ctx.Error(myError)
		ctx.Abort()
		return true
	}
	return false
}

func abortWhenFalse(ctx *gin.Context, flag bool, myError *errors.ProxyError) bool {
	if !flag {
		_ = ctx.Error(myError)
		ctx.Abort()
		return true
	}
	return false
}

func postIDParser(ctx *gin.Context) {
	strID := ctx.Param("postID")
	postID, err := strconv.Atoi(strID)
	if abortWhenErr(ctx, err, errors.PostIDParseErr) {
		return
	}
	ctx.Set("postID", postID)
}

func serverParser(ctx *gin.Context) {
	strServer := ctx.Param("server")
	strServer = strings.ToLower(strServer)
	server, ok := serverMap[strServer]
	if abortWhenFalse(ctx, ok, errors.UnknownServerErr) {
		return
	}
	ctx.Set("server", server)
}

func ParamHelperPostInfo(ctx *gin.Context) {
	serverParser(ctx)
	postIDParser(ctx)

	diffStr := ctx.DefaultQuery("diff", "3")
	diff, err := strconv.Atoi(diffStr)
	if abortWhenErr(ctx, err, errors.DiffParseErr) {
		return
	}
	ctx.Set("diff", diff)

	strMethod := ctx.Param("method")
	method, ok := methodMap[strMethod]
	if abortWhenFalse(ctx, ok, errors.MethodParseErr) {
		return
	}
	ctx.Set("method", method)
}

func ParamHelperPostList(ctx *gin.Context) {
	serverParser(ctx)

	username := ctx.DefaultQuery("username", "")
	username = strings.ToLower(username)
	ctx.Set("username", username)

	offsetStr := ctx.DefaultQuery("offset", "0")
	offset, err := strconv.ParseUint(offsetStr, 10, 32)
	if abortWhenErr(ctx, err, errors.OffsetParseErr) {
		return
	}
	ctx.Set("offset", offset)

	limitStr := ctx.DefaultQuery("limit", "20")
	limit, err := strconv.ParseUint(limitStr, 10, 32)
	if abortWhenErr(ctx, err, errors.LimitParseErr) {
		return
	}
	ctx.Set("limit", limit)
}

func ParamHelperAsset(ctx *gin.Context) {
	serverParser(ctx)
	postIDParser(ctx)
}
