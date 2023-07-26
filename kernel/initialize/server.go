package initialize

import (
	"context"
	"github.com/gin-gonic/gin"
	"goGIn/kernel"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func RunServer(Router *gin.Engine) {
	// 服务启动
	server := &http.Server{
		Addr:    kernel.FdConfig.App.Addr,
		Handler: Router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Println("\n// ============================================== "+
		"\n// |  服务启动成功"+
		"\n// |  ------------------------------------------"+
		"\n// | SERVER-NAME: ", kernel.FdConfig.App.Name,
		"\n// | SERVER-ADDR: ", kernel.FdConfig.App.Addr,
		"\n// ==============================================\n ")

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("停止服务 ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("服务退出")
}
