package model

import "time"

// User 定义了用户模型
type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Username  string    `json:"username" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	Role      string    `json:"role"` // 新增字段，可以是 "admin", "teacher", "student"
}
