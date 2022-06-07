package model
import (
	"gorm.io/gorm"
)

//题库表
type Luntan struct {
	gorm.Model `json:"-"` //隐藏字段不输出
	Title      string     `gorm:"size:255;type:char(255)"` // 设置字段大小为255
	Content    string
	Img        string
	UserId     string
	Status     int64
}