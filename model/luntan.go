package model

import (
	"gorm.io/gorm"
)

//`gorm:"type:int(1);not null;default:0;comment:状态值 1正常 2异常"`
//题库表
type Luntan struct {
	gorm.Model `json:"-"` //隐藏字段不输出
	Title      string     `gorm:"size:255;type:char(255);not null;default:标题;comment:标题"` // 设置字段大小为255
	Content    string     `gorm:"not null;default:内容;comment:内容"`
	Img        string     `gorm:"not null;default:Img;comment:论坛图片地址"`
	UserId     string     `gorm:"int(10);not null;default:100001;comment:用户ID"`
	Nickname   string     `gorm:"size:255;type:char(255);not null;default:张三;comment:用户昵称"`
	Status     int64      `gorm:"type:int(1);not null;default:0;comment:状态值 1正常 2异常"`
	Sort       int64      `gorm:"type:int(1);not null;default:2;comment:是否置顶 1 是 2否"`
	Zan        int64      `gorm:"type:int(11);not null;default:1;comment:点赞数量"`
	CommentNum int64      `gorm:"type:int(11);default:0;comment:评论数量"`
}
