package model

import (
	"gorm.io/gorm"
)

//`gorm:"type:int(1);not null;default:0;comment:状态值 1正常 2异常"`
//题库表
type LuntanComment struct {
	gorm.Model `json:"-"` //隐藏字段不输出
	LuntanId   int64      `gorm:"int(11);not null;default:99999999;comment:帖子ID"`
	Content    string     `gorm:"not null;default:内容;comment:评论内容"`
	UserId     int64      `gorm:"int(10);not null;default:100001;comment:用户ID"`
	Nickname   string     `gorm:"size:255;type:char(255);not null;default:张三;comment:用户昵称"`
	Status     int64      `gorm:"type:int(1);not null;default:0;comment:状态值 1正常 2异常"`
	Openid     string     `gorm:"size:255;type:char(255);not null;default:0;comment:用户openid"` // 设置字段大小为255
	IsZan      int64      `gorm:"type:int(1);not null;default:0;comment:用户是否点赞 1点赞2没有点赞"`
}
