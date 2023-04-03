package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
)
import "github.com/6QHTSK/Bestdori-Proxy/errors"

func ErrorHandler(context *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			debug.PrintStack()
			context.JSON(http.StatusInternalServerError, errors.UnknownErr)
		}
	}()
	context.Next()
	for _, e := range context.Errors {
		err := e.Err
		if ProxyError, ok := err.(*errors.ProxyError); ok {
			context.JSON(ProxyError.HttpCode, ProxyError)
		} else {
			errors.UnknownErr.ErrMsg = err.Error()
			context.JSON(http.StatusInternalServerError, errors.UnknownErr)
		}
	}
}
