package controllers

import (
	"blog/config"
	"blog/models"
	"blog/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
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

	db := config.GetDBConect()
	// 用户名是否存在
	var existingUser models.User
	if err := db.Where("username = ?", req.Username).First(&existingUser).Error; err == nil {
		utils.BadRequest(c, "用户名已存在")
		return
	}

	// 邮箱是否存在
	if err := db.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		utils.BadRequest(c, "邮箱已被注册")
		return
	}

	// 创建用户
	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
	if err := db.Create(&user).Error; err != nil {
		utils.InternalServerError(c, "注册失败")
		return
	}

	// 生成JWT token
	token, err := utils.GenToken(user.ID, user.Username)
	if err != nil {
		utils.InternalServerError(c, "生成token失败")
		return
	}

	utils.Success(c, AuthResponse{
		Token: token,
		User:  user,
	})
}

// 用户登录
func (ac *AuthController) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	db := config.GetDBConect()
	// 查找用户
	var user models.User
	if err := db.Where("username = ?", req.Username).First(&user).Error; err != nil {
		utils.Unauthorized(c, "用户名或密码错误")
		return
	}

	// 验证密码
	if !user.CheckPsw(req.Password) {
		utils.Unauthorized(c, "用户名或密码错误")
		return
	}

	// 生成 JWT token
	token, err := utils.GenToken(user.ID, user.Username)
	if err != nil {
		utils.InternalServerError(c, "生成token失败")
		return
	}

	utils.Success(c, AuthResponse{
		Token: token,
		User:  user,
	})
}

// 用户信息
func (ac *AuthController) GetProfile(c *gin.Context) {
	// 中间件中设置了"user_id"键，获取用户ID
	userID, exists := c.Get("user_id")

	if !exists {
		utils.Unauthorized(c, "未授权")
		return
	}

	var user models.User
	db := config.GetDBConect()
	if err := db.First(&user, userID).Error; err != nil {
		utils.NotFound(c, "用户不存在")
		return
	}

	utils.Success(c, user)
}
