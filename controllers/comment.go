package controllers

import (
	"blog/config"
	"blog/models"
	"blog/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentController struct{}

type CreateCommentRequest struct {
	Content string `json:"content" binding:"required,min=1,max=1000"`
}

// 创建评论
func (cc *CommentController) CreateComment(c *gin.Context) {
	// 获取文章ID
	postID, err := strconv.ParseUint(c.Param("post_id"), 10, 32)
	if err != nil {
		utils.BadRequest(c, "不合理的文章ID")
		return
	}

	// 获取用户ID
	var req CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "无效的请求数据: "+err.Error())
		return
	}

	// 获取用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		utils.Unauthorized(c, "用户未认证")
		return
	}

	// 检查文章是否存在
	db := config.GetDBConect()
	post := models.Post{}
	if err := db.First(&post, postID).Error; err != nil {
		utils.BadRequest(c, "文章不存在")
		return
	}

	// 创建评论
	comment := models.Comment{
		PostID:  uint(postID),
		UserID:  userID.(uint),
		Content: req.Content,
	}

	if err := db.Create(&comment).Error; err != nil {
		utils.InternalServerError(c, "创建评论失败")
		return
	}

	// 返回评论信息
	db.Preload("User").First(&comment, comment.ID)

	utils.Success(c, comment)
}

// 获得文章的评论列表
func (cc *CommentController) GetComments(c *gin.Context) {
	// 获取文章ID
	postID, err := strconv.ParseUint(c.Param("post_id"), 10, 32)
	if err != nil {
		utils.BadRequest(c, "不合理的文章ID")
		return
	}

	// 检查文章是否存在
	db := config.GetDBConect()
	post := models.Post{}
	if err := db.First(&post, postID).Error; err != nil {
		utils.BadRequest(c, "文章不存在")
		return
	}

	// 分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 50 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	// 获取评论列表
	var comments []models.Comment
	// 查询评论列表，预加载用户信息
	if err := config.DB.Preload("User").
		Where("post_id = ?", postID).
		Order("created_at ASC").
		Limit(pageSize).
		Offset(offset).
		Find(&comments).Error; err != nil {
		utils.InternalServerError(c, "获取评论列表失败")
		return
	}

	// 获取总数
	var total int64
	config.DB.Model(&models.Comment{}).Where("post_id = ?", postID).Count(&total)

	utils.Success(c, gin.H{
		"comments":  comments,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}
