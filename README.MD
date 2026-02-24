# 1、运行环境
linux ubuntu 24.04
go 1.25.6
数据库：sqlite

# 2、依赖安装步骤
go mod init "blog"
go mod tidy

## 初始化数据
go test ./test/ -run TestInitDB -v

# 3、启动方式
go run main.go





 # 目录结构
 ```
├── config
│   └── config.go           # 读取数据库配置
├── controllers
│   ├── auth.go             # 注册、登录等控制器
│   ├── comment.go          # 评论控制器
│   └── post.go             # 文章控制器
├── db
│   └── sqlite.db           # 生成的sqlite数据库
├── main.go                 # 程序入口
├── middleware
│   ├── auth.go             # 权限中间件
│   ├── error.go            # 错误处理中间件
│   └── logger.go           # 日志中间件
├── models
│   ├── comment.go          # comment模型
│   ├── post.go             # post模型
│   └── user.go             # user模型
├── README.MD               # README文件
├── routes
│   └── routes.go           # 路由配置
├── STEP.MD
├── test
│   ├── api_test.go         # api测试
│   ├── database_test.go    # 数据库初始化测试
│   └── test.py             # Python测试
└── utils
    ├── jwt.go              # 权限控制
    └── response.go         # 请求响应
 ```

# 数据库设计
## Users表
ID
USERNAME
EMAIL
PASSWORD
CREATED_AT
UPDATED_AT
DELETED_AT

## Posts表
ID
TITLE
CONTENT
USER_ID
CREATED_AT
UPDATED_AT
DELETED_AT

## Comments表
ID
CONTENT
USER_ID
POST_ID
CREATED_AT
UPDATED_AT
DELETED_AT


# API接口
## 认证接口
POST  /api/v1/auth/register    # 用户注册


## 文章接口
GET   /api/v1/posts            # 获取文章列表

## 评论接口
GET   /api/v1/comments/posts/:post_id    # 获取文章评论

# 其他接口
GET   /health                            # 健康检查


# 运行项目



# 使用示例
## 用户注册
``` ```

## 用户登录
``` ```

## 创建文章
``` ```
