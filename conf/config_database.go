package conf

import "strconv"

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

func (dbConfig *DatabaseConfig) DSN() string {
	// 构建 PostgreSQL 连接字符串
	return "host=" + dbConfig.Host +
		" user=" + dbConfig.User +
		" password=" + dbConfig.Password +
		" dbname=" + dbConfig.Name +
		" port=" + strconv.Itoa(dbConfig.Port) +
		" search_path=" + dbConfig.Schema +
		" sslmode=disable"
}
