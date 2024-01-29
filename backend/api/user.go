package api

import (
	"child-coding-platform/backend/database"
	"child-coding-platform/backend/model"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)

	// 对密码进行哈希处理
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	// 创建用户
	result := database.DB.Create(&user)
	if result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": result.Error.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "注册成功"})
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var inputUser model.User
	var foundUser model.User
	json.NewDecoder(r.Body).Decode(&inputUser)

	// 根据用户名查找用户
	result := database.DB.Where("username = ?", inputUser.Username).First(&foundUser)
	if result.Error != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "用户名或密码错误"})
		return
	}

	// 检查密码
	err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(inputUser.Password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "用户名或密码错误"})
		return
	}

	// 返回响应
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "登录成功"})
}
