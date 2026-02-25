package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
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

type Config struct {
	Server ServerConfig `mapstructure:"server"`
	JWT    JWTConfig    `mapstructure:"jwt"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
	Host string `mapstructure:"host"`
	Mode string `mapstructure:"mode"`
}

type JWTConfig struct {
	Secret string `mapstructure:"secret"`
	Expire uint   `mapstructure:"expire"`
}

var CONFIG *Config

func init() {
	// 设置配置文件名称（不含扩展名）
	viper.SetConfigName("config")
	// 设置配置文件类型
	viper.SetConfigType("yaml")
	// 添加配置文件搜索路径
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.app")

	// 读取环境变量
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APP")

	// 设置默认值
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("server.host", "0.0.0.0")
	viper.SetDefault("server.mode", "debug")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: Error reading config file: %v", err)
		log.Println("Using default values and environment variables")
	} else {
		log.Printf("Using config file: %s", viper.ConfigFileUsed())
	}
}

// LoadConfig 加载并解析配置文件，将 Viper 中的配置数据映射到 Config 结构体
// 返回解析后的配置对象指针和可能的错误
// 配置数据来源包括：配置文件、环境变量和默认值（在 init 函数中已设置）
func LoadConfig() (*Config, error) {
	var config Config

	// 使用 viper.Unmarshal 将配置数据解析到 config 结构体中
	// 如果解析失败，返回错误
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	// 返回解析成功的配置对象
	return &config, nil
}

func GetConfig() *Config {
	CONFIG, _ := LoadConfig()
	return CONFIG
}
