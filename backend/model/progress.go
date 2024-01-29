package model

import (
	"gorm.io/gorm"
)

type Progress struct {
	gorm.Model
	CourseID  uint
	UserID    uint
	Completed bool
}
