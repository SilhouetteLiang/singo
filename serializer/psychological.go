package serializer

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"singo/model"
)

// Luntan 内容序列化器
type LunTan struct {
	ID        uint   `json:"id"`
	Content   string `form:"content" json:"content"`
	Img       string `form:"img" json:"img"`
	UserId    string `form:"userid" json:"userid"`
	Nickname  string `form:"nickname" json:"nickname"`
	Status    int64  `form:"status" json:"status"`
	CreatedAt int64
}

// Luntan 内容序列化器
type LunTans struct {
	//ID         uint   `json:"id"`
	//Content    string `gorm:"not null;default:内容;comment:内容"`
	//Img        string `gorm:"not null;default:Img;comment:论坛图片地址"`
	//UserId     string `gorm:"int(10);not null;default:100001;comment:用户ID"`
	//Nickname   string `gorm:"size:255;type:char(255);not null;default:张三;comment:用户昵称"`
	//Status     int64  `gorm:"type:int(1);not null;default:0;comment:状态值 1正常 2异常"`
	//Sort       int64  `gorm:"type:int(1);not null;default:2;comment:是否置顶 1 是 2否"`
	Zan       int64 `gorm:"type:int(11);not null;default:1;comment:点赞数量"`
	IsDisplay int64
	//CommentNum int64  `gorm:"type:int(11);default:0;comment:评论数量"`
}

//Name    string `gorm:"size:255;type:char(255);not null;default:0;comment:测试题目名称"` // 设置字段大小为255
//TestNum int64  `gorm:"type:int(11);not null;default:0;comment:已使用过的测试人数 "`        // 设置字段大小为255
//Price   int64  `gorm:"type:int(11);not null;default:0;comment:测试价格"`              // 设置字段大小为255
// Luntan 内容序列化器
type TestInfos struct {
	Name    string `form:"name" json:"name"`
	TestNum int64  `form:"testnum" json:"testnum"`
	Price   int64  `form:"price" json:"price"`
}

//Title      string     `gorm:"size:255;type:char(255);not null;default:标题;comment:标题"` // 设置字段大小为255
//Content    string     `gorm:"not null;default:内容;comment:内容"`
//Img        string     `gorm:"not null;default:Img;comment:论坛图片地址"`
//UserId     string     `gorm:"int(10);not null;default:100001;comment:用户ID"`
//Nickname   string     `gorm:"size:255;type:char(255);not null;default:张三;comment:用户昵称"`
//Status     int64      `gorm:"type:int(1);not null;default:0;comment:状态值 1正常 2异常"`
//Sort       int64      `gorm:"type:int(1);not null;default:2;comment:是否置顶 1 是 2否"`
//Zan        int64      `gorm:"type:int(11);not null;default:1;comment:点赞数量"`
//CommentNum int64      `gorm:"type:int(11);default:0;comment:评论数量"`

// Psychological 内容序列化器
type Psychological struct {
	ID        uint   `json:"id"`
	Status    int64  `json:"status"`
	Title     string `json:"title"`
	CreatedAt int64  `json:"created_at"`
	A         string
	B         string
	C         string
	D         string
	Type      int64 `json:"type"`
}

// Craft 内容序列化器
type Craft struct {
	Object  string `form:"object" json:"object"`
	Scene   string `form:"scene" json:"scene"`
	Content string `form:"content" json:"content"`
	Source  string `form:"source" json:"source"`
	Status  int64  `form:"status" json:"status"`
}

