package model

import (
	"gorm.io/gorm"
)

//表格
//1. 测试结果说明：得分，解释，类型，影响
//`gorm:"size:255;type:char(255);not null;default:0;comment:内容"` // 设置字段大小为255

// content	内容模型
type TestResult struct {
	gorm.Model
	Grade     string `gorm:"type:char(100);not null;default:0;comment:得分"` // 设置字段大小为255
	Explian   string `gorm:"not null;default:0;comment:解释"`                // 设置字段大小为255
	Type      int64  `gorm:"type:int(1);not null;default:0;comment:类型 "`   // 设置字段大小为255
	Influence string `gorm:"not null;default:0;comment:影响"`                // 设置字段大小为255
	Status    int64  `gorm:"type:int(1);not null;default:0;comment:状态值 1正常 2异常"`
	BeiYong1  string `gorm:"size:255;type:char(255);not null;default:0;comment:备用1"` // 设置字段大小为255
	BeiYong2  string `gorm:"size:255;type:char(255);not null;default:0;comment:备用2"` // 设置字段大小为255
	BeiYong3  string `gorm:"size:255;type:char(255);not null;default:0;comment:备用3"` // 设置字段大小为255
	BeiYong4  string `gorm:"size:255;type:char(255);not null;default:0;comment:备用4"` // 设置字段大小为255
	BeiYong5  string `gorm:"size:255;type:char(255);not null;default:0;comment:备用5"` // 设置字段大小为255

}

// GetTestResult 用ID获取内容
func GetTestResult(ID interface{}) (Content, error) {
	var content Content
	result := DB.First(&content, ID)
	return content, result.Error
}
