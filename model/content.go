package model

import (
	"gorm.io/gorm"
)

// content	内容模型
type Content struct {
	gorm.Model
	Content  string
	Status   string
	UserName string
}

// GetConten 用ID获取内容
func GetContent(ID interface{}) (Content, error) {
	var content Content
	result := DB.First(&content, ID)
	return content, result.Error
}
