package service

import (
	"github.com/gin-gonic/gin"
	"singo/model"
	"singo/serializer"
)

// PsychologicalService 管理用户新增内容服务
type PsychologicalService struct {
	Title  string `form:"title" json:"title" binding:"required,min=5,max=30"`
	Status uint64 `form:"status" json:"status" binding:"required"`
}

//type ResPsychological struct {
//	Item []Item
//}
//
//// PsychologicalService 管理用户新增内容服务
//type Item struct {
//	Title  string `form:"title" json:"title" binding:"required,min=5,max=30"`
//	Status uint64 `form:"status" json:"status" binding:"required"`
//}
//content := model.Content{
//Content:  service.Content,
//UserName: service.UserName,
//}
//获取题目列表
func (service *PsychologicalService) GetSubjectList(c *gin.Context) serializer.Response {
	//psychological := model.ResPsychological{
	//	Items: []model.Item{
	//		Title:  service.Title,
	//		Status: service.Status,
	//	},
	//}
	psychological := model.Psychological{
		Status: service.Status,
		Title:  service.Title,
	}
	// 获取全部数据
	if err := model.DB.Find(&psychological).Error; err != nil {
		return serializer.ParamErr("获取数据失败", err)
	}
	//fmt.Printf("res : %v \n", psychological)
	//fmt.Printf("1111111111111 : %v \n", psychological)
	//fmt.Printf("2222222222222 : %v \n", psychological)

	return serializer.BuildPsychologicalResponses(psychological)
}

//用户输入内容
func (service *PsychologicalService) Input(c *gin.Context) serializer.Response {
	psychological := model.Psychological{
		Title:  service.Title,
		Status: service.Status,
	}
	// 新增数据
	if err := model.DB.Create(&psychological).Error; err != nil {
		return serializer.ParamErr("新增失败", err)
	}
	return serializer.BuildPsychologicalResponses(psychological)
}
