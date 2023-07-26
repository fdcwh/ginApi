package initialize

import (
	"fmt"
	"goGIn/kernel"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

// GormMysql 初始化Mysql数据库
func GormMysql() {
	c := kernel.FdConfig.Mysql
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", c.Username, c.Password, c.Host, c.Port, c.Database, c.Charset)
	LogLevel := logger.Info

	switch c.LogMode {
	case "silent", "Silent":
		LogLevel = logger.Silent
	case "error", "Error":
		LogLevel = logger.Error
	case "warn", "Warn":
		LogLevel = logger.Warn
	case "info", "Info":
		LogLevel = logger.Info
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  LogLevel,    // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度 【255、191】根据类型配置
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   fmt.Sprintf("%s", c.Prefix),
			SingularTable: true, // 使用单数表名
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 取消外键约束
		Logger:                                   newLogger,
		SkipDefaultTransaction:                   true, // 取消系统默认事务 设为false仍然遵循系统默认事物
	})
	//
	gormDB.InstanceSet("gorm:table_options", "ENGINE=InnoDB")

	if err != nil {
		kernel.FdLog.Error(fmt.Sprintf("连接数据库失败 : %s", err.Error()))
		panic("连接数据库失败: " + err.Error())
	}
	db, err := gormDB.DB()
	if err != nil {
		kernel.FdLog.Error(fmt.Sprintf("连接数据库失败 : %s", err.Error()))
		panic("连接数据库失败: " + err.Error())
	}

	// 连接池
	db.SetMaxIdleConns(c.MaxIdleConns) // 设置空闲的最大连接数
	db.SetMaxOpenConns(c.MaxOpenConns) // 设置与数据库的最大打开连接数
	// db.SetConnMaxLifetime(time.Hour) // 设置可以重用连接的最长时间

	kernel.FdGorm = gormDB
}
