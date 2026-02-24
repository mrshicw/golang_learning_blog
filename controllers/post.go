package controllers

import (
	"blog/config"
	"blog/models"
	"blog/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostController struct{}

type CreatePostRequest struct {
	Title   string `json:"title" binding:"required,min=1,max=200"`
	Content string `json:"content" binding:"required,min=1"`
}

type UpdatePostRequest struct {
	Title   string `json:"title" binding:"required,min=1,max=200"`
	Content string `json:"content" binding:"required,min=1"`
}

// 创建文章
func (pc *PostController) CreatePost(c *gin.Context) {
	// 绑定json
	var req CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}
	// 获取用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		utils.Unauthorized(c, "用户未认证")
		return
	}
	// 创建文章
	post := models.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  userID.(uint),
	}
	db := config.GetDBConect()
	if err := db.Create(&post).Error; err != nil {
		utils.InternalServerError(c, "创建文章失败")
		return
	}
	// 预加载用户信息
	db.Preload("User").First(&post, post.ID)

	utils.Success(c, post)
}

// 获取文章列表
func (pc *PostController) GetPosts(c *gin.Context) {
	//
	var posts []models.Post

	// 获取分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 50 {
		pageSize = 50
	}
	offset := (page - 1) * pageSize

	db := config.GetDBConect()
	if err := db.Preload("User").
		Limit(pageSize).
		Offset(offset).
		Find(&posts).Error; err != nil {
		utils.InternalServerError(c, "获取文章列表失败")
		return
	}
	var total int64
	db.Model(&models.Post{}).Count(&total)

	utils.Success(c, gin.H{
		"posts":     posts,
		"page":      page,
		"page_size": pageSize,
		"total":     total,
	})
}

// 获取单个文章详情
func (pc *PostController) GetPost(c *gin.Context) {
	//
	postID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的文章ID")
		return
	}

	var post models.Post
	db := config.GetDBConect()
	if err := db.Preload("User").First(&post, postID).Error; err != nil {
		utils.NotFound(c, "文章不存在")
		return
	}

	utils.Success(c, post)
}

// 更新文章
func (pc *PostController) UpdatePost(c *gin.Context) {
	// 解析文章ID
	postID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的文章ID")
		return
	}
	// 绑定json
	var req UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}
	// 获取用户ID
	userID, existes := c.Get("user_id")
	if !existes {
		utils.Unauthorized(c, "用户未认证")
		return
	}
	// 验证文章是否存在以及是否属于当前用户
	var post models.Post
	db := config.GetDBConect()
	if err := db.First(&post, postID).Error; err != nil {
		utils.NotFound(c, "文章不存在")
		return
	}
	// 验证文章是否属于当前用户
	if post.UserID != userID.(uint) {
		utils.Forbidden(c, "没有权限修改该文章")
		return
	}

	// 更新文章
	post.Title = req.Title
	post.Content = req.Content
	if err := db.Save(&post).Error; err != nil {
		utils.InternalServerError(c, "更新文章失败")
		return
	}
	// 预加载用户信息
	db.Preload("User").First(&post, post.ID)

	utils.Success(c, post)
}

// 删除文章
func (pc *PostController) DeletePost(c *gin.Context) {
	//
	postID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的文章ID")
		return
	}

	// 获取用户ID
	userID, existes := c.Get("user_id")
	if !existes {
		utils.Unauthorized(c, "用户未认证")
		return
	}

	// 验证文章是否存在以及是否属于当前用户
	var post models.Post
	db := config.GetDBConect()
	if err := db.First(&post, postID).Error; err != nil {
		utils.NotFound(c, "文章不存在")
		return
	}

	// 验证文章是否属于当前用户
	if post.UserID != userID.(uint) {
		utils.Forbidden(c, "没有权限删除该文章")
		return
	}

	// 删除文章
	if err := db.Delete(&post).Error; err != nil {
		utils.InternalServerError(c, "删除文章失败")
		return
	}

	utils.Success(c, gin.H{})
}
