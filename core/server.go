package core

import (
	"Server/conf"
	"fmt"

	"github.com/gin-gonic/gin"
)

// RunHTTP 启动 HTTP 服务
func RunHTTP(cfg *conf.Config) error {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"port":    cfg.Server.Port,
		})
	})

	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	fmt.Printf("Starting server on %s\n", addr)

	return r.Run(addr)
}
