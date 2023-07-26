package config

type Local struct {
	Path      string `yaml:"path"`       // 本地文件访问路径
	StorePath string `yaml:"store-path"` // 本地文件存储路径
}
