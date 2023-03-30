package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
import "Bestdori-Proxy/errors"

func ErrorHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
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
}
