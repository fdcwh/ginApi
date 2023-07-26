package router

import (
	"github.com/gin-gonic/gin"
	"goGIn/kernel/exception"
	"goGIn/kernel/middleware"
	"goGIn/kernel/res"
	"goGIn/router/api"
)

func RoutersInit() *gin.Engine {
	Router := gin.Default()
	// 跨域
	Router.Use(middleware.Cors())

	// 异常
	Router.NoMethod(exception.HandleNotFound)
	Router.NoRoute(exception.HandleNotFound)

	// 统一前缀

	//

	//
	Router.GET("/", func(c *gin.Context) {
		res.Success.SetData(gin.H{
			"message": "ok",
		}).ToJson(c)
	})

	//
	api.InitPublicRouter(Router)
	api.InitArticleRouter(Router)

	return Router
}
