package main

import (
	"github.com/gin-gonic/gin"
	"goGIn/kernel/initialize"
	"goGIn/router"
)

func main() {
	gin.SetMode(gin.DebugMode) // Debug
	// gin.SetMode(gin.ReleaseMode)
	// 配置
	initialize.InitViper()
	// 日志
	initialize.InitLog()
	//
	initialize.InitValidator("zh")
	// redis
	initialize.InitClientRedis()
	// GormMysql
	initialize.GormMysql()

	Router := router.RoutersInit()

	initialize.RunServer(Router)
}
