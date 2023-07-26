package initialize

import (
	"context"
	"github.com/redis/go-redis/v9"
	"goGIn/kernel"

	"go.uber.org/zap"
)

// InitClientRedis TODO:: 文档 https://redis.uptrace.dev/zh/
// InitClientRedis 基础单客户端连接方式、 可调整为 UniversalClient 或增加其他连接方式
func InitClientRedis() {
	rCfg := kernel.FdConfig.Redis
	rdbClient := redis.NewClient(&redis.Options{
		Addr:     rCfg.Addr,
		Password: rCfg.Password,
		Username: rCfg.Username,
		DB:       rCfg.Db,
	})
	pong, err := rdbClient.Ping(context.Background()).Result()
	if err != nil {
		kernel.FdLog.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		kernel.FdLog.Info("redis connect ping response:", zap.String("pong", pong))
		kernel.FdRedis = rdbClient
	}
}
