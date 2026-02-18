package controllers

import (
	"golang_learning_blog/models"
	"golang_learning_blog/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min6"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

// 用户注册
func (ac *AuthController) Register(c *gin.Context) {
	// 绑定json
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	// 用户名是否存在

	// 邮箱是否存在

	// 创建用户
	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
	// 存储到数据库

	// 生成JWT token
	token := ""

	utils.Success(c, AuthResponse{
		Token: token,
		User:  user,
	})
}

// 用户登录
func (ac *AuthController) Login(c *gin.Context) {
	// TODO
	utils.Success(c, AuthResponse{
		Token: "",
		User:  models.User{},
	})
}

// 用户信息
func (ac *AuthController) GetProfile(c *gin.Context) {
	// TODO
	// 这里需要实现从JWT token中解析用户信息的逻辑，暂时返回一个空的用户对象
	utils.Success(c, models.User{})
}
