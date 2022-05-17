package model

import (
	"gorm.io/gorm"
)

//话术表
type Craft struct {
	gorm.Model        //`json:"-"` //隐藏字段不输出
	Title      string `gorm:"size:255;type:char(255)"` // 设置字段大小为255
	A          string
	B          string
	Status     int64 `gorm:"type:int(1)"`
}

func GetCraft(ID interface{}) (Craft, error) {
	var Craft Craft
	result := DB.First(&Craft, ID)
	return Craft, result.Error
}
