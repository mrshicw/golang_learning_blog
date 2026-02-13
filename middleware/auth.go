package middleware

import "github.com/gin-gonic/gin"

// JWT认证中间件
func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//ctx.Set("user_id", claims.UserID)
		//ctx.Set("username", claims.Username)
		ctx.Next()
	}
}
