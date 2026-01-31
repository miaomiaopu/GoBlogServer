package conf

// DefaultConfigPath is the default path to the configuration file.
// 默认配置文件路径
const DefaultConfigPath = "conf/config.yaml"

// Config is the main configuration structure.
// 主配置结构体
// 通过 enter.go 加载
type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	Log      LogConfig      `yaml:"log"`
}
