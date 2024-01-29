package api

import (
	"child-coding-platform/backend/database"
	"child-coding-platform/backend/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateCourse 处理创建课程的请求
func CreateCourse(c *gin.Context) {
	var course model.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 创建课程
	result := database.DB.Create(&course)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "课程创建成功", "course": course})
}

// UpdateCourse 处理修改课程的请求
func UpdateCourse(c *gin.Context) {
	// 需要从请求中获取课程 ID 和更新的信息
	// ...

	// 找到并更新课程
	// ...

	// 返回响应
	// ...
}

// GetCourses 处理获取课程列表的请求
func GetCourses(c *gin.Context) {
	var courses []model.Course
	result := database.DB.Find(&courses)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"courses": courses})
}

// GetCourse 处理获取单个课程详情的请求
func GetCourse(c *gin.Context) {
	courseID := c.Param("id")
	var course model.Course
	result := database.DB.First(&course, courseID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "课程未找到"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"course": course})
}
