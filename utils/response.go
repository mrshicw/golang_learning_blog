package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Success
func Success(c *gin.Context, data interface{}) {
	c.JSON(
		http.StatusOK,
		Response{
			Code:    200,
			Message: "success",
			Data:    data,
		})
}

func Error(c *gin.Context, code int, message string) {
	c.JSON(
		code,
		Response{
			Code:    code,
			Message: message,
		})
}

// 400
func BadRequest(c *gin.Context, message string) {
	Error(c, http.StatusBadRequest, message)
}

// 401
func Unauthorized(c *gin.Context, message string) {
	Error(c, http.StatusUnauthorized, message)
}

// 403
func Forbidden(c *gin.Context, message string) {
	Error(c, http.StatusForbidden, message)
}

// 404
func NotFound(c *gin.Context, message string) {
	Error(c, http.StatusNotFound, message)
}

// 500
func InternalServerError(c *gin.Context, message string) {
	Error(c, http.StatusInternalServerError, message)
}

// health
func Health(c *gin.Context) {
	c.JSON(200,
		gin.H{
			"status":  "OK",
			"message": "API is Running ... ",
		})
}
