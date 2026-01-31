package main

import (
	"Server/core"
	appflag "Server/flag"
	"Server/global"
	"log"
)

func main() {
	opt := appflag.Parse()

	cfg, err := core.InitConfig(opt.ConfigPath)
	if err != nil {
		log.Fatalf("init config failed: %v", err)
	}
	global.GlobalConfig = cfg

	logger, cleanup, err := core.InitLogger(global.GlobalConfig.Log)
	if err != nil {
		log.Fatalf("init logger failed: %v", err)
	}
	global.GlobalLogger = logger

	// 确保日志缓冲区的内容在程序退出前被写入
	defer cleanup()

	if err := core.RunHTTP(cfg); err != nil {
		global.GlobalLogger.Fatalw("failed to run HTTP server", "error", err)
	}

}
