package conf

// ServerConfig holds server-related configuration.
// 服务器配置结构体
type ServerConfig struct {
	Port    int    `yaml:"port"`
	GinMode string `yaml:"gin_mode"`
}
