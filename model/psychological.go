package model

import (
	"gorm.io/gorm"
)

//`gorm:"size:255;type:char(255);not null;default:0;comment:内容"` // 设置字段大小为255

//题库表
type Psychological struct {
	gorm.Model `json:"-"` //隐藏字段不输出
	Title      string     `gorm:"size:255;type:char(255);not null;default:0;comment:标题"` // 设置字段大小为255
	A          string     `gorm:"not null;default:0;comment:选项A"`
	B          string     `gorm:"not null;default:0;comment:选项B"`
	C          string     `gorm:"not null;default:0;comment:选项C"`
	D          string     `gorm:"not null;default:0;comment:选项D"`
	Type       int64      `gorm:"type:int(1);not null;default:0;comment:类型 "`
	Status     int64      `gorm:"type:int(1);not null;default:0;comment:状态值 1正常 2异常"`
	BeiYong1   string     `gorm:"size:255;type:char(255);not null;default:0;comment:备用1"` // 设置字段大小为255
	BeiYong2   string     `gorm:"size:255;type:char(255);not null;default:0;comment:备用2"` // 设置字段大小为255
	BeiYong3   string     `gorm:"size:255;type:char(255);not null;default:0;comment:备用3"` // 设置字段大小为255
	BeiYong4   string     `gorm:"size:255;type:char(255);not null;default:0;comment:备用4"` // 设置字段大小为255
	BeiYong5   string     `gorm:"size:255;type:char(255);not null;default:0;comment:备用5"` // 设置字段大小为255

}

func GetPsychological(ID interface{}) (Psychological, error) {
	var psychological Psychological
	result := DB.First(&psychological, ID)
	return psychological, result.Error
}
