package api

import (
	"github.com/gin-gonic/gin"
	"goGIn/app/api/article"
)

func InitArticleRouter(Router *gin.Engine) {
	publicRouter := Router.Group("article")

	publicRouter.GET("AddTag", article.AddTag)

	//
}
