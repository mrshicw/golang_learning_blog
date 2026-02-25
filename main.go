package main

import (
	"blog/config"
	"blog/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	// 加载配置文件
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	host := cfg.Server.Host
	port := cfg.Server.Port
	// 设置 Gin 模式
	gin.SetMode(cfg.Server.Mode)
	// fmt.Println("mode", cfg.Server.Mode)

	// 初始化日志
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Info("启动博客应用")

	// 打开数据库连接
	config.OpenDB()

	// 设置路由
	r := routes.SetupRoutes()

	// 启动服务器
	r.Run(host + ":" + port)

	// 起初测试创建Gin引擎
	// r := gin.Default()

	// 健康检查
	//r.GET("/health", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "Healthy",
	//	})
	//})

	//r.Run()
}
