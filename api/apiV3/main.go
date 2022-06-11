package apiV3

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"singo/serializer"
	"singo/service"
	"strconv"
)

type ret struct {
	object []string
	scens  []string
}

//  状态检查页面
func Ping(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "Pong",
		Data: "这里是V3数据",
	})
}

// 1.首页 获取场景
func Index(c *gin.Context) {
	var service service.IndexSearchService
	if err := c.ShouldBind(&service); err == nil {
		fmt.Printf("2  : %v \n", service)
		res := service.Index(c) //TestSelect
		c.JSON(200, res)
	} else {
		c.JSON(200, "ErrorResponse")
	}

	//array := [9]string{"亲子关系", "亲密关系", "职场晋升", "刚毕业", "怎么和领导相处", "学习", "干家务", "感恩", "爱"}
	//c.JSON(200, gin.H{
	//	"code": 200,
	//	"msg":  "获取成功",
	//	"data": array,
	//})
}

//2.首页搜索
func IndexSearch(c *gin.Context) {
	var service service.IndexSearchService
	if err := c.ShouldBind(&service); err == nil {
		fmt.Printf("2  : %v \n", service)
		res := service.IndexSearch(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, "ErrorResponse")
	}
}

// 3.首页 获取对象和场景
func IndexBasicInfo(c *gin.Context) {
	object := []string{"爸爸", "妈妈", "孩子", "父亲", "母亲", "儿子", "女儿", "朋友", "领导", "老师", "老公", "老婆", "男朋友", "女朋友"}
	scens := []string{"亲子关系", "亲密关系", "职场晋升", "刚毕业", "怎么和领导相处"}
	//mapppp["scene"] = "3213"
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": &ret{
			object: object,
			scens:  scens,
		},
		"object": object,
		"scens":  scens,
	})
}

//4.首页 发布话术
func IndexPublishScript(c *gin.Context) {
	var service service.SpeechSkillService
	if err := c.ShouldBind(&service); err == nil {
		fmt.Printf("service  : %v \n", service)
		res := service.InputSpeechCraft(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, "ErrorResponse")
	}

	//array := [5]string{"亲子关系", "亲密关系", "职场晋升", "刚毕业", "怎么和领导相处"}
	//c.JSON(200, gin.H{
	//	"code": 200,
	//	"msg":  "获取成功",
	//	"data": array,
	//})
}

//5.论坛 列表
func ForumList(c *gin.Context) {
	var service service.Luntan
	if err := c.ShouldBind(&service); err == nil {
		res := service.LuntanList(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, " ErrorResponse(err)")
	}
	//array := [5]string{"亲子关系", "亲密关系", "职场晋升", "刚毕业", "怎么和领导相处"}
	//c.JSON(200, gin.H{
	//	"code": 200,
	//	"msg":  "获取成功",
	//	"data": array,
	//})
}

//6.论坛 发布论坛
func ForumPublishInvitation(c *gin.Context) {
	var service service.Luntan
	if err := c.ShouldBind(&service); err == nil {
		fmt.Printf("service  : %v \n", service)
		res := service.LuntanAdd(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, "ErrorResponse")
	}
	//array := [5]string{"亲子关系", "亲密关系", "职场晋升", "刚毕业", "怎么和领导相处"}
	//c.JSON(200, gin.H{
	//	"code": 200,
	//	"msg":  "获取成功",
	//	"data": array,
	//})
}

//7.论坛 评论列表
func ForumCommentList(c *gin.Context) {
	commentId := c.Request.URL.Query().Get("commentId")
	id, _ := strconv.ParseInt(commentId, 10, 64)
	//name := c.Param("name")
	//c.String(http.StatusOK, "Hello %s", id)
	var service service.IndexSearchService
	if err := c.ShouldBind(&service); err == nil {
		fmt.Printf("2  : %v \n", service)
		res := service.ForumCommentList(c, id) //TestSelect
		c.JSON(200, res)
	} else {
		c.JSON(200, "ErrorResponse")
	}
	//array := [5]string{"亲子关系", "亲密关系", "职场晋升", "刚毕业", "怎么和领导相处"}
	//c.JSON(200, gin.H{
	//	"code": 200,
	//	"msg":  "获取成功",
	//	"data": array,
	//})
}

