package model

import (
	"gorm.io/gorm"
)

//题库表
type Psychological struct {
	gorm.Model `json:"-"` //隐藏字段不输出
	Status     uint64     `gorm:"type:int(1)"`
	Title      string     `gorm:"size:255;type:char(255)"` // 设置字段大小为255
	A          string
	B          string
	C          string
	D          string
}

func GetPsychological(ID interface{}) (Psychological, error) {
	var psychological Psychological
	result := DB.First(&psychological, ID)
	return psychological, result.Error
}
