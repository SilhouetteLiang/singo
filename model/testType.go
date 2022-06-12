package model

import (
	"gorm.io/gorm"
)

//表格
//1. 测试结果说明：得分，解释，类型，影响
//`gorm:"size:255;type:char(255);not null;default:0;comment:内容"` // 设置字段大小为255

// content	内容模型
type TestType struct {
	gorm.Model
	Name    string `gorm:"size:255;type:char(255);not null;default:0;comment:测试题目名称"` // 设置字段大小为255
	TestNum int64  `gorm:"type:int(11);not null;default:0;comment:已使用过的测试人数 "`        // 设置字段大小为255
	Price   int64  `gorm:"type:int(11);not null;default:0;comment:测试价格"`              // 设置字段大小为255
	Status  int64  `gorm:"type:int(1);not null;default:0;comment:状态值 1正常 2异常"`
	//BeiYong1 string `gorm:"size:255;type:char(255);not null;default:0;comment:备用1"` // 设置字段大小为255
	//BeiYong2 string `gorm:"size:255;type:char(255);not null;default:0;comment:备用2"` // 设置字段大小为255
	//BeiYong3 string `gorm:"size:255;type:char(255);not null;default:0;comment:备用3"` // 设置字段大小为255
	//BeiYong4 string `gorm:"size:255;type:char(255);not null;default:0;comment:备用4"` // 设置字段大小为255
	//BeiYong5 string `gorm:"size:255;type:char(255);not null;default:0;comment:备用5"` // 设置字段大小为255
}
type TestInfo struct {
	Name    string `gorm:"size:255;type:char(255);not null;default:0;comment:测试题目名称"` // 设置字段大小为255
	TestNum int64  `gorm:"type:int(11);not null;default:0;comment:已使用过的测试人数 "`        // 设置字段大小为255
	Price   int64  `gorm:"type:int(11);not null;default:0;comment:测试价格"`              // 设置字段大小为255
}

// GetTestResult 用ID获取内容
func GetTestType(ID interface{}) (Content, error) {
	var content Content
	result := DB.First(&content, ID)
	return content, result.Error
}
