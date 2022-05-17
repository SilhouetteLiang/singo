package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"singo/model"
	"singo/serializer"
)

// 新增题目
type PsychologicalService struct {
	A      string
	B      string
	C      string
	D      string
	Title  string `form:"title" json:"title"`
	Status uint64 `form:"status" json:"status"`
}

// 输入用户信息
type UserInfoService struct {
	Nickname  string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	UserName  string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Avatar    string `form:"avatar" json:"avatar"`
	Tell      string `form:"tell" json:"tell"`
	DrawTag   string `form:"draw_tag" json:"draw_tag"`
	AccessNum int64  `form:"access_num" json:"access_num"`
}

//新增话术
type SpeechSkillService struct {
	A      string
	B      string
	Title  string `form:"title" json:"title"`
	Status int64  `form:"status" json:"status"`
}

type Item struct {
	Title  string
	Status uint64
}

var db *gorm.DB

//获取题目列表
func (service *PsychologicalService) GetSubjectList(c *gin.Context) serializer.Response {
	//var user model.ResPsychological
	psychological := make([]model.Psychological, 0)

	// 获取全部数据
	if err := model.DB.Find(&psychological).Error; err != nil {
		return serializer.ParamErr("获取数据失败", err)
	}
	//打印结果
	//fmt.Printf("result : %v \n", psychological)

	return serializer.BuildPsychologicalResponses(psychological)
}

//输入题目
func (service *PsychologicalService) InputTitle(c *gin.Context) serializer.Response {
	fmt.Printf("service  %v  \n", service)

	psychological := model.Psychological{
		Title:  service.Title,
		Status: service.Status,
		A:      service.A,
		B:      service.B,
		C:      service.C,
		D:      service.D,
	}
	fmt.Printf("psychological  %v  \n", psychological)

	// 新增数据
	if err := model.DB.Create(&psychological).Error; err != nil {
		return serializer.ParamErr("新增失败", err)
	}
	return serializer.BuildPsychologicalResponse(psychological)
}

//输入话术
func (service *SpeechSkillService) InputSpeechCraft(c *gin.Context) serializer.Response {
	fmt.Printf("service  %v  \n", service)

	speechSkillService := model.Craft{
		Title:  service.Title,
		Status: service.Status,
		A:      service.A,
		B:      service.B,
	}
	fmt.Printf("speechSkillService  %v  \n", speechSkillService)

	// 新增数据
	if err := model.DB.Create(&speechSkillService).Error; err != nil {
		return serializer.ParamErr("新增失败", err)
	}
	return serializer.BuildCraftResponse(speechSkillService)
}

//新增用户
func (service *UserInfoService) InputUserInfo(c *gin.Context) serializer.Response {
	fmt.Printf("service  %v  \n", service)

	user := model.User{
		Nickname:  service.Nickname,
		Avatar:    service.Avatar,
		Tell:      service.Tell,
		AccessNum: service.AccessNum,
		DrawTag:   service.DrawTag,
	}

	fmt.Printf("user  %v  \n", user)

	// 新增数据
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.ParamErr("新增失败", err)
	}
	return serializer.BuildUsersResponse(user)
}
