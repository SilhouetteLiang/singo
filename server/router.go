package server

import (
	"os"
	"singo/api/apiV1"
	"singo/api/apiV10"
	"singo/api/apiV2"
	"singo/api/apiV3"
	"singo/api/apiV4"
	"singo/api/apiV5"
	"singo/api/apiV6"
	"singo/api/apiV7"
	"singo/api/apiV8"
	"singo/api/apiV9"
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
	// 基础路由
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
	//个人PC端博客介绍项目---网页版展示个人信息生平资质和介绍  网络名片   最后变成了人事招聘网站
	v2 := r.Group("/api/v2")
	{
		v2.POST("ping", apiV2.Pings)
		v2.GET("lottery/algorithm", apiV2.LotteryAlgorithm)
	}
	//心里测试项目--小众人群的社交软件
	v3 := r.Group("/api/v3")
	{
		v3.POST("ping", apiV3.Ping)
		v3.GET("psychologicalTest/index", apiV3.Index)                              //1首页 GET
		v3.POST("psychologicalTest/index/search", apiV3.IndexSearch)                //2首页 POST搜索
		v3.GET("psychologicalTest/index/basicinfo", apiV3.IndexBasicInfo)           //3首页 GET获取对象和场景
		v3.POST("psychologicalTest/index/publish/script", apiV3.IndexPublishScript) //4首页 POST发布话术

		v3.GET("psychologicalTest/forum/list", apiV3.ForumList)                             //1论坛 GET列表 		desc
		v3.POST("psychologicalTest/forum/publish/invitation", apiV3.ForumPublishInvitation) //2论坛 POST发布论坛
		v3.GET("psychologicalTest/forum/comment/list", apiV3.ForumCommentList)              //3论坛 GET评论列表	desc
		v3.POST("psychologicalTest/forum/publish/comment", apiV3.ForumPublishComment)       //4论坛 POST发布评论
		v3.POST("psychologicalTest/forum/zan", apiV3.ForumZan)                              //5论坛 POST点赞

		v3.GET("psychologicalTest/evaluation/index", apiV3.EvaluationIndex)           //1测评 GET首页
		v3.GET("psychologicalTest/evaluation/xingge", apiV3.EvaluationXingge)         //2测评 GET 性格测试
		v3.GET("psychologicalTest/evaluation/MBTI", apiV3.EvaluationMBTI)             //3测评 GET MBTI测试
		v3.GET("psychologicalTest/evaluation/kuaile", apiV3.EvaluationKuaile)         //4测评 GET 快乐指数测试
		v3.GET("psychologicalTest/evaluation/qingshang", apiV3.EvaluationQingshang)   //5测评 GET 情商测试
		v3.POST("psychologicalTest/evaluation/xingge", apiV3.EvaluationXingges)       //6测评 POST 性格测试
		v3.POST("psychologicalTest/evaluation/MBTI", apiV3.EvaluationMBTIs)           //7测评 POST MBTI测试
		v3.POST("psychologicalTest/evaluation/kuaile", apiV3.EvaluationKuailes)       //8测评 POST 快乐指数测试
		v3.POST("psychologicalTest/evaluation/qingshang", apiV3.EvaluationQingshangs) //9测评 POST 情商测试
		v3.GET("psychologicalTest/evaluation/zonghe", apiV3.EvaluationZonghe)         //10测评 GET 综合测试
		v3.POST("psychologicalTest/evaluation/zonghes", apiV3.EvaluationZonghes)      //11测评 POST 综合测试

		v3.GET("psychologicalTest/mine/index", apiV3.MineIndex)           //1我的 GET首页
		v3.GET("psychologicalTest/mine/report", apiV3.MineReport)         //2我的 GET报告
		v3.GET("psychologicalTest/mine/public", apiV3.MinePublic)         //3我的 GET我发布的话术 desc
		v3.GET("psychologicalTest/mine/invitation", apiV3.MineInvitation) //4我的 GET邀请用户
		v3.GET("psychologicalTest/mine/rule", apiV3.MineRule)             //5我的 GET规则说明 积分规则 积分享受的权益
		v3.POST("psychologicalTest/mine/getuserinfo", apiV3.MineUserinfo) //6我的 POST从微信获取我的信息  只有昵称和头像
		v3.POST("psychologicalTest/mine/getUid", apiV3.MineReturnUid)     //7我的 POST返回系统定义的uid为openid
		v3.POST("psychologicalTest/pay/pay", apiV3.Pay)                   //8我的 POST支付
		v3.GET("psychologicalTest/pay/notify", apiV3.Notify)              //9我的 GET支付回调

	}
	//点餐小工具   工具类的
	v4 := r.Group("/api/v4")
	{
		v4.POST("ping", apiV4.Pings)
		v4.GET("lottery/algorithm", apiV4.LotteryAlgorithm)

		v3.GET("psychologicalTest/subject/list", apiV3.SubjectList)      //题目列表接口
		v3.GET("psychologicalTest/user/list", apiV3.UserList)            //用户列表
		v3.POST("psychologicalTest/subject/input", apiV3.InputTitle)     //输入题目
		v3.POST("psychologicalTest/craft/input", apiV3.InputSpeechCraft) //输入话术
		v3.POST("psychologicalTest/userinfo/input", apiV3.InputUserInfo) //获取用户信息
		v3.GET("psychologicalTest/subject/result1", apiV3.Result1)       //结果1
		v3.GET("psychologicalTest/subject/result2", apiV3.Result2)       //结果2
		v3.GET("psychologicalTest/subject/result3", apiV3.Result3)       //结果3
		v3.POST("psychologicalTest/luntan/add", apiV3.LuntanAdd)         //1.新增论坛
		v3.GET("psychologicalTest/luntan/list", apiV3.LuntanList)        //2.论坛列表
		//首页
		v3.POST("psychologicalTest/subject/list", apiV3.SubjectList) //2.1性格色彩测试接口
	}
	//社交类的----小众人群的社交软件
	v5 := r.Group("/api/v5")
	{
		v5.POST("ping", apiV5.Pings)
		v5.GET("lottery/algorithm", apiV5.LotteryAlgorithm)
	}
	//内容创业  问答社区---验证身份证、学位证、专业   高质量的问答社区
	v6 := r.Group("/api/v6")
	{
		v6.POST("ping", apiV6.Pings)
		v6.GET("lottery/algorithm", apiV6.LotteryAlgorithm)
	}
	//你关注的才是头条、内容分发平台
	v7 := r.Group("/api/v7")
	{
		v7.POST("ping", apiV7.Pings)
		v7.GET("lottery/algorithm", apiV7.LotteryAlgorithm)
	}
	//做平台  平台化(律师平台、)
	v8 := r.Group("/api/v8")
	{
		v8.POST("ping", apiV8.Pings)
		v8.GET("lottery/algorithm", apiV8.LotteryAlgorithm)
	}
	// 人工智能  抓取分析数据
	v9 := r.Group("/api/v9")
	{
		v9.POST("ping", apiV9.Pings)
		v9.GET("lottery/algorithm", apiV9.LotteryAlgorithm)
	}
	// 国外的支付工具 内容付费
	v10 := r.Group("/api/v10")
	{
		v10.POST("ping", apiV10.Pings)
		v10.GET("lottery/algorithm", apiV10.LotteryAlgorithm)
	}

	return r
}