//23我的 从微信获取我的信息
type UserInfo struct {
	UserName       string
	Nickname       string
	Avatar         string
	PasswordDigest string
	Tell           string
	DrawTag        string
	Image          string
	AccessNum      int64
	Status         string
	UserPoints     int64
	OpenId         string
	UserCode       string

	//User           string `gorm:"size:255;type:char(255);not null;default:0;comment:用户"`         // 设置字段大小为255
	//UserName       string `gorm:"size:255;type:char(255);not null;default:0;comment:用户真实名字"`     // 设置字段大小为255
	//Nickname       string `gorm:"size:255;type:char(255);not null;default:0;comment:用户昵称"`       // 设置字段大小为255
	//Avatar         string `gorm:"size:1000;not null;default:0;comment:用户头像"`                     // 设置字段大小为255
	//PasswordDigest string `gorm:"size:255;type:char(255);not null;default:0;comment:用户密码"`       // 设置字段大小为255
	//Tell           string `gorm:"size:255;type:char(255);not null;default:999999;comment:用户手机号"` // 设置字段大小为255
	//DrawTag        string `gorm:"size:255;type:char(255);not null;default:0;comment:用户标签"`       // 设置字段大小为255
	//Image          string `gorm:"size:255;type:char(255);not null;default:0;comment:用户补充图片"`     // 设置字段大小为255
	//AccessNum      int64  `gorm:"type:int(10);not null;default:1;comment:用户访问次数"`                // 设置字段大小为255
	//Status         string `gorm:"type:char(255);not null;default:0;comment:状态值 1正常 2异常"`
	//UserPoints     int64  `gorm:"type:int(10);not null;default:0;comment:用户积分"`
	//OpenId         string `gorm:"size:255;type:char(255);not null;default:0;comment:用户OPENID"`
	//UserCode       string `gorm:"size:255;type:char(255);not null;default:0;comment:用户code"`
	//Appid          string `gorm:"size:255;type:char(255);not null;default:0;comment:小程序AppId"`
	//AppSecret      string `gorm:"size:255;type:char(255);not null;default:0;comment:小程序AppSecret"`
}

type UserInfoRes struct {
	Grade string
}

//24我的 返回系统定义的uid为openid
type ReturnUid struct {
	OpenId string `form:"openId" json:"openId"`
}

// Users 内容序列化器
type LuntanComments struct {
	//LuntanId   int64      `gorm:"int(11);not null;default:99999999;comment:帖子ID"`
	//Content    string     `gorm:"not null;default:内容;comment:评论内容"`
	//UserId     int64      `gorm:"int(10);not null;default:100001;comment:用户ID"`
	//Nickname   string     `gorm:"size:255;type:char(255);not null;default:张三;comment:用户昵称"`
	//Status     int64      `gorm:"type:int(1);not null;default:0;comment:状态值 1正常 2异常"`
	LuntanId int64 `gorm:"int(11);not null;default:99999999;comment:帖子ID"`
	Content  string
	UserId   int64
	Nickname string
	Status   int64
}

// Users 内容序列化器
type KeyWordScene struct {
	KeyWord  string `gorm:"size:255;type:char(255);not null;default:0;comment:关键词"`
	Scene    string `gorm:"size:255;type:char(255);not null;default:0;comment:场景"`
	Status   int64  `gorm:"type:int(1);not null;default:0;comment:状态值 1正常 2异常"`
	BeiYong1 string `gorm:"size:255;type:char(255);not null;default:0;comment:BeiYong1"`
	BeiYong2 string `gorm:"size:255;type:char(255);not null;default:0;comment:BeiYong2"` // 设置字段大小为255
	BeiYong3 string `gorm:"size:255;type:char(255);not null;default:0;comment:BeiYong3"` // 设置字段大小为255
	BeiYong4 string `gorm:"size:255;type:char(255);not null;default:0;comment:BeiYong4"` // 设置字段大小为255
	BeiYong5 string `gorm:"size:255;type:char(255);not null;default:0;comment:BeiYong5"` // 设置字段大小为255
}

type ResPsychological struct {
	Items []Item `json:"items"`
}

type Item struct {
	Title  string `form:"title" json:"title" binding:"required,min=5,max=30"`
	Status uint64 `form:"status" json:"status" binding:"required"`
}

type XinggeResult struct {
	Title   string `form:"title" json:"title" binding:"required,min=5,max=30"`
	Answers []*Lists
	Status  uint64 `form:"status" json:"status" binding:"required"`
}

