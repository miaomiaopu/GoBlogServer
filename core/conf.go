package core

import (
	"Server/conf"
	"Server/util"
	"fmt"

	"gopkg.in/yaml.v3"
)

// InitConfig 初始化全局配置
func InitConfig(path string) (*conf.Config, error) {
	if path == "" {
		path = conf.DefaultConfigPath
	}

	data, err := util.LoadYaml(path)
	if err != nil {
		return nil, fmt.Errorf("read config %s: %w", path, err)
	}

	var cfg conf.Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}

	return &cfg, nil
}
