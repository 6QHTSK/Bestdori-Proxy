package router

import (
	"github.com/6QHTSK/Bestdori-Proxy/controller"
	"github.com/6QHTSK/Bestdori-Proxy/errors"
	"github.com/6QHTSK/Bestdori-Proxy/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() (router *gin.Engine) {
	router = gin.Default()
	router.Use(cors.Default(), middleware.ErrorHandler, middleware.AddVersionToHeader)
	postGroup := router.Group("/post", middleware.ParamHelperPostInfo)
	{
		postGroup.GET("/:server/:postID", controller.PostInfoHandler)
		postGroup.GET("/:server/:postID/:method", controller.PostInfoHandler)
	}
	postListGroup := router.Group("/post", middleware.ParamHelperPostList)
	{
		postListGroup.GET("/:server/list", controller.PostListHandler)
	}
	assetsGroup := router.Group("/assets", middleware.ParamHelperAsset)
	{
		assetsGroup.GET("/:server/:postID/cover", controller.CoverProxy)
		assetsGroup.GET("/:server/:postID/audio", controller.AudioProxy)
	}
	router.NoRoute(func(context *gin.Context) {
		_ = context.Error(errors.NoRouteErr)
	})
	return router
}