//8.论坛 发布评论
func ForumPublishComment(c *gin.Context) {
	//ForumPublishCommentService
	//var service service.ForumPublishCommentService
	//if err := c.ShouldBind(&service); err == nil {
	//	fmt.Printf("2  : %v \n", service)
	//	res := service.liang(c) //TestSelect
	//	c.JSON(200, res)
	//} else {
	//	c.JSON(200, "ErrorResponse")
	//}

	var service service.ForumPublishCommentService
	if err := c.ShouldBind(&service); err == nil {
		fmt.Printf("service  : %v \n", service)
		res := service.ForumPublishComment(c)
		c.JSON(200, res)
	} else {
		fmt.Printf("err  : %v \n", err)

		c.JSON(200, "ErrorResponse")
	}

	//array := [5]string{"亲子关系", "亲密关系", "职场晋升", "刚毕业", "怎么和领导相处"}
	//c.JSON(200, gin.H{
	//	"code": 200,
	//	"msg":  "获取成功",
	//	"data": array,
	//})
}

//8.论坛 点赞
func ForumZan(c *gin.Context) {
	var service service.ForumZanService
	if err := c.ShouldBind(&service); err == nil {
		fmt.Printf("service  : %v \n", service)
		res := service.ForumZan(c)
		c.JSON(200, res)
	} else {
		fmt.Printf("err  : %v \n", err)

		c.JSON(200, "ErrorResponse")
	}
	//array := [5]string{"亲子关系", "亲密关系", "职场晋升", "刚毕业", "怎么和领导相处"}
	//c.JSON(200, gin.H{
	//	"code": 200,
	//	"msg":  "获取成功",
	//	"data": array,
	//})
}

//9.测评 首页
func EvaluationIndex(c *gin.Context) {
	array := [5]string{"性格色彩", "MBTI测试", "快乐指数", "情商测试"}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": array,
	})
}

//10.测评 性格测试
func EvaluationXingge(c *gin.Context) {

	commentId := c.Request.URL.Query().Get("type")
	id, _ := strconv.ParseInt(commentId, 10, 64)
	//name := c.Param("name")
	//c.String(http.StatusOK, "Hello %s", id)
	var service service.PsychologicalService
	if err := c.ShouldBind(&service); err == nil {
		res := service.EvaluationXingge(c, id) //TestSelect
		c.JSON(200, res)
	} else {
		c.JSON(200, "ErrorResponse")
	}

	//array := [5]string{"亲子关系", "亲密关系", "职场晋升", "刚毕业", "怎么和领导相处"}
	//c.JSON(200, gin.H{
	//	"code": 200,
	//	"msg":  "获取成功",
	//	"data": array,
	//})
}

//11.测评 MBTI测试
func EvaluationMBTI(c *gin.Context) {
	var service service.PsychologicalService
	if err := c.ShouldBind(&service); err == nil {
		res := service.EvaluationMBTI(c) //TestSelect
		c.JSON(200, res)
	} else {
		c.JSON(200, "ErrorResponse")
	}
	//array := [5]string{"亲子关系", "亲密关系", "职场晋升", "刚毕业", "怎么和领导相处"}
	//c.JSON(200, gin.H{
	//	"code": 200,
	//	"msg":  "获取成功",
	//	"data": array,
	//})
}

//12.测评 快乐指数测试
func EvaluationKuaile(c *gin.Context) {
	var service service.PsychologicalService
	if err := c.ShouldBind(&service); err == nil {
		res := service.EvaluationKuaile(c) //TestSelect
		c.JSON(200, res)
	} else {
		c.JSON(200, "ErrorResponse")
	}
	//array := [5]string{"亲子关系", "亲密关系", "职场晋升", "刚毕业", "怎么和领导相处"}
	//c.JSON(200, gin.H{
	//	"code": 200,
	//	"msg":  "获取成功",
	//	"data": array,
	//})
}

