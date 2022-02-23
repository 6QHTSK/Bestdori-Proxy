package router

import (
	"github.com/6QHTSK/ayachan-bestdoriAPI/Controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() (router *gin.Engine) {
	router = gin.Default()
	router.Use(cors.Default())
	router.GET("/:chartID/map", Controller.GetMap)
	router.GET("/:chartID", Controller.GetChartInfo)
	router.GET("/list", Controller.GetChartList)
	router.GET("/engine", Controller.GetEngine)
	return router
}
