package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/thedevsaddam/gojsonq/v2"
	"io/ioutil"
	"net/http"
	"singo/model"
	"singo/serializer"
)

//新增论坛
type Luntan struct {
	Content  string `form:"content" json:"content"`
	Img      string `form:"img" json:"img"`
	UserId   string `form:"userid" json:"userid"`
	Nickname string `form:"nickname" json:"nickname"`
	Status   int64  `form:"status" json:"status"`
	OpenId   string `form:"openid" json:"openid"`
}

// 新增题目
type PsychologicalService struct {
	A      string
	B      string
	C      string
	D      string
	Title  string `form:"title" json:"title"`
	Status int64  `form:"status" json:"status"`
	Type   int64  `form:"type" json:"type"`
}

//Title      string     `gorm:"size:255;type:char(255);not null;default:0;comment:标题"` // 设置字段大小为255
//A          string     `gorm:"not null;default:0;comment:选项A"`
//B          string     `gorm:"not null;default:0;comment:选项B"`
//C          string     `gorm:"not null;default:0;comment:选项C"`
//D          string     `gorm:"not null;default:0;comment:选项D"`
//Type       int64      `gorm:"type:int(1);not null;default:0;comment:类型 "`
//Status     int64      `gorm:"type:int(1);not null;default:0;comment:状态值 1正常 2异常"`
//BeiYong1   string     `gorm:"size:255;type:char(255);not null;default:0;comment:备用1"` // 设置字段大小为255
//BeiYong2   string     `gorm:"size:255;type:char(255);not null;default:0;comment:备用2"` // 设置字段大小为255
//BeiYong3   string     `gorm:"size:255;type:char(255);not null;default:0;comment:备用3"` // 设置字段大小为255
//BeiYong4   string     `gorm:"size:255;type:char(255);not null;default:0;comment:备用4"` // 设置字段大小为255
//BeiYong5   string     `gorm:"size:255;type:char(255);not null;default:0;comment:备用5"` // 设置字段大小为255

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
	Object  string `form:"object" json:"object"`
	Scene   string `form:"scene" json:"scene"`
	Content string `form:"content" json:"content"`
	Source  string `form:"source" json:"source"`
	Status  int64  `form:"status" json:"status"`
	OpenId  string `form:"openid" json:"openid"`
}

type Item struct {
	Title  string
	Status uint64
}

// 1.首页
type IndexService struct {
	Status int64 `form:"status" json:"status"`
}

// 2.搜索
type IndexSearchService struct {
	Keyword string `form:"keyword" json:"keyword"`
	Status  int64  `form:"status" json:"status"`
	OpenId  string `form:"openid" json:"openid"`
}

//8.发布评论
type ForumPublishCommentService struct {
	LuntanId int64  `form:"luntanId" json:"luntanId"`
	Content  string `form:"content" json:"content"`
	UserId   int64  `form:"userId" json:"userId"`
	Nickname string `form:"nickname" json:"nickname"`
	Status   int64  `form:"status" json:"status"`
	OpenId   string `form:"openid" json:"openid"`
}

//8.论坛 点赞
type ForumZanService struct {
	LuntanId int64  `form:"luntanId" json:"luntanId"`
	Type     int64  `form:"type" json:"type"`
	OpenId   string `form:"openid" json:"openid"`
}

//9.测评 首页
type EvaluationIndexService struct {
	Name    string `form:"name" json:"name"`
	TestNum int64  `form:"testnum" json:"testnum"`
	Price   int64  `form:"price" json:"price"`
}

//14测评 post 性格测试
type XinggesService struct {
	Grade  int64  `form:"grade" json:"grade"`
	OpenId string `form:"openid" json:"openid"`
}

//15测评 post MBTI测试
type MBTIService struct {
	Grade  int64  `form:"grade" json:"grade"`
	OpenId string `form:"openid" json:"openid"`
}

//16测评 post 快乐指数测试
type KuaileService struct {
	Grade  int64  `form:"grade" json:"grade"`
	OpenId string `form:"openid" json:"openid"`
}

