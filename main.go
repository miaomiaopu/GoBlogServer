package main

import (
	"flag"
	"fmt"
	"log"

	"Server/conf"

	"github.com/gin-gonic/gin"
)

func main() {
	// Define command-line flag for config file path, -c
	// 读取配置文件路径的命令行参数, -c
	cfgPath := flag.String("c", "conf/config.yaml", "path to config yaml")
	flag.Parse()

	// Load configuration from the specified file
	// 从指定文件加载配置
	cfg, err := conf.LoadConfig(*cfgPath)
	if err != nil {
		log.Fatalf("load config failed: %v", err)
	}

	// Initialize Gin router
	// 初始化Gin路由
	r := gin.Default()

	// Define a simple ping route to test the server
	// 定义一个简单的ping路由来测试服务器
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"port":    cfg.Server.Port,
		})
	})

	// Start the server on the configured port
	// 在配置的端口上启动服务器
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	fmt.Printf("Starting server on %s (config=%s)\n", addr, *cfgPath)
	if err := r.Run(addr); err != nil {
		log.Fatalf("server exited: %v", err)
	}
}
