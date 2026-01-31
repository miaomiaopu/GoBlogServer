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

	if err := core.RunHTTP(cfg); err != nil {
		log.Fatalf("server exited: %v", err)
	}
}
