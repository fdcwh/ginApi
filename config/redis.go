package config

type Redis struct {
	Addr string `yaml:"addr"` // 服务器地址:端口
	Db   int    `yaml:"db"`   // 库
	// 当redis服务器版本在6.0以上时，作为ACL认证信息配合密码一起使用，
	// ACL是redis 6.0以上版本提供的认证功能，6.0以下版本仅支持密码认证。
	// 默认为空，不进行认证。
	Username string `yaml:"Username"`
	// 当redis服务器版本在6.0以上时，作为ACL认证信息配合密码一起使用，
	// 当redis服务器版本在6.0以下时，仅作为密码认证。
	// ACL是redis 6.0以上版本提供的认证功能，6.0以下版本仅支持密码认证。
	// 默认为空，不进行认证。
	Password string `yaml:"password"` // 密码
}
