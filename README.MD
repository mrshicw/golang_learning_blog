# Blog API
用golang语言，基于gin框架、gorm库的个人博客后端API

# 基本功能
- 基于JWt认证授权的注册、登录
- 文章CRUD操作
- 评论发表、查看功能
- 统一日志记录
- 统一错误处理
- 数据库自动迁移

# 1、运行环境
- linux ubuntu 24.04
- go 1.25.6
- gin Web框架
- gorm orm库
- sqlite 数据库
- jwt 身份认证
- logrus 日志库
- bcrypt 密码加密

# 2、依赖安装步骤
```bash
go mod init "blog"
go mod tidy
```

```bash
go test ./test/ -run TestInitDB -v
```

# 3、启动方式
```bash
go run main.go
```

 # 目录结构
 ```
├── main.go                 # 程序入口
├── config
│   ├── config.go           # 读取数据库配置
│   ├── config.yaml         # 配置文件
├── controllers
│   ├── auth.go             # 注册、登录等控制器
│   ├── comment.go          # 评论控制器
│   └── post.go             # 文章控制器
├── db
│   └── sqlite.db           # 生成的sqlite数据库
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
- ID            # 主键，唯一
- USERNAME      # 用户名，唯一
- EMAIL         # 邮箱，唯一
- PASSWORD      # 密码，加密
- CREATED_AT
- UPDATED_AT
- DELETED_AT

## Posts表
- ID            # 主键，唯一
- TITLE         # 标题
- CONTENT       # 内容
- USER_ID       # 外键，Users表
- CREATED_AT
- UPDATED_AT
- DELETED_AT

## Comments表
- ID            # 主键，唯一
- CONTENT       # 内容
- USER_ID       # 外键，Users表
- POST_ID       # 外键，Posts表
- CREATED_AT
- UPDATED_AT
- DELETED_AT

# 使用示例
## 用户注册
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
     -H "Content-Type: application/json" \
     -d '{"Username": "bloguser", "Email": "bloguser@example.com", "Password": "bloguser123666"}'
```
```
{ "code":    200,
  "message": "success",
  "data":    { "token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6ImJsb2d1c2VyIiwiZXhwIjoxNzcyMDY2MDY2LCJuYmYiOjE3NzE5Nzk2NjYsImlhdCI6MTc3MTk3OTY2Nn0.VM3hvk81zOJ-1kN0GmDlkxpLydE6ubrDQvdqxMy60e8",
               "user":  { "ID":2,
                            "CreatedAt":"2026-02-25T08:34:26.320750911+08:00",
                            "UpdatedAt":"2026-02-25T08:34:26.320750911+08:00",
                            "DeletedAt":null,
                            "Username":"bloguser",
                            "Password":"$2a$10$xR030jQm1Jr24EOBSiXBl.VCekPJFqSb83YVCyvwT5BThdzRdmMua",
                            "Email":"bloguser@example.com"
                        }
            }
}
```

## 用户登录
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
     -H "Content-Type: application/json" \
     -d '{"Username": "bloguser", "Password": "bloguser123666"}'
```
```
{ "code":    200,
  "message": "success",
  "data":    { "token":  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6ImJsb2d1c2VyIiwiZXhwIjoxNzcyMDY2NDYxLCJuYmYiOjE3NzE5ODAwNjEsImlhdCI6MTc3MTk4MDA2MX0.pgkdb_m8yZ7iJg39Jfkps5OIdYN31H88NCBPi7AjQpI",
               "user":   { "ID":2,
                          "CreatedAt":"2026-02-25T08:34:26.320750911+08:00",
                          "UpdatedAt":"2026-02-25T08:34:26.320750911+08:00",
                          "DeletedAt":null,
                          "Username":"bloguser",
                          "Password":"$2a$10$xR030jQm1Jr24EOBSiXBl.VCekPJFqSb83YVCyvwT5BThdzRdmMua",
                          "Email":"bloguser@example.com"
                         }
            }
}
```

## 获取用户信息
```bash
curl -X POST http://localhost:8080/api/v1/authed/profile \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