type Lists struct {
	Answerno string `form:"answerno" json:"answerno" `
	Content  string `form:"content" json:"content" `
}

// Users 内容序列化器
type Users struct {
	UserName       string
	PasswordDigest string
	Nickname       string
	Avatar         string `gorm:"size:1000"`
	User           string `gorm:"size:255;type:char(255)"` // 设置字段大小为255
	Tell           string `gorm:"size:255;type:char(12)"`
	DrawTag        string `gorm:"size:255;type:char(255)"`
	Image          string `gorm:"size:255;type:char(255)"`
	AccessNum      int64
	Status         string
}

//序列化响应 新增用户
//func BuildIndexSearchResponse(KeyWordScene model.KeyWordScene) Response {
//	return Response{
//		Data: BuildIndexSearch(user),
//	}
//}
//
////序列化内容 新增用户
//func BuildIndexSearch(KeyWordScene model.KeyWordScene) Users {
//	return Users{
//		Nickname:  KeyWordScene.Nickname,
//		Avatar:    KeyWordScene.Avatar,
//		Tell:      KeyWordScene.Tell,
//		AccessNum: KeyWordScene.AccessNum,
//		DrawTag:   KeyWordScene.DrawTag,
//	}
//}

//序列化响应 2 搜索
func BuildUsersResponse(user model.User) Response {
	return Response{
		Data: BuildUsers(user),
	}
}

//序列化内容 2 搜索
func BuildUsers(user model.User) Users {
	return Users{
		Nickname:  user.Nickname,
		Avatar:    user.Avatar,
		Tell:      user.Tell,
		AccessNum: user.AccessNum,
		DrawTag:   user.DrawTag,
	}
}

//序列化响应 输入话术
func BuildCraftResponse(psychological model.Craft) Response {
	return Response{
		Data: BuildCraft(psychological),
	}
}

//序列化内容 输入话术
func BuildCraft(psychological model.Craft) Craft {
	return Craft{
		Object:  psychological.Object,
		Scene:   psychological.Scene,
		Content: psychological.Content,
		Source:  psychological.Source,
		Status:  psychological.Status,
	}
}

//序列化响应 新增论坛
func BuildLuntanResponse(luntan model.Luntan) Response {
	return Response{
		Data: BuildLuntan(luntan),
	}
}

//序列化内容 新增论坛
func BuildLuntan(Luntan model.Luntan) LunTan {
	return LunTan{
		ID:        Luntan.ID,
		Content:   Luntan.Content,
		Img:       Luntan.Img,
		UserId:    Luntan.UserId,
		Nickname:  Luntan.Nickname,
		Status:    Luntan.Status,
		CreatedAt: Luntan.CreatedAt.Unix(),
	}
}

//序列化响应 输入题目
func BuildPsychologicalResponse(psychological model.Psychological) Response {
	return Response{
		Data: BuildpSychologicals(psychological),
	}
}

//序列化内容 输入题目
func BuildpSychologicals(psychological model.Psychological) Psychological {
	return Psychological{
		ID:        psychological.ID,
		Title:     psychological.Title,
		Type:      psychological.Type,
		A:         psychological.A,
		B:         psychological.B,
		C:         psychological.C,
		D:         psychological.D,
		Status:    psychological.Status,
		CreatedAt: psychological.CreatedAt.Unix(),
	}
}

//序列化心理测试响应
func BuildPsychologicalResponses(psychological []model.Psychological) Response {
	return Response{
		//Data: BuildpSychological(psychological),
		Data: psychological,
	}
}

//  序列化内容
func BuildpSychological(psychological []model.Psychological) model.Psychological {
	return model.Psychological{}
}

type Response1 struct {
	Code  int         `json:"code"`
	Count int64       `json:"count"`
	Data  interface{} `json:"data,omitempty"`
	Msg   string      `json:"msg"`
	Error string      `json:"error,omitempty"`
}

func ParamErr1(msg string, err error) Response1 {
	if msg == "" {
		msg = "参数错误"
	}
	return Err1(CodeParamErr, msg, err)
}

