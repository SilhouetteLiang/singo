package server

import (
	"os"
	"singo/api/apiV1"
	"singo/api/apiV2"
	"singo/api/apiV3"
	"singo/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())



	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", apiV1.Ping)

		// 用户登录
		v1.POST("user/register", apiV1.UserRegister)

		// 用户登录
		v1.POST("user/login", apiV1.UserLogin)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("user/me", apiV1.UserMe)
			auth.DELETE("user/logout", apiV1.UserLogout)
			auth.POST("user/input", apiV1.UserInput)
		}

	}

	//彩票算法接口
	v2 := r.Group("/api/v2")
	{
		v2.POST("ping", apiV2.Pings)
		v2.GET("lottery/algorithm", apiV2.LotteryAlgorithm)

	}

	//心里测试小工具
	v3 := r.Group("/api/v3")
	{
		v3.POST("ping", apiV3.Pings)
		v3.GET("psychologicalTest/subject/list", apiV3.SubjectList)      //题目列表接口
		v3.GET("psychologicalTest/user/list", apiV3.UserList)            //用户列表
		v3.POST("psychologicalTest/subject/input", apiV3.InputTitle)     //输入题目
		v3.POST("psychologicalTest/craft/input", apiV3.InputSpeechCraft) //输入话术
		v3.POST("psychologicalTest/userinfo/input", apiV3.InputUserInfo) //获取用户信息
		v3.GET("psychologicalTest/subject/result1", apiV3.Result1)       //结果1
		v3.GET("psychologicalTest/subject/result2", apiV3.Result2)       //结果2
		v3.GET("psychologicalTest/subject/result3", apiV3.Result3)       //结果3
		v3.GET("psychologicalTest/index", apiV3.Index)                   //首页

		//v3.GET("psychologicalTest/subject/list", apiV3.SubjectList) //2、题目结果返回接口
		//v3.GET("psychologicalTest/subject/list", apiV3.SubjectList) //2.1性格色彩测试接口
		//v3.GET("psychologicalTest/subject/list", apiV3.SubjectList) //2.2荣格测试接口
		//v3.GET("psychologicalTest/subject/list", apiV3.SubjectList) //2.3 蜘蛛图接口
		//v3.GET("psychologicalTest/subject/list", apiV3.SubjectList) //3、论坛发帖接口

	}

	return r
}
