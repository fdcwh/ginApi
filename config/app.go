package config

type App struct {
	Name          string `yaml:"name"`
	Addr          string `yaml:"addr"`           // 端口值
	UseMultipoint bool   `yaml:"use-multipoint"` // 多点登录拦截
	LimitCountIP  int    `yaml:"iplimit-count"`
	LimitTimeIP   int    `yaml:"iplimit-time"`
	RouterPrefix  string `yaml:"router-prefix"`
}