// Err 通用错误处理
func Err1(errCode int, msg string, err error) Response1 {
	res := Response1{
		Code: errCode,
		Msg:  msg,
	}
	// 生产环境隐藏底层报错
	if err != nil && gin.Mode() != gin.ReleaseMode {
		res.Error = err.Error()
	}
	return res
}

//序列化 论坛 响应
func BuildLuntanResponses(Luntan []model.Luntan, count int64) Response1 {
	return Response1{
		//Data: BuildpSychological(psychological),
		Data:  Luntan,
		Count: count,
	}
}

//1. 首页
type SceneService struct {
	Scene  string `gorm:"size:255;type:char(255);not null;default:0;comment:场景"`
	Status int64  `gorm:"type:int(1);not null;default:0;comment:状态值 1正常 2异常"`
}

func BuildKeyWordSceneResponse(KeyWordScene []model.Scene) Response {
	return Response{
		Data: KeyWordScene,
	}
}

//2. 首页搜索
type Crafts struct {
	Content string `gorm:"not null;default:0;comment:内容"`
}

func BuildIndexSearchResponse(Craft []model.Crafts) Response {
	return Response{
		Data: Craft,
	}
}

//7. 论坛 评论列表
func BuildForumCommentListResponse(Craft []model.LuntanComment, count int64) Response {
	return Response{
		Data:  Craft,
		Count: count,
	}
}

//序列化响应 8 论坛发布评论
func BuildPublishCommentResponse(LuntanComment model.LuntanComment) Response {
	return Response{
		Data: BuildPublishComment(LuntanComment),
	}
}

//序列化内容 8 论坛发布评论
func BuildPublishComment(LuntanComment model.LuntanComment) LuntanComments {
	return LuntanComments{
		LuntanId: LuntanComment.LuntanId,
		Content:  LuntanComment.Content,
		Nickname: LuntanComment.Nickname,
		Status:   LuntanComment.Status,
	}
}

//序列化响应 8 论坛 点赞
func BuildForumZanResponse(Luntan model.LuntanZan, isDisplay int64) Response {
	return Response{
		Data: BuildForumZan(Luntan, isDisplay),
	}
}

//序列化内容 8 论坛 点赞
func BuildForumZan(Luntan model.LuntanZan, isDisplay int64) LunTans {
	return LunTans{
		//LuntanId   int64      `gorm:"int(11);not null;default:99999999;comment:帖子ID"`
		//Content    string     `gorm:"not null;default:内容;comment:评论内容"`
		//UserId     int64      `gorm:"int(10);not null;default:100001;comment:用户ID"`
		//Nickname   string     `gorm:"size:255;type:char(255);not null;default:张三;comment:用户昵称"`
		//Status     int64      `gorm:"type:int(1);not null;default:0;comment:状态值 1正常 2异常"`
		Zan:       Luntan.Zan,
		IsDisplay: isDisplay,
	}
}

//序列化响应 9.测评 首页
func BuildEvaluationIndexResponse(TestInfo []model.TestInfo) Response {
	return Response{
		Data: TestInfo,
	}
}

//序列化内容 9 测评 首页
//func BuildEvaluationIndex(TestInfo []model.TestInfo) TestInfos {
//	return TestInfos{
//		//Name    string `gorm:"size:255;type:char(255);not null;default:0;comment:测试题目名称"` // 设置字段大小为255
//		//TestNum int64  `gorm:"type:int(11);not null;default:0;comment:已使用过的测试人数 "`        // 设置字段大小为255
//		//Price   int64  `gorm:"type:int(11);not null;default:0;comment:测试价格"`              // 设置字段大小为255
//		Name:    TestInfo.Name,
//		TestNum: TestInfo.TestNum,
//		Price:   TestInfo.Price,
//	}
//}

