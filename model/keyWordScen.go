package model

import (
	"gorm.io/gorm"
)

//4. 关键字场景： 关键词，场景
//`gorm:"size:255;type:char(255);not null;default:0;comment:内容"` // 设置字段大小为255
// KeyWordScene
type KeyWordScene struct {
	gorm.Model
	KeyWord  string `gorm:"size:255;type:char(255);not null;default:0;comment:关键词"`
	Scene    string `gorm:"size:255;type:char(255);not null;default:0;comment:场景"`
	Status   int64  `gorm:"type:int(1);not null;default:0;comment:状态值 1正常 2异常"`
	BeiYong1 string `gorm:"size:255;type:char(255);not null;default:0;comment:BeiYong1"`
	BeiYong2 string `gorm:"size:255;type:char(255);not null;default:0;comment:BeiYong2"` // 设置字段大小为255
	BeiYong3 string `gorm:"size:255;type:char(255);not null;default:0;comment:BeiYong3"` // 设置字段大小为255
	BeiYong4 string `gorm:"size:255;type:char(255);not null;default:0;comment:BeiYong4"` // 设置字段大小为255
	BeiYong5 string `gorm:"size:255;type:char(255);not null;default:0;comment:BeiYong5"` // 设置字段大小为255
}

type Scene struct {
	Scene   string `gorm:"size:255;type:char(255);not null;default:0;comment:场景"`
	KeyWord string `gorm:"size:255;type:char(255);not null;default:0;comment:关键词"`
}

//Name string `gorm:"type:char(30);comment:'姓名'"`

// GetKeyWordScene 用ID获取内容
func GetKeyWordScene(ID interface{}) (Content, error) {
	var content Content
	result := DB.First(&content, ID)
	return content, result.Error
}
