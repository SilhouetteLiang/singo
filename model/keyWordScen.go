package model

import (
	"gorm.io/gorm"
)

//4. 关键字场景： 关键词，场景

// KeyWordScene	内容模型
type KeyWordScene struct {
	gorm.Model
	KeyWord string
	Scene   string
}

// GetKeyWordScene 用ID获取内容
func GetKeyWordScene(ID interface{}) (Content, error) {
	var content Content
	result := DB.First(&content, ID)
	return content, result.Error
}