//序列化响应 10 测评 get 性格测试
func BuildEvaluationXinggeResponse(Psychological []model.Psychological) Response {
	var dat = []*XinggeResult{}
	var arr = []*Lists{}

	var A map[int]string
	A = make(map[int]string)
	for k, item := range Psychological {
		A[k] = item.A
	}
	var B map[int]string
	B = make(map[int]string)
	for k, item := range Psychological {
		B[k] = item.B
	}
	var C map[int]string
	C = make(map[int]string)
	for k, item := range Psychological {
		C[k] = item.C
	}
	var D map[int]string
	D = make(map[int]string)
	for k, item := range Psychological {
		D[k] = item.D
	}

	for k, item := range Psychological {

		dat = append(dat, &XinggeResult{
			Title: item.Title,
			Answers: []*Lists{
				{Answerno: "A", Content: A[k]},
				{Answerno: "B", Content: B[k]},
				{Answerno: "C", Content: C[k]},
				{Answerno: "D", Content: D[k]},
			},
		})
	}

	fmt.Printf("res == %v /n", arr)
	return Response{
		Data: dat,
		//Data: A,
		//Data: arr,
	}
}

//14测评 post 性格测试
func BuildResXinggesResponse(User model.UserReport) Response {
	return Response{
		Data: User,
	}
}

//, apiV3.EvaluationXingges)             //14测评 post 性格测试
//v3.POST("psychologicalTest/evaluation/MBTI", apiV3.EvaluationMBTIs)                 //15测评 post MBTI测试
//v3.POST("psychologicalTest/evaluation/kuaile", apiV3.EvaluationKuailes)             //16测评 post 快乐指数测试
//v3.POST("psychologicalTest/evaluation/qingshang", apiV3.EvaluationQingshangs)       //17测评 post 情商测试

//15测评 post MBTI测试
func BuildResMBTIsResponse(User model.UserReport) Response {
	return Response{
		Data: User,
	}
}

//16测评 post 快乐指数测试
func BuildResKuailesResponse(User model.UserReport) Response {
	return Response{
		Data: User,
	}
}

//17测评 post 情商测试
func BuildResQingshangsResponse(User model.UserReport) Response {
	return Response{
		Data: User,
	}
}

//18.我的 首页
func BuildMineIndexResponse(User model.UserMine) Response {
	return Response{
		Data: User,
	}
}

//19.我的 报告
func BuildMineReportResponse(UserReportList []model.UserReportList) Response {
	return Response{
		Data: UserReportList,
	}
}

//20我的 GET我发布的话术 desc
func BuildMinePublicResponse(MyCraft []model.MyCraft) Response {
	return Response{
		Data: MyCraft,
	}
}

//2121我的 GET邀请用户 desc
func BuildMineInvitationResponse(UserInvitationInfo []model.UserInvitationInfo) Response {
	return Response{
		Data: UserInvitationInfo,
	}
}

//23我的 从微信获取我的信息
func BuildMineUserinfoResponse(userInfo model.User) Response {
	return Response{
		Data: BuildMineUserinfo(userInfo),
		//Data: 3,
	}
}

//23我的 从微信获取我的信息
func BuildMineUserinfo(userInfo model.User) UserInfo {
	return UserInfo{
		UserName:       userInfo.UserName,
		Nickname:       userInfo.Nickname,
		Avatar:         userInfo.Avatar,
		PasswordDigest: userInfo.PasswordDigest,
		Tell:           userInfo.Tell,
		DrawTag:        userInfo.DrawTag,
		Image:          userInfo.Image,
		AccessNum:      userInfo.AccessNum,
		Status:         userInfo.Status,
		UserPoints:     userInfo.UserPoints,
		OpenId:         userInfo.OpenId,
		UserCode:       userInfo.UserCode,
	}
}

//24我的 返回系统定义的uid为openid
func BuildMineReturnUidResponse(userInfo model.User) Response {
	return Response{
		Data: BuildMineReturnUid(userInfo),
	}
}

//24我的 返回系统定义的uid为openid
func BuildMineReturnUid(userInfo model.User) ReturnUid {
	return ReturnUid{
		OpenId: userInfo.OpenId,
	}
}
