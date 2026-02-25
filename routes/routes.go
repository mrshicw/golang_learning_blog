package routes

import (
	// "golang_learning_blog/controllers"
	// "golang_learning_blog/utils"

	"blog/controllers"
	"blog/middleware"
	"time"

	"github.com/gin-gonic/gin"
)

// health
// 临时的健康检查处理函数，实际项目中应该替换为真正的处理函数
func Health(c *gin.Context) {
	c.JSON(200,
		gin.H{
			"status":  "OK",
			"message": time.Now().Format(time.RFC3339),
		})
}

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	// 1. 全局中间件
	r.Use(middleware.LoggerMiddleware())
	r.Use(middleware.ErrorHandlingMiddleware())
	r.Use(gin.Recovery())

	// 创建控制器实例
	authController := &controllers.AuthController{}
	postController := &controllers.PostController{}
	commentController := &controllers.CommentController{}

	// API路由组
	api := r.Group("/api/v1")
	{
		// 1、不需要认证路由：注册、登录认证
		auth := api.Group("/auth")
		{
			// 注册
			// /api/v1/auth/register
			auth.POST("/register", authController.Register)
			// 登录
			// /api/v1/auth/login
			auth.POST("/login", authController.Login)
		}

		// 2、认证路由：用户信息、文章、评论
		authed := api.Group("/authed")
		// +认证
		authed.Use(middleware.AuthMiddleWare())
		{
			// 获取用户信息
			// /api/v1/authed/profile
			authed.POST("/profile", authController.GetProfile)

			//  文章
			posts := authed.Group("/posts")
			{
				// 发表文章
				// /api/v1/authed/posts
				posts.POST("", postController.CreatePost)
				// 修改文章
				// /api/v1/authed/posts/:id
				posts.PUT("/:id", postController.UpdatePost)
				// 删除文章
				// /api/v1/authed/posts/:id
				posts.DELETE("/:id", postController.DeletePost)

			}

			// 评论
			comments := authed.Group("")
			{
				// 对文章ID发表评论
				// /api/v1/authed/posts/:post_id/comments
				comments.POST("/posts/:post_id/comments", commentController.CreateComment)
			}
		}

		// 3、不需要认证路由：文章
		public := api.Group("/public")
		{
			// 查看所有评论列表
			// /api/v1/public/posts
			public.GET("/posts", postController.GetPosts)
			// 查看文章ID评论
			// /api/v1/public/posts/:id
			public.GET("/posts/:id", postController.GetPost)
		}

		// 4、不需要认证路由：公开评论
		comments := api.Group("/comments")
		{
			// 查看评论
			// /api/v1/comments/post/:post_id
			comments.GET("/post/:post_id", commentController.GetComments)
		}
	}

	// 健康检查
	r.GET("/health", Health)
	r.POST("/health", Health)

	return r
}