//17测评 post 情商测试
type QingShangService struct {
	Grade  int64  `form:"grade" json:"grade"`
	OpenId string `form:"openid" json:"openid"`
}

//Grade         int64  `gorm:"type:int(10);not null;default:0;comment:用户得分"`
//Status        int64  `gorm:"type:int(1);not null;default:0;comment:状态值 1正常 2异常"`
//UserName      string `gorm:"size:255;type:char(255);not null;default:0;comment:用户名字"`   // 设置字段大小为255
//NickName      string `gorm:"size:255;type:char(255);not null;default:0;comment:用户昵称"`   // 设置字段大小为255
//Uid           int64  `gorm:"type:int(1);not null;default:10001;comment:用户ID"`           // 设置字段大小为255
//OpenId        string `gorm:"size:255;type:char(255);not null;comment:用户opendid"`        // 设置字段大小为255
//ReportContent string `gorm:"size:255;type:char(255);not null;default:0;comment:报告附加内容"` // 设置字段大小为255
//ReportType    int64  `gorm:"int(10);not null;default:1;comment:报告类型"`                   // 设置字段大小为255
//ReportName    string `gorm:"size:255;type:char(255);not null;default:0;comment:报告名字"`   // 设置字段大小为255

//Name    string `gorm:"size:255;type:char(255);not null;default:0;comment:测试题目名称"` // 设置字段大小为255
//TestNum int64  `gorm:"type:int(11);not null;default:0;comment:已使用过的测试人数 "`        // 设置字段大小为255
//Price   int64  `gorm:"type:int(11);not null;default:0;comment:测试价格"`              // 设置字段大小为255
//Status  int64  `gorm:"type:int(1);not null;default:0;comment:状态值 1正常 2异常"`
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

//18我的 首页
type MineIndexService struct {
	Nickname   string `form:"nickname" json:"nickname"`
	Avatar     string `form:"nickname" json:"nickname"`
	UserPoints int64  `form:"userpoints" json:"userpoints"`
}

//Grade         int64  `gorm:"type:int(10);not null;default:0;comment:用户得分"`
//Status        int64  `gorm:"type:int(1);not null;default:0;comment:状态值 1正常 2异常"`
//UserName      string `gorm:"size:255;type:char(255);not null;default:0;comment:用户名字"`   // 设置字段大小为255
//NickName      string `gorm:"size:255;type:char(255);not null;default:0;comment:用户昵称"`   // 设置字段大小为255
//Uid           int64  `gorm:"type:int(1);not null;default:10001;comment:用户ID"`           // 设置字段大小为255
//ReportContent string `gorm:"size:255;type:char(255);not null;default:0;comment:报告附加内容"` // 设置字段大小为255
//ReportContent string `gorm:"size:255;type:char(255);not null;default:0;comment:报告附加内容"` // 设置字段大小为255
//ReportType    int64  `gorm:"int(10);not null;default:1;comment:报告类型"`                   // 设置字段大小为255
//ReportName    string `gorm:"size:255;type:char(255);not null;default:0;comment:报告名字"`   // 设置字段大小为255
//19我的 报告
type MineReportService struct {
	Grade      string `form:"grade" json:"grade"`
	UserName   string `form:"username" json:"username"`
	UserPoints int64  `form:"userpoints" json:"userpoints"`
	ReportName string `form:"reportname" json:"reportname"`
}

//23我的 从微信获取我的信息
type MineUserinfoService struct {
	UserName string `form:"userName" json:"userName"`
	Nickname string `form:"nickName" json:"nickname"`
	Avatar   string `form:"avatar" json:"avatar"`
	Tell     string `form:"tell" json:"tell"`
	DrawTag  string `form:"drawTag" json:"drawTag"`
	OpenId   string `form:"openId" json:"openId"`
}

//23我的 从微信获取我的信息
type MineReturnUidService struct {
	Code      string `form:"code" json:"code"`
	Appid     string `form:"appid" json:"appid"`
	Appsecret string `form:"appsecret" json:"appsecret"`
}

