package logx

import (
	"Server/conf"
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	sugar *zap.SugaredLogger
}

// 初始化日志
func Init(cfg conf.LogConfig) (*Logger, func(), error) {
	var sugar *zap.SugaredLogger

	level := zapcore.InfoLevel
	if err := level.UnmarshalText([]byte(cfg.Level)); err != nil {
		return nil, nil, fmt.Errorf("invalid log level: %w", err)
	}

	// 解析输出目标
	hasStdout := false
	hasFile := false
	if strings.TrimSpace(cfg.Outputs) == "" {
		hasStdout = true
	} else {
		for p := range strings.SplitSeq(cfg.Outputs, ",") {
			switch strings.TrimSpace(p) {
			case "stdout":
				hasStdout = true
			case "file":
				hasFile = true
			}
		}
	}

	// 默认日志文件路径
	if hasFile {
		if cfg.AppLogPath == "" {
			cfg.AppLogPath = conf.DefaultAppLogPath
		}
		if cfg.ErrorLogPath == "" {
			cfg.ErrorLogPath = conf.DefaultErrorLogPath
		}
	}

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewConsoleEncoder(encoderCfg)

	// 组合多个输出目标
	var (
		cores   []zapcore.Core
		closers []func()
	)

	// 控制台输出
	if hasStdout {
		cores = append(cores, zapcore.NewCore(
			encoder,
			zapcore.AddSync(os.Stdout),
			level,
		))
	}

	// 文件输出
	if hasFile {
		appWS, closeApp, err := zap.Open(cfg.AppLogPath)
		if err != nil {
			return nil, nil, err
		}
		errWS, closeErr, err := zap.Open(cfg.ErrorLogPath)
		if err != nil {
			closeApp()
			return nil, nil, err
		}
		closers = append(closers, closeApp, closeErr)

		// app.log: 所有级别(>= level)
		cores = append(cores, zapcore.NewCore(
			encoder,
			appWS,
			level,
		))

		// error.log: 仅错误级别(>= Error 且 >= level)
		errLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= level && lvl >= zapcore.ErrorLevel
		})
		cores = append(cores, zapcore.NewCore(
			encoder,
			errWS,
			errLevel,
		))
	}

	// 如果没有指定任何输出目标，默认输出到控制台
	if len(cores) == 0 {
		cores = append(cores, zapcore.NewCore(
			encoder,
			zapcore.AddSync(os.Stdout),
			level,
		))
	}

	logger := zap.New(zapcore.NewTee(cores...))
	sugar = logger.Sugar()

	cleanup := func() {
		for _, c := range closers {
			c()
		}
		_ = logger.Sync()
	}
	return &Logger{sugar: sugar}, cleanup, nil
}

// 直接记录消息内容, 例如: logx.Info("This is an info message")
// 输出: This is an info message
// 通过 sugar.Debug/Info/Warn/Error/Fatal 实现
func (log *Logger) Debug(args ...any) {
	if log != nil && log.sugar != nil {
		log.sugar.Debug(args...)
	}
}
func (log *Logger) Info(args ...any) {
	if log != nil && log.sugar != nil {
		log.sugar.Info(args...)
	}
}
func (log *Logger) Warn(args ...any) {
	if log != nil && log.sugar != nil {
		log.sugar.Warn(args...)
	}
}
func (log *Logger) Error(args ...any) {
	if log != nil && log.sugar != nil {
		log.sugar.Error(args...)
	}
}
func (log *Logger) Fatal(args ...any) {
	if log != nil && log.sugar != nil {
		log.sugar.Fatal(args...)
	}
}

// 带键值对的日志, 例如: logx.Infow("message", "key1", value1, "key2", value2)
// 输出: message key1=value1 key2=value2
func (log *Logger) Debugw(msg string, keysAndValues ...any) {
	if log != nil && log.sugar != nil {
		log.sugar.Debugw(msg, keysAndValues...)
	}
}
func (log *Logger) Infow(msg string, keysAndValues ...any) {
	if log != nil && log.sugar != nil {
		log.sugar.Infow(msg, keysAndValues...)
	}
}
func (log *Logger) Warnw(msg string, keysAndValues ...any) {
	if log != nil && log.sugar != nil {
		log.sugar.Warnw(msg, keysAndValues...)
	}
}
func (log *Logger) Errorw(msg string, keysAndValues ...any) {
	if log != nil && log.sugar != nil {
		log.sugar.Errorw(msg, keysAndValues...)
	}
}
func (log *Logger) Fatalw(msg string, keysAndValues ...any) {
	if log != nil && log.sugar != nil {
		log.sugar.Fatalw(msg, keysAndValues...)
	}
}
