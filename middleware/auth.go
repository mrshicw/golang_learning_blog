package middleware

import (
	"blog/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 1. 从请求头中获取Authorization字段
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			utils.Unauthorized(ctx, "需要授权")
			ctx.Abort()
			return
		}

		// 2. 检查Bearer前缀
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			utils.Unauthorized(ctx, "无效的token格式")
			ctx.Abort()
			return
		}

		// 3. 解析token
		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			utils.Unauthorized(ctx, "无效的token")
			ctx.Abort()
			return
		}

		// 4. 将用户信息存储在上下文中，供后续处理函数使用
		ctx.Set("user_id", claims.UserID)
		ctx.Set("username", claims.Username)

		// 5. 继续处理请求
		ctx.Next()

	}

}
