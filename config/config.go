package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// 获取SQLite存储目录及文件
func GetDBFile() string {
	filename := "sqlite.db"

	_, currentFile, _, _ := runtime.Caller(0)

	// 获取当前运行文件所在目录
	pwdDir := filepath.Dir(currentFile)

	// 获得上一级目录，即项目根目录
	projectDir := filepath.Dir(pwdDir)

	// 数据库目录为项目根目录下的 "db" 文件夹
	dbDir := filepath.Join(projectDir, "db")

	// 创建目录（如果不存在）
	os.MkdirAll(dbDir, 0755)

	dbPath := filepath.Join(dbDir, filename)

	return dbPath
}

// 获取环境变量
// 如果不存在，返回默认值
func myGetEvn(key, val string) string {
	// 标准库函数，os.Getevn(key)
	if val := os.Getenv(key); val != "" {
		return val
	}
	return val
}

func OpenDB() {
	DBFile := GetDBFile()
	db, err := gorm.Open(sqlite.Open(DBFile))
	if err != nil {
		log.Fatal("打开数据库失败！")
	}
	DB = db
}

func GetDBConect() *gorm.DB {
	return DB
}
