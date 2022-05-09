package apiV3

import (
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

//subject/list
func SubjectList(c *gin.Context) {
	var service service.PsychologicalService

	if err := c.ShouldBind(&service); err == nil {
		res := service.GetSubjectList(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, " ErrorResponse(err)")
	}
}

// UserLogin 用户输入接口
func Input(c *gin.Context) {
	var service service.PsychologicalService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Input(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, "ErrorResponse")
	}
}
