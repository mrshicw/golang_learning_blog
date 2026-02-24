package test

import (
	"blog/config"
	"blog/models"
	"log"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// go test ./test/ -run TestInitDB -v
// 初始化数据库、创建数据库表
func TestInitDB(t *testing.T) {
	var err error
	DBFile := config.GetDBFile()

	DB, err := gorm.Open(sqlite.Open(DBFile), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	log.Println("创建SQLite数据库成功")

	// 自动迁移数据库表结构
	err = DB.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Comment{},
	)

	if err != nil {
		log.Fatal("自动迁移数据库失败：", err)
	}

	log.Println("自动迁移数据库成功")
}

// go test ./test/ -run TestCreateData -v
// 向数据库插入数据
func TestCreateData(t *testing.T) {
	config.OpenDB()
	db := config.GetDBConect()

	// user表
	user := models.User{
		Username: "user1",
		Email:    "user1@163.com",
		Password: "pwd1",
	}
	db.Where("username = ?", "user1").Delete(&models.User{})
	if err := db.Create(&user).Error; err != nil {
		log.Fatal("插入User数据失败", user, err)
	}

	// post表
	post := models.Post{
		Title:   "title1",
		Content: "content1",
	}
	db.Where("title = ?", "title1").Delete(&models.Post{})
	if err := db.Create(&post).Error; err != nil {
		log.Fatal("插入Post数据失败", post, err)
	}

	// comments 表
	//comment := models.Comment{
	//	Content: "评论1",
	//}
	//if err := db.Create(&comment).Error; err != nil {
	//	log.Fatal("插入Comment数据失败", comment, err)
	//}
}
