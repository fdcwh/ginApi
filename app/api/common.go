package api

import (
	"github.com/gin-gonic/gin"
	"goGIn/kernel"
	"goGIn/kernel/res"
)

func GetConfigData(c *gin.Context) {
	res.Success.SetData(gin.H{
		"FdConfig": kernel.FdConfig,
	}).ToJson(c)
}
