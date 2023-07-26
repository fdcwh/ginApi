package api

import (
	"github.com/gin-gonic/gin"
	"goGIn/app/api"
)

func InitPublicRouter(Router *gin.Engine) {
	publicRouter := Router.Group("public")

	publicRouter.GET("config", api.GetConfigData)

	//
}
