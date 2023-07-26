package config

type JWT struct {
	SigningKey  string `yaml:"signingKey"`  // jwt签名
	ExpiresTime string `yaml:"expiresTime"` // 过期时间
	BufferTime  string `yaml:"buffer-time"` // 缓冲时间
	Issuer      string `yaml:"issuer"`      // 签发者
}
