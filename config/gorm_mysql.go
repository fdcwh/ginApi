package config

type Mysql struct {
	Host         string `yaml:"host"`     // 服务器地址
	Port         string `yaml:"port"`     //:端口
	Database     string `yaml:"database"` // 数据库名
	Username     string `yaml:"username"` // 数据库用户名
	Password     string `yaml:"password"` // 数据库密码
	Prefix       string `yaml:"prefix"`   //全局表前缀，单独定义TableName则不生效
	Charset      string `yaml:"charset"`
	Singular     bool   `yaml:"singular"`                //是否开启全局禁用复数，true表示开启
	Engine       string `yaml:"engine" default:"InnoDB"` //数据库引擎，默认InnoDB
	MaxIdleConns int    `yaml:"maxIdleConns"`            // 空闲中的最大连接数
	MaxOpenConns int    `yaml:"maxOpenConns"`            // 打开到数据库的最大连接数
	LogMode      string `yaml:"logMode"`                 // 日志 Silent、Error、Warn、Info
}
