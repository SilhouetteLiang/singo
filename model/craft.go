package model

import (
	"gorm.io/gorm"
)

//话术表
type Craft struct {
	gorm.Model        //`json:"-"` //隐藏字段不输出
	Status     int64  `gorm:"type:int(1)"`
	object     string `gorm:"size:255;type:char(255)"` // 设置字段大小为255
	scene      string `gorm:"size:255;type:char(255)"` // 设置字段大小为255
	content    string
	source     string `gorm:"size:255;type:char(255)"` // 设置字段大小为255
	keyword    string `gorm:"size:255;type:char(255)"` // 设置字段大小为255
	Title      string `gorm:"size:255;type:char(255)"` // 设置字段大小为255
}

func GetCraft(ID interface{}) (Craft, error) {
	var Craft Craft
	result := DB.First(&Craft, ID)
	return Craft, result.Error
}
