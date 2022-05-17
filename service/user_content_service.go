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
	Content  string `form:"content" json:"content" binding:"required"`
}

//用户输入内容
func (service *UserInputService) InputContent(c *gin.Context) serializer.Response {
	content := model.Content{
		Content:  service.Content,
		UserName: service.UserName,
	}
	// 新增数据
	if err := model.DB.Create(&content).Error; err != nil {
		return serializer.ParamErr("新增失败", err)
	}
	return serializer.BuildUserResponses(content)
}
