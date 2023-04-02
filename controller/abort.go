package controller

import (
	"Bestdori-Proxy/errors"
	"github.com/gin-gonic/gin"
)

func unsupported(ctx *gin.Context) {
	_ = ctx.Error(errors.UnsupportedHandler)
	ctx.Abort()
}

func abortWhenErr(ctx *gin.Context, err error, myError error) bool {
	if err != nil {
		_ = ctx.Error(myError)
		ctx.Abort()
		return true
	}
	return false
}
