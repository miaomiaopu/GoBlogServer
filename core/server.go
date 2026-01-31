package core

import (
	"Server/conf"
	"Server/logx"
	"fmt"

	"github.com/gin-gonic/gin"
)

// RunHTTP 启动 HTTP 服务
func RunHTTP(cfg *conf.Config) error {
	gin.SetMode(cfg.Server.GinMode)
	r := gin.New()
	// 日志中间件
	r.Use(logx.GinLogger(), logx.GinRecovery())

	// 健康检查接口
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"port":    cfg.Server.Port,
		})
	})

	// 启动服务器
	addr := fmt.Sprintf(":%d", cfg.Server.Port)

	return r.Run(addr)
}
