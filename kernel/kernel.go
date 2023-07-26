package kernel

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"goGIn/config"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

const (
	ConfigEnv         = "FD_CONF"
	ConfigDefaultFile = "config.yaml"
	ConfigTestFile    = "config.test.yaml"
	ConfigReleaseFile = "config.release.yaml"
)

var (
	FdConfig             config.Server
	FdGorm               *gorm.DB
	FdRedis              *redis.Client
	FdViper              *viper.Viper
	FdLog                *zap.Logger
	FdTrans              ut.Translator
	FdConcurrencyControl = &singleflight.Group{}
)
