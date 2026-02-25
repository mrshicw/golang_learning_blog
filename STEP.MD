# 系统搭建过程
github 的 .gitignore
*.db
*.gitignore

# 启动main服务程序
pwd
blog/
vim main.go
main.go => routes/routes.go 配置的路由

go mod init "blog"
go mod tidy
go run main.go

# 添加基本的models
pwd
models/
user.go
post.go
comment.go

# 读取配置文件
pwd
config/
config.go <= database_test.go
vim config.go


# 数据库迁移程序，系统启动前运行
pwd
test/
database_test.go => config.go
vim database_test.go

迁移数据库
go test ./test/ -run TestInitDB -v

插入测试数据
go test ./test/ -run TestCreateData -v


# 设置路由
pwd
routes/
routes.go

routes.go  <=  main.go
routes.go  =>  utils/response.go
routes.go  =>  controller/auth.go
routes.go  =>  controller/post.go
routes.go  =>  controller/comment.go

/api/v1/auth/register
/api/v1/auth/login

/api/v1/profile
/api/v1/posts
/api/v1/posts/:id
/api/v1/posts/:post_id/comments
/api/v1/comments
/api/v1/comments/post/:post_id


# 完成response
pwd
utils/
vim response.go


# 编写controllers
pwd
auth.go     <= routes.go
post.go     <= routes.go
comment.go  <= routes.go
controllers/
vim auth.go
vim post.go
vim comment.go

# 编写中间件middleware
middleware/auth.go <= routes.go
