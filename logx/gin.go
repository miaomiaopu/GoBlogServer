package logx

import (
	"net/http"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
	"Server/global"
)

// GinLogger 使用 logx 的 zap 记录 gin 访问日志。
// 记录 HTTP 请求的状态码、方法、路径、客户端 IP、延迟时间、用户代理和错误信息。
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		if raw != "" {
			path = path + "?" + raw
		}

		status := c.Writer.Status()
		latency := time.Since(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		ua := c.Request.UserAgent()
		errMsg := c.Errors.ByType(gin.ErrorTypePrivate).String()

		if global.GlobalLogger == nil {
			return
		}

		// 5xx 用 Error，其它用 Info
		if status >= http.StatusInternalServerError {
			global.GlobalLogger.Errorw("http_request",
				"status", status,
				"method", method,
				"path", path,
				"ip", clientIP,
				"latency", latency.String(),
				"ua", ua,
				"error", errMsg,
			)
			return
		}

		global.GlobalLogger.Infow("http_request",
			"status", status,
			"method", method,
			"path", path,
			"ip", clientIP,
			"latency", latency.String(),
			"ua", ua,
			"error", errMsg,
		)
	}
}

// GinRecovery 使用 logx 记录 panic 恢复日志。
func GinRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if rec := recover(); rec != nil {
				if global.GlobalLogger != nil {
					global.GlobalLogger.Errorw("panic_recovered",
						"error", rec,
						"method", c.Request.Method,
						"path", c.Request.URL.Path,
						"ip", c.ClientIP(),
						"stack", string(debug.Stack()),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
