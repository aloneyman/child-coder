package model

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	gorm.Model
	CourseID  uint      `json:"course_id"`
	UserID    uint      `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
