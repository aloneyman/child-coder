package database

import (
	"child-coding-platform/backend/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(sqlite.Open("child_coding_platform.db"), &gorm.Config{})
	if err != nil {
		panic("无法连接数据库")
	}

	DB = connection

	connection.AutoMigrate(&model.User{}, &model.Course{})
}
