package core

import (
	"Server/conf"
	"Server/logx"
)

// InitLogger 初始化全局日志
func InitLogger(cfg conf.LogConfig) (*logx.Logger, func(), error) {
	// 调用 logx 包的 Init 方法初始化日志
	logger, cleanup, err := logx.Init(cfg)
	if err != nil {
		return nil, nil, err
	}
	return logger, cleanup, nil
}
