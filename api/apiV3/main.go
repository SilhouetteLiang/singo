package apiV3

import (
	"fmt"
	"github.com/gin-gonic/gin"
	dev "github.com/pyihe/wechat-sdk"
	"github.com/w3liu/go-common/constant/timeformat"
	"os"
	"singo/serializer"
	"singo/service"
	"strconv"
	"sync/atomic"
	"time"
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
	pageId := c.Request.URL.Query().Get("pageId")
	id, _ := strconv.Atoi(pageId)
	//id, _ := strconv.ParseInt(pageId, 10, 64)
	var service service.Luntan
	if err := c.ShouldBind(&service); err == nil {
		res := service.LuntanList(c, id)
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
	pageId := c.Request.URL.Query().Get("pageId")
	pageIds, _ := strconv.Atoi(pageId)
	id, _ := strconv.ParseInt(commentId, 10, 64)
	//name := c.Param("name")
	//c.String(http.StatusOK, "Hello %s", id)
	var service service.IndexSearchService
	if err := c.ShouldBind(&service); err == nil {
		fmt.Printf("2  : %v \n", service)
		res := service.ForumCommentList(c, id, pageIds) //TestSelect
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
	var service service.EvaluationIndexService
	if err := c.ShouldBind(&service); err == nil {
		fmt.Printf("service  : %v \n", service)
		res := service.EvaluationIndex(c)
		c.JSON(200, res)
	} else {
		fmt.Printf("err  : %v \n", err)

		c.JSON(200, "ErrorResponse")
	}
	//array := [5]string{"性格色彩", "MBTI测试", "快乐指数", "情商测试"}
	//c.JSON(200, gin.H{
	//	"code": 200,
	//	"msg":  "获取成功",
	//	"data": array,
	//})
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
	var service service.XinggesService
	if err := c.ShouldBind(&service); err == nil {
		res := service.ResXingges(c) //TestSelect
		c.JSON(200, res)
	} else {
		c.JSON(200, "ErrorResponse")
	}
}

//15.测评 MBTI测试
func EvaluationMBTIs(c *gin.Context) {
	var service service.MBTIService
	if err := c.ShouldBind(&service); err == nil {
		res := service.ResMBTIs(c) //TestSelect
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

//16.测评 快乐指数测试
func EvaluationKuailes(c *gin.Context) {
	var service service.KuaileService
	if err := c.ShouldBind(&service); err == nil {
		res := service.ResKuailes(c) //TestSelect
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

//17.测评 情商测试
func EvaluationQingshangs(c *gin.Context) {
	var service service.QingShangService
	if err := c.ShouldBind(&service); err == nil {
		res := service.ResQingshangs(c) //TestSelect
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

//23我的 从微信获取我的信息
func MineUserinfo(c *gin.Context) {
	var service service.MineUserinfoService
	if err := c.ShouldBind(&service); err == nil {
		fmt.Printf("service  : %v \n", service)
		res := service.MineUserinfo(c)
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

//24我的 返回系统定义的uid为openid
func MineReturnUid(c *gin.Context) {
	var service service.MineReturnUidService
	if err := c.ShouldBind(&service); err == nil {
		fmt.Printf("service  : %v \n", service)
		res := service.MineReturnUid(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, "ErrorResponse")
	}
}

//25我的 支付
func Pay(c *gin.Context) {
	var services service.PayServices
	if err := c.ShouldBind(&services); err == nil {
		fmt.Printf("services  : %v \n", services)
	} else {
		fmt.Printf("err  : %v \n", err)
		c.JSON(200, "ErrorResponse")
	}
	var service service.PayService
	var appId, mchId, apiKey, apiSecret string
	appId = "wx6902b88cb7e7ea61"
	mchId = "1513195931"
	apiKey = "sdewCSEwefafnk798moaklfja23rerwc"
	apiSecret = "365bb1ab8d91e0eb9dcbd3a8d9ef8b12"

	client := dev.NewPayer(dev.WithAppId(appId), dev.WithMchId(mchId), dev.WithApiKey(apiKey), dev.WithSecret(apiSecret))

	var orderId string
	t2 := time.Now()
	//yy.ext = time.Now().Format("2006-01-02 15:04:05")
	orderId = Generate(t2)
	fmt.Printf("orderId %v /n", orderId)

	// unified order（统一下单）
	param := dev.NewParam()
	param.Add("nonce_str", "jgOUy(#oiDhbLMkTU")
	param.Add("body", "yourBody")
	param.Add("out_trade_no", orderId)
	param.Add("total_fee", services.Amount)
	//map[string]interface{}{"total": amount * 100, "currency": "CNY"}
	param.Add("spbill_create_ip", "39.105.90.51")
	param.Add("notify_url", "https://www.qxxa.top/api/v3/psychologicalTest/pay/notify")
	param.Add("trade_type", "JSAPI")
	param.Add("openid", services.Openid)
	result, err := client.UnifiedOrder(param)
	if err != nil {
		handleErr(err)
		c.JSON(200, "ErrorResponse")
	}
	fmt.Printf("result %v /n", result)

	//appId, _ = result.GetString("apppid")
	//prepayId, _ := result.GetString("prepay_id")
	//param = dev.NewParam()
	//param.Add("appId", appId)
	//param.Add("timeStamp", time.Now().Unix())
	//param.Add("nonceStr", "yourNonceStr1")
	//param.Add("package", "prepay_id="+prepayId)
	//param.Add("signType", "MD5")
	timeUnix := strconv.FormatInt(time.Now().Unix(), 10)
	//use to evoke wechat pay
	param.Add("appId", appId)
	param.Add("nonceStr", result.Data()["nonce_str"])
	//param.Add("prepay_id", result.Data()["prepay_id"])
	param.Add("package", `prepay_id=`+`result.Data()["prepay_id"]`)
	param.Add("timeStamp", timeUnix)
	param.Add("key", apiKey)
	sign := param.Sign("MD5")
	//param.Add("sign", sign)

	//result1, err := client.UnifiedOrder(param)
	//if err != nil {
	//	handleErr(err)
	//}
	//result map[appid:wx6902b88cb7e7ea61 mch_id:1513195931 nonce_str:DNRDy7I29yq0ioT9 prepay_id:wx21124652538668b2579dbeea9aae2b0000 result_code:SUCCESS return_code:SUCCESS return_msg:OK sign:9578AB835F2E7A04B689AB34937E5282 trade_type:JSAPI]
	appid := result.Data()["mch_id"]
	fmt.Printf("appid %v /n", appid)

	service.TimeStamp = timeUnix
	service.NonceStr = result.Data()["nonce_str"]
	service.Package = result.Data()["prepay_id"]
	service.SignType = "MD5"
	service.PaySign = sign

	for _, item := range result.Data() {
		fmt.Printf("item %v /n", item)

	}

	fmt.Printf("sign %v /n", sign)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": service,
	})
	//fmt.Printf("result1 %v /n", result1)

	// download bill
	//param = dev.NewParam()
	//param.Add("nonce_str", "yourNonceStr")
	//param.Add("bill_date", "yourDate")
	//param.Add("bill_type", "ALL")
	//param.Add("tar_type", "GZIP")
	//err = client.DownloadBill(param, "./bill")
	//if err != nil {
	//	handleErr(err)
	//}

	// get phone for mini program user
	//result, err = client.GetUserPhoneForMini("code", "encryptedData", "iv")
	//if err != nil {
	//	handleErr(err)
	//}
	//var phone string
	//if countryCode := result.Get("countryCode"); countryCode != nil && countryCode.(string) == "86" {
	//	purePhone := result.Get("purePhoneNumber")
	//	phone = purePhone.(string)
	//} else {
	//	phoneNumber := result.Get("phoneNumber")
	//	phone = phoneNumber.(string)
	//}
	//fmt.Printf("user phone is %s\n", phone)
	//var service service.MineReturnUidService
	//if err := c.ShouldBind(&service); err == nil {
	//	fmt.Printf("service  : %v \n", service)
	//	res := service.Pay(c)
	//	c.JSON(200, res)
	//} else {
	//	c.JSON(200, "ErrorResponse")
	//}
}

//生成单号
//06123xxxxx
//sum 最少10位,sum 表示全部单号位数
//func MakeYearDaysRand(sum int) string {
//	//年
//	strs := time.Now().Format("06")
//	//一年中的第几天
//	days := strconv.Itoa(GetDaysInYearByThisYear())
//	count := len(days)
//	if count < 3 {
//		//重复字符0
//		days = strings.Repeat("0", 3-count) + days
//	}
//	//组合
//	strs += days
//	//剩余随机数
//	sum = sum - 5
//	if sum < 1 {
//		sum = 5
//	}
//	//0~9999999的随机数
//	ran := GetRand()
//	pow := math.Pow(10, float64(sum)) - 1
//	//fmt.Println("sum=>", sum)
//	//fmt.Println("pow=>", pow)
//	result := strconv.Itoa(ran.Intn(int(pow)))
//	count = len(result)
//	//fmt.Println("result=>", result)
//	if count < sum {
//		//重复字符0
//		result = strings.Repeat("0", sum-count) + result
//	}
//	//组合
//	strs += result
//	return strs
//}
//
////年中的第几天
//func GetDaysInYearByThisYear() int {
//	now := time.Now()
//	total := 0
//	arr := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
//	y, month, d := now.Date()
//	m := int(month)
//	for i := 0; i < m-1; i++ {
//		total = total + arr[i]
//	}
//	if (y%400 == 0 || (y%4 == 0 && y%100 != 0)) && m > 2 {
//		total = total + d + 1
//	} else {
//		total = total + d
//	}
//	return total
//}

var num int64

//生成24位订单号
//前面17位代表时间精确到毫秒，中间3位代表进程id，最后4位代表序号
//1.固定24位长度订单号，毫秒+进程id+序号。
//2.同一毫秒内只要不超过一万次并发，则订单号不会重复。
func Generate(t time.Time) string {
	s := t.Format(timeformat.Continuity)
	m := t.UnixNano()/1e6 - t.UnixNano()/1e9*1e3
	ms := sup(m, 3)
	p := os.Getpid() % 1000
	ps := sup(int64(p), 3)
	i := atomic.AddInt64(&num, 1)
	r := i % 10000
	rs := sup(r, 4)
	n := fmt.Sprintf("%s%s%s%s", s, ms, ps, rs)
	return n

}

//对长度不足n的数字前面补0
func sup(i int64, n int) string {
	m := fmt.Sprintf("%d", i)
	for len(m) < n {
		m = fmt.Sprintf("0%s", m)
	}
	return m
}

func Notify(c *gin.Context) {

	fmt.Println("微信支付回调成功")
	var resXml string

	//var appId, mchId, apiKey, apiSecret string
	//
	//client := dev.NewPayer(dev.WithAppId(appId), dev.WithMchId(mchId), dev.WithApiKey(apiKey), dev.WithSecret(apiSecret))

	// handle refund notify
	//http.HandleFunc("/refund_notify", func(writer http.ResponseWriter, request *http.Request) {
	//	defer request.Body.Close()
	//	result, err := client.RefundNotify(request.Body)
	//	if err != nil {
	//		handleErr(err)
	//	}
	//	fmt.Printf("RefundNotify Result = %v\n", result.Data())
	//})

	fmt.Println("dsadawdwdadwd")

	err := "dsadawdasdwd"
	fmt.Printf("err %v /n", err)
	//res := "11111111111111111"
	//c.JSON(http.StatusOK, &wechat.V3NotifyRsp{Code: gopay.SUCCESS, Message: "成功"})
	resXml = "<xml>" + "<return_code><![CDATA[FAIL]]></return_code>" + "<return_msg><![CDATA[FAIL]]></return_msg>" + "</xml> "
	resXml = "<xml>" + "<return_code><![CDATA[SUCCESS]]></return_code>" + "<return_msg><![CDATA[OK]]></return_msg>" + "</xml> "
	c.JSON(200, gin.H{
		"code": 200,
		"data": resXml,
		"msg":  "获取成功",
	})

}
func handleErr(err error) {
	fmt.Printf("err %v /n", err)

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
	//var service service.Luntan
	//if err := c.ShouldBind(&service); err == nil {
	//	res := service.LuntanList(c)
	//	c.JSON(200, res)
	//} else {
	//	c.JSON(200, " ErrorResponse(err)")
	//}

	//array := [4]string{"亲子关系", "亲密关系", "职场晋升", "刚毕业"}
	//c.JSON(200, gin.H{
	//	"code": 200,
	//	"data": array,
	//	"msg":  "获取成功",
	//})
}
