package model

import (
	"gorm.io/gorm"
)

//标签表
type Tag struct {
	gorm.Model `json:"-"` //隐藏字段不输出
	Title      string     `gorm:"size:255;type:char(255)"` // 设置字段大小为255
}

func GetTag(ID interface{}) (Tag, error) {
	var Tag Tag
	result := DB.First(&Tag, ID)
	return Tag, result.Error
}
