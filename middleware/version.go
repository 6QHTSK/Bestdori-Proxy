package middleware

import (
	"github.com/6QHTSK/Bestdori-Proxy/config"
	"github.com/gin-gonic/gin"
)

func AddVersionToHeader(ctx *gin.Context) {
	ctx.Writer.Header().Set("Bestdori-Proxy-Version", config.Version)
}
