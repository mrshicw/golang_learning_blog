package routes

import (
	"golang_learning_blog/controllers"
	"golang_learning_blog/utils"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.New()

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
			// /api/v1/auth/register
			auth.POST("/register", authController.Register)
			// /api/v1/auth/login
			auth.POST("/login", authController.Login)
		}

		// 2、认证路由：用户信息、文章、评论
		authed := api.Group("")
		// +认证

		{
			// 用户信息
			// /api/v1/profile
			authed.GET("/profile", authController.GetProfile)

			// 文章
			posts := authed.Group("/posts")
			{
				// /api/v1/posts
				posts.POST("", postController.CreatePost)
				// /api/v1/posts/:id
				posts.PUT("/:id", postController.UpdatePost)
				// /api/v1/posts/:id
				posts.DELETE("/:id", postController.DeletePost)

			}

			// 评论
			comments := authed.Group("/posts/:post_id/comments")
			{
				comments.POST("", commentController.CreateComment)
			}
		}

		// 3、不需要认证路由：文章
		public := api.Group("")
		{
			// /api/v1/posts
			public.GET("/posts", postController.GetPosts)
			// /api/v1/posts/:id
			public.GET("/posts/:id", postController.GetPost)
		}

		// 4、不需要认证路由：公开评论
		// /api/v1/comments
		comments := api.Group("/comments")
		{
			// /api/v1/comments/post/:post_id
			comments.GET("/post/:post_id", commentController.GetComments)
		}
	}

	// 健康检查
	r.GET("/health", utils.Health)

	return r
}
