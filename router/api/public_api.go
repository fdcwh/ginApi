package api

import (
	"github.com/gin-gonic/gin"
	"goGIn/app/api"
	"goGIn/kernel/middleware"
)

func InitPublicRouter(Router *gin.Engine) {
	publicRouter := Router.Group("public")

	publicRouter.GET("config", api.GetConfigData)

	//

	Router.GET("setJwt", api.JwtSet)

	Router.Use(middleware.Jwt())
	{
		Router.GET("getJwt", api.JwtGet)

		Router.GET("refreshToken", api.RefreshToken)
	}
}
