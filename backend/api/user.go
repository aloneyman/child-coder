package api

import (
	"child-coding-platform/backend/database"
	"child-coding-platform/backend/model"
	"encoding/json"
	"net/http"
)

// RegisterUser 处理用户注册请求
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)

	// 密码加密等处理

	// 创建用户
	database.DB.Create(&user)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// LoginUser 处理用户登录请求
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	var foundUser model.User
	json.NewDecoder(r.Body).Decode(&user)

	// 根据用户名查找用户
	database.DB.Where("username = ?", user.Username).First(&foundUser)

	// 检查密码等验证逻辑

	// 返回响应
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(foundUser)
}
