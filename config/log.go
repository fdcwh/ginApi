package config

// Log  配置文件结构体
type Log struct {
	BaseDirectoryName  string `yaml:"BaseDirectoryName" default:"logs"`
	InfoDirectoryName  string `yaml:"InfoDirectoryName" default:"info"`
	WarnDirectoryName  string `yaml:"WarnDirectoryName" default:"warn"`
	ErrorDirectoryName string `yaml:"ErrorDirectoryName" default:"error"`
	DebugDirectoryName string `yaml:"DebugDirectoryName" default:"debug"`
	InfoFileName       string `yaml:"InfoFileName" default:"info.log"`
	WarnFileName       string `yaml:"WarnFileName" default:"warn.log"`
	ErrorFileName      string `yaml:"ErrorFileName" default:"error.log"`
	DebugFileName      string `yaml:"DebugFileName" default:"Debug.log"`
	LogFileMaxSize     int    `yaml:"LogFileMaxSize" default:"128"`
	LogFileMaxBackups  int    `yaml:"LogFileMaxBackups" default:"180"`
	LogFileMaxAge      int    `yaml:"LogFileMaxAge" default:"1"`
	LogFileCompress    bool   `yaml:"LogFileCompress" default:"false"`
	LogConsoleOut      bool   `yaml:"LogConsoleOut"  default:"false"`
}
