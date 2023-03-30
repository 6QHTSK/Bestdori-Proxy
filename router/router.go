package router

import (
	"Bestdori-Proxy/controller"
	"Bestdori-Proxy/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() (router *gin.Engine) {
	router = gin.Default()
	router.Use(cors.Default(), middleware.ErrorHandler())
	bestdori := router.Group("/bestdori")
	{
		bestdori.GET("/chart", controller.GetChartList)
		bestdori.GET("/chart/:chartID", controller.GetChartInfo)
		bestdori.GET("/chart/:chartID/:method", controller.GetChartInfo)
		bestdori.GET("/cover/:chartID", controller.CoverProxy)
		bestdori.GET("/music/:chartID", controller.MusicProxy)
	}
	router.GET("/sonolus/*path", controller.SonolusProxy)
	return router
}
