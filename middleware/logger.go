package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LoggerMiddleware() gin.HandlerFunc {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	return gin.LoggerWithFormatter(
		func(param gin.LogFormatterParams) string {
			// 记录请求日志
			logger.WithFields(logrus.Fields{
				"client_ip":   param.ClientIP,
				"time_stamp":  param.TimeStamp.Format("2006-01-02 15:04:05"),
				"method":      param.Method,
				"path":        param.Path,
				"status_code": param.StatusCode,
				"latency":     param.Latency,
				"user_agent":  param.Request.UserAgent(),
			}).Info("请求日志")
			return ""
		})
}
