package model

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string
	Description string
	Courses     []Course `gorm:"many2many:course_categories;"`
}
