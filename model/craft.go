package model

import (
	"gorm.io/gorm"
)

//`gorm:"size:255;type:char(255);not null;default:0;comment:内容"` // 设置字段大小为255
//话术表
type Craft struct {
	gorm.Model        //`json:"-"` //隐藏字段不输出
	Object     string `gorm:"size:255;type:char(255);not null;default:0;comment:对象"` // 设置字段大小为255
	Scene      string `gorm:"size:255;type:char(255);not null;default:0;comment:场景"` // 设置字段大小为255
	Content    string `gorm:"not null;default:0;comment:内容"`
	Source     string `gorm:"size:255;type:char(255);not null;default:0;comment:来源"`  // 设置字段大小为255
	Keyword    string `gorm:"size:255;type:char(255);not null;default:0;comment:关键词"` // 设置字段大小为255
	Status     int64  `gorm:"type:int(1);not null;default:0;comment:状态值 1正常 2异常"`
	Title      string `gorm:"not null;default:0;comment:标题"`
	BeiYong1   string `gorm:"size:255;type:char(255);not null;default:0;comment:BeiYong1"` // 设置字段大小为255
	BeiYong2   string `gorm:"size:255;type:char(255);not null;default:0;comment:BeiYong2"` // 设置字段大小为255
	BeiYong3   string `gorm:"size:255;type:char(255);not null;default:0;comment:BeiYong3"` // 设置字段大小为255
	BeiYong4   string `gorm:"size:255;type:char(255);not null;default:0;comment:BeiYong4"` // 设置字段大小为255
	BeiYong5   string `gorm:"size:255;type:char(255);not null;default:0;comment:BeiYong5"` // 设置字段大小为255
}
type Crafts struct {
	Content string `gorm:"not null;default:0;comment:内容"`
	Source  string `gorm:"size:255;type:char(255);not null;default:0;comment:来源"` // 设置字段大小为255
}

func GetCraft(ID interface{}) (Craft, error) {
	var Craft Craft
	result := DB.First(&Craft, ID)
	return Craft, result.Error
}
