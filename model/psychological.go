package model

import (
	//"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// ontent	内容模型
type Psychological struct {
	gorm.Model
	Status uint64
	Title  string
}
type ResPsychological struct {
	Items []Item
}

// PsychologicalService 管理用户新增内容服务
type Item struct {
	Title  string
	Status uint64
}

// GetPsychological 用ID获取内容
func GetPsychological(ID interface{}) (Psychological, error) {
	var psychological Psychological
	result := DB.First(&psychological, ID)
	return psychological, result.Error
}

// GetConten 用ID获取内容
//func GetUsers(ID interface{}) (User, error) {
//	var user User
//	result := DB.First(&user, ID)
//	return user, result.Error
//}
