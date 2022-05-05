package service

import (
	//"singo/model"
//	//"singo/serializer"
	"github.com/gin-gonic/gin"
	"singo/model"
	"singo/serializer"
)

// UserInputService 管理用户新增内容服务
type UserInputService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Content string `form:"content" json:"content" binding:"required"`
}



//用户输入内容
func (service *UserInputService) InputContent(c *gin.Context) serializer.Response {
	content :=model.Content{
		Content: service.Content,
		UserName: service.UserName,
	}
	// 新增数据
	if err := model.DB.Create(&content).Error; err != nil {
		return serializer.ParamErr("新增失败", err)
	}
	return serializer.BuildUserResponses(content)
}


//Register 用户注册
//func (service *UserRegisterService) Register() serializer.Response {
//	user := model.User{
//		Nickname: service.Nickname,
//		UserName: service.UserName,
//		Status:   model.Active,
//	}
//
//	// 表单验证
//	if err := service.valid(); err != nil {
//		return *err
//	}
//
//	// 加密密码
//	if err := user.SetPassword(service.Password); err != nil {
//		return serializer.Err(
//			serializer.CodeEncryptError,
//			"密码加密失败",
//			err,
//		)
//	}
//
//	// 创建用户
//	if err := model.DB.Create(&user).Error; err != nil {
//		return serializer.ParamErr("注册失败", err)
//	}
//
//	return serializer.BuildUserResponse(user)
//}

































