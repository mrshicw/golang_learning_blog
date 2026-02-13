package controllers

import (
	"golang_learning_blog/models"
	"golang_learning_blog/utils"

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
	//
	post := models.Post{}
	utils.Success(c, post)
}

// 获取文章列表
func (pc *PostController) GetPosts(c *gin.Context) {
	//
	utils.Success(c, gin.H{})
}

// 获取单个文章详情
func (pc *PostController) GetPost(c *gin.Context) {
	//
	post := models.Post{}
	utils.Success(c, post)
}

// 更新文章
func (pc *PostController) UpdatePost(c *gin.Context) {
	//
	post := models.Post{}
	utils.Success(c, post)
}

// 删除文章
func (pc *PostController) DeletePost(c *gin.Context) {
	//
	utils.Success(c, gin.H{})
}
