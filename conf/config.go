package conf

import (
	"os"

	"gopkg.in/yaml.v3"
)

// ServerConfig holds server-related configuration.
// 服务器配置结构体
type ServerConfig struct {
	Port int `yaml:"port"`
}

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

// Config is the main configuration structure.
// 主配置结构体
type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}

// LoadConfig reads the configuration from a YAML file.
// 读取配置文件
func LoadConfig(path string) (*Config, error) {
	if path == "" {
		path = "conf/config.yaml"
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	// Unmarshal the YAML data into the Config struct
	// 将YAML数据解码到Config结构体中
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
