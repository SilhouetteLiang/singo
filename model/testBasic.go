package model

import (
	"gorm.io/gorm"
)

//2. 测试基本信息：类型，计分规则，测试种类

// content	内容模型
type TestBasic struct {
	gorm.Model
	Type         int64
	ScoringRules string
	TestType     int64
}

// GetTestBasic 用ID获取内容
func GetTestBasic(ID interface{}) (Content, error) {
	var content Content
	result := DB.First(&content, ID)
	return content, result.Error
}
