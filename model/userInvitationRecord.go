package model

import (
	"gorm.io/gorm"
)

//`gorm:"size:255;type:char(255);not null;default:0;comment:内容"` // 设置字段大小为255

// UserInvitationRecord	用户邀请记录表
type UserInvitationRecord struct {
	gorm.Model
	Status             int64  `gorm:"type:int(1);not null;default:0;comment:状态值 1正常 2异常"`
	UserName           string `gorm:"size:255;type:char(255);not null;default:0;comment:用户名字"` // 设置字段大小为255
	NickName           string `gorm:"size:255;type:char(255);not null;default:0;comment:用户昵称"` // 设置字段大小为255
	OpenId             string `gorm:"size:255;type:char(255);not null;comment:用户opendid"`      // 设置字段大小为255
	InvitationId       string `gorm:"size:255;type:char(255);not null;default:0;comment:邀请用户的openid"`
	InvitationNickName string `gorm:"size:255;type:char(255);not null;default:0;comment:邀请用户的昵称"`
	InvitationName     string `gorm:"size:255;type:char(255);not null;default:0;comment:邀请用户的名字"`
	InvitationAvatar   string `gorm:"size:255;type:char(255);not null;default:0;comment:邀请用户的头像"`
	InvitationIntegral string `gorm:"size:255;type:char(255);not null;default:0;comment:邀请用户的积分指数"`
}

type UserInvitationInfo struct {
	InvitationId       string `gorm:"size:255;type:char(255);not null;default:0;comment:邀请用户的openid"`
	InvitationNickName string `gorm:"size:255;type:char(255);not null;default:0;comment:邀请用户的昵称"`
	InvitationAvatar   string `gorm:"size:255;type:char(255);not null;default:0;comment:邀请用户的头像"`
	InvitationIntegral string `gorm:"size:255;type:char(255);not null;default:0;comment:邀请用户的积分指数"`
}

// GetUserInvitationRecord 用ID获取内容 用户邀请记录表
func GetUserInvitationRecord(ID interface{}) (UserInvitationRecord, error) {
	var UserInvitationRecord UserInvitationRecord
	result := DB.First(&UserInvitationRecord, ID)
	return UserInvitationRecord, result.Error
}
