package apiV10

import (
	"github.com/gin-gonic/gin"
	"singo/serializer"
)

// Ping 状态检查页面
func Pings(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "Pong",
		Data:	"这里是数据",
	})
}

func LotteryAlgorithm(c *gin.Context) {


	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "Pong",
		Data:	"这里是数据",
	})
}