``` 
{ "code":     200,
  "message":  "success",
  "data":   { "ID":2,
                "CreatedAt":"2026-02-25T08:34:26.320750911+08:00",
                "UpdatedAt":"2026-02-25T08:34:26.320750911+08:00",
                "DeletedAt":null,
                "Username":"bloguser",
                "Password":"$2a$10$xR030jQm1Jr24EOBSiXBl.VCekPJFqSb83YVCyvwT5BThdzRdmMua",
                "Email":"bloguser@example.com"
            }
}
```

## 创建文章
```bash
curl -X POST http://localhost:8080/api/v1/authed/posts \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer YOUR_JWT_TOKEN" \
     -d '{"Title":   "Title 1","Content": "Content 1",}'
```

``` 
{ "code":     200,
  "message":  "success",
  "data":     { "ID":3,
                 "CreatedAt":"2026-02-25T09:03:08.732893815+08:00",
                 "UpdatedAt":"2026-02-25T09:03:08.732893815+08:00",
                 "DeletedAt":null,
                 "Title":"Title 1",
                 "Content":"Content 1",
                 "UserID":2,
                 "User": { "ID":2,
                           "CreatedAt":"2026-02-25T08:34:26.320750911+08:00",
                           "UpdatedAt":"2026-02-25T08:34:26.320750911+08:00",
                           "DeletedAt":null,
                           "Username":"bloguser",
                           "Password":"$2a$10$xR030jQm1Jr24EOBSiXBl.VCekPJFqSb83YVCyvwT5BThdzRdmMua",
                           "Email":"bloguser@example.com"
                        }
                }
}
```

## 更新文章
```bash
curl -X PUT http://localhost:8080/api/v1/authed/posts/1 \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer YOUR_JWT_TOKEN" \
     -d '{"Title": "Title Update","Content": "Content Update",}'
```

```
{ "code":     200,
  "message":  "success",
  "data":    { "ID":1,
               "CreatedAt":"2026-02-24T19:46:19.651819119+08:00",
               "UpdatedAt":"2026-02-25T10:12:10.477844152+08:00",
               "DeletedAt":null,
               "Title":"Title Update",
               "Content":"Content Update",
               "UserID":1,
               "User":  { "ID":1,
                          "CreatedAt":"2026-02-24T19:42:49.63947599+08:00",
                          "UpdatedAt":"2026-02-24T19:42:49.63947599+08:00",
                          "DeletedAt":null,
                          "Username":"shicw",
                          "Password":"$2a$10$TEYiTpqlwi6hBZFxfdj/1uZyZZ8w/1dqZ0Xb1sCChSYrQQ4WL6fui",
                          "Email":"shicw@example.com"
                        }
            }
}
```

## 删除文章
```bash
curl -X DELETE http://localhost:8080/api/v1/authed/posts/2 \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

``` 
{ "code":    200,
  "message": "success",
  "data":    {}
}
```
## 对文章ID发表评论
```bash
/api/v1/posts/:post_id/comments
curl -X POST http://localhost:8080/api/v1/authed/posts/1/comments \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer YOUR_JWT_TOKEN" \
     -d '{"Content": "This is a comment" }'

