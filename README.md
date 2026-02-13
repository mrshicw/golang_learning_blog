# golang_learning_blog

cd golang_learning_blog

go mod init "golang_learning_blog"

go mod tidy


# 配置文件
vim .env


# 定义models
mkdir -p models && cd models/
vim user.go
package models
....
vim comment.go
vim post.go


# 配置文件
mkdir -p config && cd config/
vim config.go
package config
...

# 初始化数据库、迁移数据库
mkdir -p test && cd test/
vim database_test.go
package test
func TestInitDB(t *testing.T) {}
# go test ./test/ -run TestInitDB -v
// 插入数据
go test ./test/ -run TestCreateData -v

# 定义Response
mkdir -p utils && cd utils
vim response.go
package uitls
...

# 定义路由
mkdir -p routes && cd routes
vim routes.go
package routes
...

go clean -modcache -cache -i -r
# 启动服务器
mkdir -p setup && cd setup
vim main.go
package main
...

# 添加控制器
mkdir -p controllers && cd controllers
vim auth.go
package controllers
... struct, Login ...

# 添加中间件
mkdir -p middleware && cd middleware
vim auth.go
package middleware