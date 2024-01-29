package api

import (
	"child-coding-platform/backend/database"
	"child-coding-platform/backend/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateComment 处理创建评论的请求
func CreateComment(c *gin.Context) {
	var comment model.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 创建评论
	result := database.DB.Create(&comment)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "评论创建成功", "comment": comment})
}

// GetCommentsByCourse 处理获取特定课程的所有评论的请求
func GetCommentsByCourse(c *gin.Context) {
	courseID := c.Param("course_id")
	var comments []model.Comment
	result := database.DB.Where("course_id = ?", courseID).Find(&comments)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"comments": comments})
}

// UpdateComment 处理编辑评论的请求
func UpdateComment(c *gin.Context) {
	commentID := c.Param("id")
	var comment model.Comment
	if err := database.DB.First(&comment, commentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "评论未找到"})
		return
	}

	// TODO: 确保当前用户是评论的作者或管理员

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&comment)

	c.JSON(http.StatusOK, gin.H{"message": "评论更新成功", "comment": comment})
}

// DeleteComment 处理删除评论的请求
func DeleteComment(c *gin.Context) {
	commentID := c.Param("id")
	var comment model.Comment
	if err := database.DB.First(&comment, commentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "评论未找到"})
		return
	}

	// TODO: 确保当前用户是评论的作者或管理员

	database.DB.Delete(&comment)

	c.JSON(http.StatusOK, gin.H{"message": "评论删除成功"})
}

func Default(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "hello coder"})
}
