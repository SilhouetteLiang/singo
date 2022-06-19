package model

import (
	"gorm.io/gorm"
)

//`gorm:"size:255;type:char(255);not null;default:0;comment:内容"` // 设置字段大小为255

// content	内容模型
type UserReport struct {
	gorm.Model
	Grade         int64  `gorm:"type:int(10);not null;default:0;comment:用户得分"`
	Status        int64  `gorm:"type:int(1);not null;default:0;comment:状态值 1正常 2异常"`
	UserName      string `gorm:"size:255;type:char(255);not null;default:0;comment:用户名字"`   // 设置字段大小为255
	NickName      string `gorm:"size:255;type:char(255);not null;default:0;comment:用户昵称"`   // 设置字段大小为255
	Uid           int64  `gorm:"type:int(1);not null;default:10001;comment:用户ID"`           // 设置字段大小为255
	OpenId        string `gorm:"size:255;type:char(255);not null;comment:用户opendid"`        // 设置字段大小为255
	ReportContent string `gorm:"size:255;type:char(255);not null;default:0;comment:报告附加内容"` // 设置字段大小为255
	ReportType    int64  `gorm:"int(10);not null;default:1;comment:报告类型"`                   // 设置字段大小为255
	ReportName    string `gorm:"size:255;type:char(255);not null;default:0;comment:报告名字"`   // 设置字段大小为255
}

type UserReportList struct {
	Grade         int64  `gorm:"type:int(10);not null;default:0;comment:用户得分"`
	Status        int64  `gorm:"type:int(1);not null;default:0;comment:状态值 1正常 2异常"`
	UserName      string `gorm:"size:255;type:char(255);not null;default:0;comment:用户名字"`   // 设置字段大小为255
	NickName      string `gorm:"size:255;type:char(255);not null;default:0;comment:用户昵称"`   // 设置字段大小为255
	Uid           int64  `gorm:"type:int(1);not null;default:10001;comment:用户ID"`           // 设置字段大小为255
	ReportContent string `gorm:"size:255;type:char(255);not null;default:0;comment:报告附加内容"` // 设置字段大小为255
	ReportType    int64  `gorm:"int(10);not null;default:1;comment:报告类型"`                   // 设置字段大小为255
	ReportName    string `gorm:"size:255;type:char(255);not null;default:0;comment:报告名字"`   // 设置字段大小为255
}

// GetUserReport 用ID获取内容
func GetUserReport(ID interface{}) (Content, error) {
	var content Content
	result := DB.First(&content, ID)
	return content, result.Error
}