```

``` 
{ "code":    200,
  "message": "success",
  "data":    { "ID":2,
               "CreatedAt":"2026-02-25T10:14:48.284135625+08:00",
               "UpdatedAt":"2026-02-25T10:14:48.284135625+08:00",
               "DeletedAt":null,
               "Content":  "This is a comment",
               "UserID":   1,
               "User": { "ID":1,
                         "CreatedAt":"2026-02-24T19:42:49.63947599+08:00",
                         "UpdatedAt":"2026-02-24T19:42:49.63947599+08:00",
                         "DeletedAt":null,
                         "Username":"shicw",
                         "Password":"$2a$10$TEYiTpqlwi6hBZFxfdj/1uZyZZ8w/1dqZ0Xb1sCChSYrQQ4WL6fui",
                         "Email":"shicw@example.com"
                         },
                "PostID": 1,
                "Post":  { "ID":0,
                           "CreatedAt":"0001-01-01T00:00:00Z",
                           "UpdatedAt":"0001-01-01T00:00:00Z",
                           "DeletedAt":null,
                           "Title":    "",
                           "Content":  "",
                           "UserID":   0,
                           "User":   { "ID":0,
                                       "CreatedAt":"0001-01-01T00:00:00Z",
                                       "UpdatedAt":"0001-01-01T00:00:00Z",
                                       "DeletedAt":null,
                                       "Username":"",
                                       "Password":"",
                                       "Email":""
                                    }
                        }
            }
}
```

## 获取全部评论
```bash
curl -X GET http://localhost:8080/api/v1/public/posts
```

``` 
{ "code":    200,
  "message": "success",
  "data":   { "page":1,
              "page_size":10,
              "posts": [{ "ID":1,
                          "CreatedAt":"2026-02-24T19:46:19.651819119+08:00",
                          "UpdatedAt":"2026-02-24T19:46:51.919602915+08:00",
                          "DeletedAt":null,
                          "Title":"My Update Post",
                          "Content":"This is the content of my update post.",
                          "UserID":1,
                          "User":  { "ID":1,
                                     "CreatedAt":"2026-02-24T19:42:49.63947599+08:00",
                                     "UpdatedAt":"2026-02-24T19:42:49.63947599+08:00",
                                     "DeletedAt":null,
                                     "Username":"shicw",
                                     "Password":"$2a$10$TEYiTpqlwi6hBZFxfdj/1uZyZZ8w/1dqZ0Xb1sCChSYrQQ4WL6fui",
                                     "Email":"shicw@example.com"
                                     }
                        },
                        { "ID":3,
                          "CreatedAt":"2026-02-25T09:03:08.732893815+08:00",
                          "UpdatedAt":"2026-02-25T09:03:08.732893815+08:00",
                          "DeletedAt":null,
                          "Title":"Title 1",
                          "Content":"Content 1",
                          "UserID":2,
                          "User": { "ID":2,
                                    "CreatedAt":"2026-02-25T08:34:26.320750911+08:00",
                                    "UpdatedAt":"2026-02-25T08:34:26.320750911+08:00",
                                    "DeletedAt":null,
                                    "Username":"bloguser",
                                    "Password":"$2a$10$xR030jQm1Jr24EOBSiXBl.VCekPJFqSb83YVCyvwT5BThdzRdmMua",
                                    "Email":"bloguser@example.com"
                                  }
                        }],
              "total":2
            }
}
```
## 获取指定ID的评论
```bash
curl -X GET http://localhost:8080/api/v1/public/posts/1
```

``` 
{ "code":   200,
  "message":"success",
  "data":   { "ID":1,
              "CreatedAt":"2026-02-24T19:46:19.651819119+08:00",
              "UpdatedAt":"2026-02-24T19:46:51.919602915+08:00",
              "DeletedAt":null,
              "Title":"My Update Post",
              "Content":"This is the content of my update post.",
              "UserID":1,
              "User": { "ID":1,
                        "CreatedAt":"2026-02-24T19:42:49.63947599+08:00",
                        "UpdatedAt":"2026-02-24T19:42:49.63947599+08:00",
                        "DeletedAt":null,
                        "Username":"shicw",
                        "Password":"$2a$10$TEYiTpqlwi6hBZFxfdj/1uZyZZ8w/1dqZ0Xb1sCChSYrQQ4WL6fui",
                        "Email":"shicw@example.com"
                      }
            }
}
```
## 健康检查
```bash
curl -X GET http://localhost:8080/health
```

``` 
{ "message": "2026-02-25T09:07:51+08:00",
  "status":  "OK"
}
```