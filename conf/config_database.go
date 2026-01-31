package conf

// DatabaseConfig holds database-related configuration.
// 数据库配置结构体
type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Schema   string `yaml:"schema"`
}
