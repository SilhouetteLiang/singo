package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//`gorm:"size:255;type:char(255);not null;default:0;comment:内容"` // 设置字段大小为255

// User 用户模型
type User struct {
	gorm.Model
	User           string `gorm:"size:255;type:char(255);not null;default:0;comment:用户"`         // 设置字段大小为255
	UserName       string `gorm:"size:255;type:char(255);not null;default:0;comment:用户真实名字"`     // 设置字段大小为255
	Nickname       string `gorm:"size:255;type:char(255);not null;default:0;comment:用户昵称"`       // 设置字段大小为255
	Avatar         string `gorm:"size:1000;not null;default:0;comment:用户头像"`                     // 设置字段大小为255
	PasswordDigest string `gorm:"size:255;type:char(255);not null;default:0;comment:用户密码"`       // 设置字段大小为255
	Tell           string `gorm:"size:255;type:char(255);not null;default:999999;comment:用户手机号"` // 设置字段大小为255
	DrawTag        string `gorm:"size:255;type:char(255);not null;default:0;comment:用户标签"`       // 设置字段大小为255
	Image          string `gorm:"size:255;type:char(255);not null;default:0;comment:用户补充图片"`     // 设置字段大小为255
	AccessNum      int64  `gorm:"size:255;type:char(255);not null;default:1;comment:用户访问次数"`     // 设置字段大小为255
	Status         string `gorm:"type:char(255);not null;default:0;comment:状态值 1正常 2异常"`
}

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
	// Active 激活用户
	Active string = "active"
	// Inactive 未激活用户
	Inactive string = "inactive"
	// Suspend 被封禁用户
	Suspend string = "suspend"
)

// GetUser 用ID获取用户
func GetUser(ID interface{}) (User, error) {
	var user User
	result := DB.First(&user, ID)
	return user, result.Error
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}
