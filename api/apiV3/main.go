package apiV3

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"singo/serializer"
	"singo/service"
)

// Ping 状态检查页面
func Pings(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "Pong",
		Data: "这里是V3数据",
	})
}

//题目列表
func SubjectList(c *gin.Context) {
	var service service.PsychologicalService

	if err := c.ShouldBind(&service); err == nil {
		res := service.GetSubjectList(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, " ErrorResponse(err)")
	}
}

//用户列表
func UserList(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello world",
		"tag":     11,
		//"tag":    ("亲子关系","亲密关系","职场晋升","刚毕业")
	})

}

func Result1(c *gin.Context) {
	ageMap := make(map[string]int, 10)
	ageMap["红"] = 19
	ageMap["黄"] = 21
	ageMap["蓝"] = 31
	ageMap["绿"] = 41
	c.JSON(200, gin.H{
		"code": 200,
		"data": ageMap,
		"msg":  "获取成功",
	})
}
func Result2(c *gin.Context) {
	ageMap := make(map[string]int, 10)
	ageMap["外倾感觉"] = 19
	ageMap["内倾感觉"] = 21
	ageMap["外倾直觉"] = 31
	ageMap["内倾直觉"] = 41
	ageMap["外倾思考"] = 31
	ageMap["内倾思考"] = 41
	ageMap["外倾情感"] = 31
	ageMap["内倾情感"] = 41
	c.JSON(200, gin.H{
		"code": 200,
		"data": ageMap,
		"msg":  "获取成功",
	})
}
func Result3(c *gin.Context) {
	ageMap := make(map[string]int, 10)
	ageMap["盈利能力"] = 19
	ageMap["成长能力"] = 19
	ageMap["现金能力"] = 39
	ageMap["偿债能力"] = 19
	ageMap["运营能力"] = 29
	c.JSON(200, gin.H{
		"code": 200,
		"data": ageMap,
		"msg":  "获取成功",
	})
}

// 输入题目
func InputTitle(c *gin.Context) {
	var service service.PsychologicalService
	if err := c.ShouldBind(&service); err == nil {
		fmt.Printf("service  : %v \n", service)
		res := service.InputTitle(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, "ErrorResponse")
	}
}

// 输入话术
func InputSpeechCraft(c *gin.Context) {

	var service service.SpeechSkillService
	if err := c.ShouldBind(&service); err == nil {
		fmt.Printf("service  : %v \n", service)
		res := service.InputSpeechCraft(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, "ErrorResponse")
	}

}

// 输入用户信息
func InputUserInfo(c *gin.Context) {
	//UserInfoService
	var service service.UserInfoService
	if err := c.ShouldBind(&service); err == nil {
		fmt.Printf("service  : %v \n", service)
		res := service.InputUserInfo(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, "ErrorResponse")
	}
}

// 首页
func Index(c *gin.Context) {
	array := [4]string{"亲子关系", "亲密关系", "职场晋升", "刚毕业"}
	c.JSON(200, gin.H{
		"code": 200,
		"data": array,
		"msg":  "获取成功",
	})
}

//论坛新增
func LuntanAdd(c *gin.Context) {
	var service service.Luntan
	if err := c.ShouldBind(&service); err == nil {
		fmt.Printf("service  : %v \n", service)
		res := service.LuntanAdd(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, "ErrorResponse")
	}
}
//论坛列表
func LuntanList(c *gin.Context) {
	array := [4]string{"亲子关系", "亲密关系", "职场晋升", "刚毕业"}
	c.JSON(200, gin.H{
		"code": 200,
		"data": array,
		"msg":  "获取成功",
	})
}
