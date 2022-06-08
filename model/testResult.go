package model

import (
	"gorm.io/gorm"
)

//表格
//1. 测试结果说明：得分，解释，类型，影响

// content	内容模型
type TestResult struct {
	gorm.Model
	Grade     int64
	Explian   string
	Type      int64
	Influence string
}

// GetTestResult 用ID获取内容
func GetTestResult(ID interface{}) (Content, error) {
	var content Content
	result := DB.First(&content, ID)
	return content, result.Error
}
