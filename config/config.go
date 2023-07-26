package config

type Server struct {
	JWT   JWT   `yaml:"jwt"`
	Log   Log   `yaml:"log"`
	Redis Redis `yaml:"redis"`
	App   App   `yaml:"app"`
	// Captcha Captcha `yaml:"captcha"`
	// gorm
	Mysql Mysql `yaml:"mysql"`
	// store
	Local Local `yaml:"local"`
}
