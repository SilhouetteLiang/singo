package model

import (
	"gorm.io/gorm"
)

//`gorm:"size:255;type:char(255);not null;default:0;comment:内容"` // 设置字段大小为255

// content	内容模型
type Content struct {
	gorm.Model
	Content  string `gorm:"size:255;type:char(255);not null;default:0;comment:内容"` // 设置字段大小为255
	Status   int64  `gorm:"type:int(1);not null;default:0;comment:状态值 1正常 2异常"`
	UserName string
}

// GetConten 用ID获取内容
func GetContent(ID interface{}) (Content, error) {
	var content Content
	result := DB.First(&content, ID)
	return content, result.Error
}