//24我的 支付
type PayService struct {
	TimeStamp string `form:"timeStamp" json:"timeStamp"`
	NonceStr  string `form:"nonceStr" json:"nonceStr"`
	Package   string `form:"package" json:"package"`
	SignType  string `form:"signType" json:"signType"`
	PaySign   string `form:"paySign" json:"paySign"`
}

type PayServices struct {
	Openid string `form:"openid" json:"openid"`
}

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

//新增论坛
func (service *Luntan) LuntanAdd(c *gin.Context) serializer.Response {
	luntan := model.Luntan{
		Content:  service.Content,
		Img:      service.Img,
		UserId:   service.UserId,
		Nickname: service.Nickname,
		Status:   service.Status,
		Openid:   service.OpenId,
	}
	fmt.Printf("luntan  %v  \n", luntan)

	// 新增数据
	if err := model.DB.Create(&luntan).Error; err != nil {
		return serializer.ParamErr("新增失败", err)
	}
	return serializer.BuildLuntanResponse(luntan)
}

//论坛列表
func (service *Luntan) LuntanList(c *gin.Context, pageId int) serializer.Response1 {
	Offset := (pageId - 1) * 10
	Luntan := make([]model.Luntan, 0)
	var count int64
	//model.DB.Count(&count)
	model.DB.Table("luntans").Count(&count)
	// 获取全部数据
	//db.Limit(10).Offset(5).Find(&users)
	if err := model.DB.Order("id desc").Limit(10).Offset(Offset).Find(&Luntan).Error; err != nil {
		return serializer.ParamErr1("获取数据失败", err)
	}
	//打印结果
	fmt.Printf("count : %v \n", count)
	//return serializer.BuildLuntanResponses(Luntan)
	//luntan := model.Luntan{
	//	Title:   service.Title,
	//	Content: service.Content,
	//	Img:     service.Img,
	//	UserId:  service.UserId,
	//	//Title      string     `gorm:"size:255;type:char(255);not null;default:标题;comment:标题"` // 设置字段大小为255
	//	//Content    string     `gorm:"not null;default:内容;comment:内容"`
	//	//Img        string     `gorm:"not null;default:Img;comment:论坛图片地址"`
	//	//UserId     string     `gorm:"int(10);not null;default:100001;comment:用户ID"`
	//	//Nickname   string     `gorm:"size:255;type:char(255);not null;default:张三;comment:用户昵称"`
	//	//Status     int64      `gorm:"type:int(1);not null;default:0;comment:状态值 1正常 2异常"`
	//	//Sort       int64      `gorm:"type:int(1);not null;default:2;comment:是否置顶 1 是 2否"`
	//	//Zan        int64      `gorm:"type:int(11);not null;default:1;comment:点赞数量"`
	//	//CommentNum int64      `gorm:"type:int(11);default:0;comment:评论数量"`
	//}
	//fmt.Printf("luntan  %v  \n", luntan)

	// 新增数据
	//if err := model.DB.Create(&luntan).Error; err != nil {
	//	return serializer.ParamErr("新增失败", err)
	//}
	return serializer.BuildLuntanResponses(Luntan, count)
}

