package model

import (
	"gorm.io/gorm"
)

//3. 对象：名称，编号
// Object	内容模型
type Object struct {
	gorm.Model
	Name   string
	Number int64
}

// GetObject 用ID获取内容
func GetObject(ID interface{}) (Content, error) {
	var content Content
	result := DB.First(&content, ID)
	return content, result.Error
}
