package model

// Course 定义了课程模型
type Course struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"` // 课程内容，可以是文本或者是更复杂的结构
}
