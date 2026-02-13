package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 日志记录中间件
func LoggerMiddleWare() gin.HandlerFunc {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		logger.WithFields(logrus.Fields{
			"status_code": param.StatusCode,
			"latency":     param.Latency,
			"client_it":   param.ClientIP,
			"method":      param.Method,
			"path":        param.Path,
			"user_agent":  param.Request.UserAgent(),
			"error":       param.ErrorMessage,
			"timestamp":   param.TimeStamp.Format(param.TimeStamp.Format(time.RFC3339)),
		}).Info("HTTP Request")
		return ""
	})
}

// 全局错误处理中间件
func ErrorHandleMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logrus.WithFields(logrus.Fields{
					"error":  err,
					"path":   ctx.Request.URL.Path,
					"mathod": ctx.Request.Method,
				}).Error("恢复异常")
			}

			ctx.JSON(500, gin.H{
				"code":    500,
				"message": "服务器内部错误",
			})
			ctx.Abort()
		}()
		ctx.Next()
	}
}