//func (service *Luntan) ForumPublishComment(c *gin.Context) interface{} {
//
//}

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
		Type:   service.Type,
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
	//Object string `form:"object" json:"object"`
	//Scene  string `form:"scene" json:"scene"`
	//Source string `form:"source" json:"source"`
	//Status int64  `form:"status" json:"status"`
	speechSkillService := model.Craft{
		Object:  service.Object,
		Scene:   service.Scene,
		Content: service.Content,
		Source:  service.Source,
		Status:  service.Status,
		Openid:  service.OpenId,
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

// 1.首页
type SceneService struct {
	Scene  string `gorm:"size:255;type:char(255);not null;default:0;comment:场景"`
	Status int64  `gorm:"type:int(1);not null;default:0;comment:状态值 1正常 2异常"`
}

//var db *gorm.DB

//func TestSelect(userId int) (userList []User) {
//	db.Raw("select * from users where user_id > ?", userId).Scan(&userList)
//	return
//}
func (service *IndexSearchService) Index(c *gin.Context) serializer.Response {
	KeyWordScene := make([]model.Scene, 0)

	//result := map[string]interface{}{}
	//model.DB.Raw("select * from users").Scan(&result)
	//result = db.Exec("select * from users ")

	//db.Table("key_word_scenes").Take(&result)
	//fmt.Printf("result : %v \n", result)
	// 获取全部数据
	//db.Table("users").Select("COALESCE(age,?)", 42).Rows()
	//db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
	//db.Where("name LIKE ?", "%jin%").Find(&users)
	//if err := db.Table("key_word_scenes").Select("Scene").Find(&KeyWordScene).Error; err != nil {
	if err := model.DB.Table("key_word_scenes").Select("Scene,key_word").Find(&KeyWordScene).Error; err != nil {
		return serializer.ParamErr("获取数据失败", err)
	}
	//打印结果
	//fmt.Printf("result : %v \n", KeyWordScene)
	return serializer.BuildKeyWordSceneResponse(KeyWordScene)
}

// 2.搜索
type Crafts struct {
	Content string `gorm:"not null;default:0;comment:内容"`
}

func (service *IndexSearchService) IndexSearch(c *gin.Context) serializer.Response {
	Craft := make([]model.Crafts, 0)
	// 获取全部数据
	//db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
	//db.Where("name LIKE ?", "%jin%").Find(&users)
	if err := model.DB.Select("content,source").Where("status = ? AND keyword LIKE ? OR content LIKE ? OR scene LIKE ?", service.Status, "%"+service.Keyword+"%", "%"+service.Keyword+"%", "%"+service.Keyword+"%").Find(&Craft).Error; err != nil {
		return serializer.ParamErr("获取数据失败", err)
	}
	//打印结果
	fmt.Printf("result : %v \n", Craft)
	return serializer.BuildIndexSearchResponse(Craft)
}

// 7.获取评论列表

func (service *IndexSearchService) ForumCommentList(c *gin.Context, id int64, pageId int) serializer.Response {
	LuntanComment := make([]model.LuntanComment, 0)
	var count int64
	Offset := (pageId - 1) * 10
	//model.DB.Count(&count)
	model.DB.Table("luntan_comments").Where("luntan_id = ? ", id).Count(&count)
	// 获取全部数据
	//db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
	//db.Where("name LIKE ?", "%jin%").Find(&users)
	if err := model.DB.Select("*").Where("luntan_id = ? ", id).Order("id desc").Limit(10).Offset(Offset).Find(&LuntanComment).Error; err != nil {
		return serializer.ParamErr("获取数据失败", err)
	}
	//打印结果
	fmt.Printf("result : %v \n", LuntanComment)
	return serializer.BuildForumCommentListResponse(LuntanComment, count)
}

// 8.论坛 发布评论
func (service *ForumPublishCommentService) ForumPublishComment(c *gin.Context) serializer.Response {
	fmt.Printf("service  %v  \n", service)
	LuntanComment := model.LuntanComment{
		LuntanId: service.LuntanId,
		Content:  service.Content,
		UserId:   service.UserId,
		Nickname: service.Nickname,
		Status:   service.Status,
		Openid:   service.OpenId,
	}
	fmt.Printf("user  %v  \n", LuntanComment)
	// 新增数据
	if err := model.DB.Create(&LuntanComment).Error; err != nil {
		return serializer.ParamErr("新增失败", err)
	}
	return serializer.BuildPublishCommentResponse(LuntanComment)
}

// 8.论坛 点赞
func (service *ForumZanService) ForumZan(c *gin.Context) serializer.Response {
	Luntan := model.Luntan{}
	model.DB.Table("luntans").Select("zan").Where("id = ?", service.LuntanId).Find(&Luntan)
	if service.Type == 1 {
		var num int64
		num = Luntan.Zan + 1
		model.DB.Table("luntans").Where("id = ?", service.LuntanId).Update("zan", num)
	}
	if service.Type == 2 {
		var num int64
		num = Luntan.Zan - 1
		model.DB.Table("luntans").Where("id = ?", service.LuntanId).Update("zan", num)
	}
	// 新增数据
	//model.DB.Table("luntans").Where("id = ?", service.LuntanId).Update("zan", 10)
	//if err := model.DB.Create(&Luntan).Error; err != nil {
	//	return serializer.ParamErr("新增失败", err)
	//}
	return serializer.BuildForumZanResponse(Luntan)
}

// 9.测评 首页
func (service *EvaluationIndexService) EvaluationIndex(c *gin.Context) serializer.Response {
	TestType := []model.TestInfo{}
	model.DB.Table("test_types").Select("name,test_num,price").Where("status = ?", 1).Find(&TestType)
	return serializer.BuildEvaluationIndexResponse(TestType)
}

// 10测评 get 性格测试
func (service *PsychologicalService) EvaluationXingge(c *gin.Context, id int64) serializer.Response {
	Psychological := []model.Psychological{}
	id = 3
	model.DB.Table("psychologicals").Select("*").Where("type = ?", id).Find(&Psychological)
	// 新增数据
	//model.DB.Table("luntans").Where("id = ?", service.LuntanId).Update("zan", 10)
	//if err := model.DB.Create(&Luntan).Error; err != nil {
	//	return serializer.ParamErr("新增失败", err)
	//}
	return serializer.BuildEvaluationXinggeResponse(Psychological)
}

// 11测评 get MBTI测试
func (service *PsychologicalService) EvaluationMBTI(c *gin.Context) serializer.Response {
	Psychological := []model.Psychological{}
	model.DB.Table("psychologicals").Select("*").Where("type = ?", 1).Find(&Psychological)
	// 新增数据
	//model.DB.Table("luntans").Where("id = ?", service.LuntanId).Update("zan", 10)
	//if err := model.DB.Create(&Luntan).Error; err != nil {
	//	return serializer.ParamErr("新增失败", err)
	//}
	return serializer.BuildEvaluationXinggeResponse(Psychological)
}

// 12测评 get 快乐指数测试
func (service *PsychologicalService) EvaluationKuaile(c *gin.Context) serializer.Response {
	Psychological := []model.Psychological{}
	arr := [4]int64{5, 6, 7, 8}
	model.DB.Table("psychologicals").Select("*").Where("type in ?", arr).Find(&Psychological)
	// 新增数据
	//model.DB.Table("luntans").Where("id = ?", service.LuntanId).Update("zan", 10)
	//if err := model.DB.Create(&Luntan).Error; err != nil {
	//	return serializer.ParamErr("新增失败", err)
	//}
	return serializer.BuildEvaluationXinggeResponse(Psychological)
}

// 13测评 get 情商测试
func (service *PsychologicalService) EvaluationQingshang(c *gin.Context) serializer.Response {
	Psychological := []model.Psychological{}
	arr := [4]int64{5, 6, 7, 8}
	model.DB.Table("psychologicals").Select("*").Where("type in ?", arr).Find(&Psychological)
	// 新增数据
	//model.DB.Table("luntans").Where("id = ?", service.LuntanId).Update("zan", 10)
	//if err := model.DB.Create(&Luntan).Error; err != nil {
	//	return serializer.ParamErr("新增失败", err)
	//}
	return serializer.BuildEvaluationXinggeResponse(Psychological)
}

// 14测评 post 性格测试
func (service *XinggesService) ResXingges(c *gin.Context) serializer.Response {
	//fmt.Printf("service  %v  \n", service)
	UserReport := model.UserReport{
		Grade:      service.Grade,
		OpenId:     service.OpenId,
		ReportName: "性格测试",
		Status:     1,
	}
	fmt.Printf("UserReport  %v  \n", UserReport)
	// 新增数据
	if err := model.DB.Create(&UserReport).Error; err != nil {
		return serializer.ParamErr("新增失败", err)
	}
	return serializer.BuildResXinggesResponse(UserReport)
}

// 15测评 post MBTI测试 EvaluationMBTIs
func (service *MBTIService) ResMBTIs(c *gin.Context) serializer.Response {
	UserReport := model.UserReport{
		Grade:      service.Grade,
		OpenId:     service.OpenId,
		ReportName: "MBTI测试",
		Status:     1,
	}
	fmt.Printf("UserReport  %v  \n", UserReport)
	// 新增数据
	if err := model.DB.Create(&UserReport).Error; err != nil {
		return serializer.ParamErr("新增失败", err)
	}
	return serializer.BuildResMBTIsResponse(UserReport)
}

// 16测评 post 快乐指数测试 EvaluationKuailes
func (service *KuaileService) ResKuailes(c *gin.Context) serializer.Response {
	UserReport := model.UserReport{
		Grade:      service.Grade,
		OpenId:     service.OpenId,
		ReportName: "快乐指数测试",
		Status:     1,
	}
	fmt.Printf("UserReport  %v  \n", UserReport)
	// 新增数据
	if err := model.DB.Create(&UserReport).Error; err != nil {
		return serializer.ParamErr("新增失败", err)
	}
	return serializer.BuildResKuailesResponse(UserReport)
}

// 17测评 post 情商测试 EvaluationQingshangs
func (service *QingShangService) ResQingshangs(c *gin.Context) serializer.Response {
	UserReport := model.UserReport{
		Grade:      service.Grade,
		OpenId:     service.OpenId,
		ReportName: "情商测试",
		Status:     1,
	}
	fmt.Printf("UserReport  %v  \n", UserReport)
	// 新增数据
	if err := model.DB.Create(&UserReport).Error; err != nil {
		return serializer.ParamErr("新增失败", err)
	}
	return serializer.BuildResQingshangsResponse(UserReport)
}

// 18.我的 首页
func (service *MineIndexService) MineIndex(c *gin.Context) serializer.Response {
	User := model.UserMine{}
	//arr := [4]int64{5, 6, 7, 8}
	uid := 1
	model.DB.Table("users").Select("id,nickname,user_points,avatar").Where("id = ?", uid).Find(&User)
	// 新增数据
	//model.DB.Table("luntans").Where("id = ?", service.LuntanId).Update("zan", 10)
	//if err := model.DB.Create(&Luntan).Error; err != nil {
	//	return serializer.ParamErr("新增失败", err)
	//}
	return serializer.BuildMineIndexResponse(User)
}

//19.我的 报告
func (service *MineReportService) MineReport(c *gin.Context) serializer.Response {
	UserReportList := []model.UserReportList{}
	//arr := [4]int64{5, 6, 7, 8}
	uid := 1
	model.DB.Table("user_reports").Select("*").Where("id = ?", uid).Find(&UserReportList)
	// 新增数据
	//model.DB.Table("luntans").Where("id = ?", service.LuntanId).Update("zan", 10)
	//if err := model.DB.Create(&Luntan).Error; err != nil {
	//	return serializer.ParamErr("新增失败", err)
	//}
	return serializer.BuildMineReportResponse(UserReportList)
}

//23我的 从微信获取我的信息
func (service *MineUserinfoService) MineUserinfo(c *gin.Context) serializer.Response {
	//UserName string `form:"userName" json:"userName"`
	//Nickname string `form:"nickname" json:"nickname"`
	//Avatar   string `form:"avatar" json:"avatar"`
	//Tell     string `form:"tell" json:"tell"`
	//DrawTag  string `form:"drawTag" json:"drawTag"`
	//OpenId   string `form:"openId" json:"openId"`
	userService := model.User{
		UserName: service.UserName,
		Nickname: service.Nickname,
		Avatar:   service.Avatar,
		Tell:     service.Tell,
		DrawTag:  service.DrawTag,
		OpenId:   service.OpenId,
	}
	fmt.Printf("userService  %v  \n", userService)
	// 新增数据
	if err := model.DB.Create(&userService).Error; err != nil {
		return serializer.ParamErr("新增失败", err)
	}
	return serializer.BuildMineUserinfoResponse(userService)
}

const (
	code2sessionURL = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	//appID           = "你的appid"		wx6902b88cb7e7ea61
	//appSecret       = "你的appsecret" 1f568972901352eb6814e4dc1b9d50e1
	//Code								053dzQml22LZp94GGMkl2Bdmcl4dzQmJ
)

//24我的 返回系统定义的uid为openid
func (service *MineReturnUidService) MineReturnUid(c *gin.Context) serializer.Response {
	//获取code
	//code := c.PostForm("code")
	UserID := model.User{}
	model.DB.Table("users").Select("id").Where("user_code = ?", service.Code).Find(&UserID)
	if UserID.ID != 0 {
		return serializer.Response{
			Error: "用户已经存在无法新增",
		}
	}
	fmt.Printf("UserID  : %v \n", UserID.ID)
	fmt.Printf("11111  : %v \n", UserID.ID)

	//调用auth.code2Session接口获取openid
	//url := fmt.Sprintf(code2sessionURL, service.Appid, service.Appsecret, service.Code)
	//resp, err := http.Get(url)
	//if err != nil {
	//	return serializer.Response{
	//		Error: "用户已经存在无法新增",
	//	}
	//}
	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//json := gojsonq.New().FromString(string(body)).Find("openid")
	//openId := json.(string)
	//fmt.Println("my openid is: ", openId)

	//调用auth.code2Session接口获取openid
	url := fmt.Sprintf(code2sessionURL, service.Appid, service.Appsecret, service.Code)
	resp, err := http.Get(url)
	if err != nil {
		return serializer.Response{
			Error: "获取微信openid失败",
		}
	}
	fmt.Printf("22  : %v \n", resp)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	fmt.Printf("331  : %v \n", body)
	fmt.Printf("332  : %v \n", err)
	//
	json := gojsonq.New().FromString(string(body)).Find("openid")
	//微信code  5分钟之后过期
	if json == nil {
		return serializer.Response{
			Error: "获取微信openid失败 需要重新code",
		}
	}
	fmt.Printf("333  : %v \n", json)
	openId := json.(string)
	fmt.Printf("334  : %v \n", json)
	fmt.Println("my openid is: ", openId)

	//id := uuid.NewV4()
	//ids := id.String()
	//fmt.Printf("id  : %v \n", id)
	//fmt.Printf("ids  : %v \n", ids)
	userService := model.User{
		UserCode:  service.Code,
		Appid:     service.Appid,
		AppSecret: service.Appsecret,
		OpenId:    openId,
	}
	fmt.Printf("userService  %v  \n", userService)
	fmt.Printf("333335  : %v \n", UserID.ID)

	// 新增数据
	if err := model.DB.Create(&userService).Error; err != nil {
		fmt.Printf("44  : %v \n", UserID.ID)

		return serializer.ParamErr("新增失败", err)
	}
	fmt.Printf("55  : %v \n", UserID.ID)

	return serializer.BuildMineReturnUidResponse(userService)
}

//25我的 支付
func (service *MineReturnUidService) Pay(c *gin.Context) serializer.Response {
	UserID := model.User{}
	id := uuid.NewV4()
	ids := id.String()
	fmt.Printf("id  : %v \n", id)
	fmt.Printf("ids  : %v \n", ids)
	model.DB.Table("users").Select("id").Where("user_code = ?", service.Code).Find(&UserID)
	if UserID.ID != 0 {
		return serializer.Response{
			Error: "用户已经存在无法新增",
		}
	}
	fmt.Printf("UserID  : %v \n", UserID.ID)
	userService := model.User{
		UserCode:  service.Code,
		Appid:     service.Appid,
		AppSecret: service.Appsecret,
		OpenId:    ids,
	}
	fmt.Printf("userService  %v  \n", userService)

	// 新增数据
	if err := model.DB.Create(&userService).Error; err != nil {
		return serializer.ParamErr("新增失败", err)
	}
	return serializer.BuildMineReturnUidResponse(userService)
}
