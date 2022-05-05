package model

import (
	//"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// ontent	内容模型
type Content struct {
	gorm.Model
	Content        string
	Status         string
	UserName       string
}


// GetConten 用ID获取内容
func GetContent(ID interface{}) (User, error) {
	var user User
	result := DB.First(&user, ID)
	return user, result.Error
}

// GetConten 用ID获取内容
func GetUsers(ID interface{}) (User, error) {
	var user User
	result := DB.First(&user, ID)
	return user, result.Error
}


