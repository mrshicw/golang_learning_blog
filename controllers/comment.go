package controllers

import (
	"golang_learning_blog/models"
	"golang_learning_blog/utils"

	"github.com/gin-gonic/gin"
)

type CommentController struct{}

type CreateCommentRequest struct {
	Content string `json:"content" binding:"required,min=1,max=1000"`
}

// 创建评论
func (cc *CommentController) CreateComment(c *gin.Context) {
	//

	comment := models.Comment{}
	utils.Success(c, comment)
}

// 获得文章的评论列表
func (cc *CommentController) GetComments(c *gin.Context) {
	//
	utils.Success(c, gin.H{})
}
