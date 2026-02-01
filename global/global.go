package global

import (
	"Server/conf"

	"gorm.io/gorm"
)

// Logger 是全局日志接口，避免依赖具体实现，避免循环依赖
type Logger interface {
	Debug(args ...any)
	Info(args ...any)
	Warn(args ...any)
	Error(args ...any)
	Fatal(args ...any)

	Debugw(msg string, keysAndValues ...any)
	Infow(msg string, keysAndValues ...any)
	Warnw(msg string, keysAndValues ...any)
	Errorw(msg string, keysAndValues ...any)
	Fatalw(msg string, keysAndValues ...any)
}

var GlobalConfig *conf.Config
var GlobalLogger Logger
var GlobalDB *gorm.DB
