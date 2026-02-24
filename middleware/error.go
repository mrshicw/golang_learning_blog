package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logrus.WithFields(logrus.Fields{
					"error":  err,
					"path":   c.Request.URL.Path,
					"method": c.Request.Method,
				}).Error("捕获到 panic 错误")

				c.JSON(500, gin.H{
					"code":    500,
					"message": "内部服务器错误",
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
