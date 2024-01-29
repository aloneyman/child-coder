package main

import (
	"child-coding-platform/backend/api"
	"child-coding-platform/backend/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect() // 确保数据库连接已经建立

	r := gin.Default()

	// 注册路由
	r.POST("/register", api.RegisterUser)
	r.POST("/login", api.LoginUser)
	r.POST("/course/create", api.AdminRequired(), api.CreateCourse)
	r.POST("/course/update", api.AdminRequired(), api.UpdateCourse)
	r.GET("/courses", api.GetCourses)
	r.GET("/course/:id", api.GetCourse)
	r.POST("/comment/create", api.CreateComment)
	r.GET("/comments/:course_id", api.GetCommentsByCourse)
	r.PUT("/comment/update/:id", api.UpdateComment)
	r.DELETE("/comment/delete/:id", api.DeleteComment)

	// 你可以在这里添加更多的路由

	r.Run() // 默认监听并在 0.0.0.0:8080 上启动服务
}