//13.测评 情商测试
func EvaluationQingshang(c *gin.Context) {
	var service service.PsychologicalService
	if err := c.ShouldBind(&service); err == nil {
		res := service.EvaluationQingshang(c) //TestSelect
		c.JSON(200, res)
	} else {
		c.JSON(200, "ErrorResponse")
	}
	//array := [5]string{"亲子关系", "亲密关系", "职场晋升", "刚毕业", "怎么和领导相处"}
	//c.JSON(200, gin.H{
	//	"code": 200,
	//	"msg":  "获取成功",
	//	"data": array,
	//})
}

//14.测评 性格测试
func EvaluationXingges(c *gin.Context) {
	array := [5]string{"亲子关系", "亲密关系", "职场晋升", "刚毕业", "怎么和领导相处"}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": array,
	})
}

//15.测评 MBTI测试
func EvaluationMBTIs(c *gin.Context) {
	array := [5]string{"亲子关系", "亲密关系", "职场晋升", "刚毕业", "怎么和领导相处"}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": array,
	})
}

//16.测评 快乐指数测试
func EvaluationKuailes(c *gin.Context) {
	array := [5]string{"亲子关系", "亲密关系", "职场晋升", "刚毕业", "怎么和领导相处"}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": array,
	})
}

//17.测评 情商测试
func EvaluationQingshangs(c *gin.Context) {
	array := [5]string{"亲子关系", "亲密关系", "职场晋升", "刚毕业", "怎么和领导相处"}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": array,
	})
}

//18.我的 首页
func MineIndex(c *gin.Context) {

	var service service.MineIndexService
	if err := c.ShouldBind(&service); err == nil {
		res := service.MineIndex(c) //TestSelect
		c.JSON(200, res)
	} else {
		c.JSON(200, "ErrorResponse")
	}
	//array := [5]string{"亲子关系", "亲密关系", "职场晋升", "刚毕业", "怎么和领导相处"}
	//c.JSON(200, gin.H{
	//	"code": 200,
	//	"msg":  "获取成功",
	//	"data": array,
	//})
}

//19.我的 报告
func MineReport(c *gin.Context) {
	var service service.MineReportService
	if err := c.ShouldBind(&service); err == nil {
		res := service.MineReport(c) //TestSelect
		c.JSON(200, res)
	} else {
		c.JSON(200, "ErrorResponse")
	}
	//array := [5]string{"亲子关系", "亲密关系", "职场晋升", "刚毕业", "怎么和领导相处"}
	//c.JSON(200, gin.H{
	//	"code": 200,
	//	"msg":  "获取成功",
	//	"data": array,
	//})
}

//20.我的 我发布的话术
func MinePublic(c *gin.Context) {
	array := [5]string{"亲子关系", "亲密关系", "职场晋升", "刚毕业", "怎么和领导相处"}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": array,
	})
}

//21.我的 邀请用户
func MineInvitation(c *gin.Context) {
	array := [5]string{"亲子关系", "亲密关系", "职场晋升", "刚毕业", "怎么和领导相处"}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": array,
	})
}

//22.我的 规则说明 积分规则 积分享受的权益
func MineRule(c *gin.Context) {
	array := [5]string{"亲子关系", "亲密关系", "职场晋升", "刚毕业", "怎么和领导相处"}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": array,
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
	var service service.Luntan
	if err := c.ShouldBind(&service); err == nil {
		res := service.LuntanList(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, " ErrorResponse(err)")
	}

	//array := [4]string{"亲子关系", "亲密关系", "职场晋升", "刚毕业"}
	//c.JSON(200, gin.H{
	//	"code": 200,
	//	"data": array,
	//	"msg":  "获取成功",
	//})
}
