package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/thedevsaddam/gojsonq/v2"
	"io/ioutil"
	"net/http"
	"singo/model"
	"singo/serializer"
)

//dsadasd
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
	Nickname string `form:"nickname" json:"nickname"`
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
	Grade  string `form:"grade" json:"grade"`
	OpenId string `form:"openid" json:"openid"`
}

//15测评 post MBTI测试
type MBTIService struct {
	Grade  string `form:"grade" json:"grade"`
	OpenId string `form:"openid" json:"openid"`
}

//16测评 post 快乐指数测试
type KuaileService struct {
	Grade  string `form:"grade" json:"grade"`
	OpenId string `form:"openid" json:"openid"`
}

//17测评 post 情商测试
type QingShangService struct {
	Grade  string `form:"grade" json:"grade"`
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

//19我的 报告
type MineReportService struct {
	OpenId string `form:"openid" json:"openid"`
}

//20我的 GET我发布的话术 desc
type MinePublicService struct {
	OpenId string `form:"openid" json:"openid"`
}

//23我的 从微信获取我的信息
type MineUserinfoService struct {
	UserName     string `form:"userName" json:"userName"`
	Nickname     string `form:"nickName" json:"nickname"`
	Avatar       string `form:"avatar" json:"avatar"`
	Tell         string `form:"tell" json:"tell"`
	DrawTag      string `form:"drawTag" json:"drawTag"`
	OpenId       string `form:"openId" json:"openId"`
	InvitationId string `form:"invitation_id" json:"invitation_id"`
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
	Amount int64  `form:"amount" json:"amount"`
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
		Status:   1,
		Openid:   service.OpenId,
	}
	fmt.Printf("luntan  %v  \n", luntan)

	// 新增数据
	if err := model.DB.Create(&luntan).Error; err != nil {
		return serializer.ParamErr("新增失败", err)
	}
	return serializer.BuildLuntanResponse(luntan)
}

//5论坛 GET列表 		desc
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
	service.Status = 1
	// 获取全部数据
	//db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
	//db.Where("name LIKE ?", "%jin%").Find(&users)
	if err := model.DB.Select("content,source").Where("status = ? AND openid = ? OR keyword LIKE ? OR content LIKE ? OR scene LIKE ?", service.Status, service.OpenId, "%"+service.Keyword+"%", "%"+service.Keyword+"%", "%"+service.Keyword+"%").Find(&Craft).Error; err != nil {
		return serializer.ParamErr("获取数据失败", err)
	}
	//打印结果
	fmt.Printf("result : %v \n", Craft)
	return serializer.BuildIndexSearchResponse(Craft)
}

//4首页 POST发布话术
func (service *SpeechSkillService) InputSpeechCraft(c *gin.Context) serializer.Response {
	speechSkillService := model.Craft{
		Object:  service.Object,
		Scene:   service.Scene,
		Content: service.Content,
		Source:  service.Source,
		Status:  1,
		Openid:  service.OpenId,
	}

	// 新增数据
	if err := model.DB.Create(&speechSkillService).Error; err != nil {
		return serializer.ParamErr("新增失败", err)
	}
	return serializer.BuildCraftResponse(speechSkillService)
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
	//fmt.Printf("service  %v  \n", service)
	LuntanComment := model.LuntanComment{
		LuntanId: service.LuntanId,
		Content:  service.Content,
		Nickname: service.Nickname,
		Status:   1,
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
	Luntan := model.LuntanZan{}
	LuntanComment := model.LuntanComment{
		LuntanId: service.LuntanId,
		Openid:   service.OpenId,
		IsZan:    3,
	}
	//key := fmt.Sprintf("%v_%v", service.LuntanId, service.OpenId)
	//fmt.Printf("key  : %v \n", key)
	//// 添加string
	//cache.RedisClient.Set("golang_redis", "golang", 0)
	////ClientRedis.Set("golang_string", "golang", 0)
	//// 获取string
	//stringCmd := cache.RedisClient.Get("golang_redis")
	//fmt.Println(stringCmd.String(), stringCmd.Args(), stringCmd.Val())
	//// 删除string
	//cache.RedisClient.Del("golang_redis")
	var isDisplay int64
	model.DB.Table("luntans").Select("zan").Where("id = ?", service.LuntanId).Find(&Luntan)
	model.DB.Table("luntan_comments").Select("id").Where("luntan_id = ? And openid = ?", service.LuntanId, service.OpenId).Find(&LuntanComment)
	//LuntanComment
	// 新增数据
	if LuntanComment.ID == 0 {
		if err := model.DB.Create(&LuntanComment).Error; err != nil {
			return serializer.ParamErr("新增luntan_comments 数据失败", err)
		}
	}
	//fmt.Printf("LuntanComment11111111111  : %v \n", LuntanComment.ID)

	if service.Type == 1 {
		fmt.Printf("Type  : %v \n", service.Type)
		var num int64
		num = Luntan.Zan + 1
		res1 := model.DB.Table("luntans").Where("id = ?", service.LuntanId).Update("zan", num)
		if res1.RowsAffected == 0 {
			return serializer.Response{
				Error: "1请查看luntanId是否正确",
			}
		}

		res2 := model.DB.Table("luntan_comments").Where("luntan_id = ? And openid = ?", service.LuntanId, service.OpenId).Update("is_zan", 1)
		if res2.RowsAffected == 0 {
			return serializer.Response{
				Error: "A该用户已经点赞不可以在点赞",
			}
		}
		isDisplay = 1
	}
	if service.Type == 2 {
		fmt.Printf("Type  : %v \n", service.Type)

		var num int64
		num = Luntan.Zan - 1
		res1 := model.DB.Table("luntans").Where("id = ?", service.LuntanId).Update("zan", num)
		if res1.RowsAffected == 0 {
			return serializer.Response{
				Error: "2请查看 luntanId 是否正确",
			}
		}

		res2 := model.DB.Table("luntan_comments").Where("luntan_id = ? And openid = ?", service.LuntanId, service.OpenId).Update("is_zan", 2)
		if res2.RowsAffected == 0 {
			return serializer.Response{
				Error: "B该用户已经取消点赞不可以连续取消赞",
			}
		}
		isDisplay = 2
	}
	// 新增数据
	//model.DB.Table("luntans").Where("id = ?", service.LuntanId).Update("zan", 10)
	//if err := model.DB.Create(&Luntan).Error; err != nil {
	//	return serializer.ParamErr("新增失败", err)
	//}
	return serializer.BuildForumZanResponse(Luntan, isDisplay)
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
func (service *MineIndexService) MineIndex(c *gin.Context, openid string) serializer.Response {
	User := model.UserMine{}
	//arr := [4]int64{5, 6, 7, 8}
	//uid := 1
	model.DB.Table("users").Select("id,nickname,user_points,avatar").Where("open_id = ?", openid).Find(&User)
	// 新增数据
	//model.DB.Table("luntans").Where("id = ?", service.LuntanId).Update("zan", 10)
	//if err := model.DB.Create(&Luntan).Error; err != nil {
	//	return serializer.ParamErr("新增失败", err)
	//}
	return serializer.BuildMineIndexResponse(User)
}

//19.我的 报告
func (service *MineReportService) MineReport(c *gin.Context, openid string) serializer.Response {
	UserReportList := []model.UserReportList{}
	//arr := [4]int64{5, 6, 7, 8}
	//uid := 1
	model.DB.Table("user_reports").Select("*").Where("open_id = ?", openid).Find(&UserReportList)
	// 新增数据
	//model.DB.Table("luntans").Where("id = ?", service.LuntanId).Update("zan", 10)
	//if err := model.DB.Create(&Luntan).Error; err != nil {
	//	return serializer.ParamErr("新增失败", err)
	//}
	return serializer.BuildMineReportResponse(UserReportList)
}

//20我的 GET我发布的话术 desc
func (service *MinePublicService) MinePublic(c *gin.Context, openid string) serializer.Response {
	MyCraft := []model.MyCraft{}
	model.DB.Table("crafts").Select("*").Where("openid = ?", openid).Find(&MyCraft)
	return serializer.BuildMinePublicResponse(MyCraft)
}

//23我的 从微信获取我的信息
func (service *MineUserinfoService) MineUserinfo(c *gin.Context) serializer.Response {
	var Point int64
	if len(service.InvitationId) == 0 {
		Point = 0
	}
	if len(service.InvitationId) != 0 {
		Point = 3
	}
	//获取用户信息
	UserInfo := model.User{}
	model.DB.Table("users").Select("id").Where("open_id = ?", service.OpenId).Find(&UserInfo)
	//1.老用户修改数据
	if UserInfo.ID != 0 { //
		UserPoints := 3 + UserInfo.UserPoints
		model.DB.Table("users").Where("open_id = ?", service.OpenId).Update("user_points", UserPoints)
	}
	//2.新用户新增数据
	if UserInfo.ID != 0 { //
		userService := model.User{
			UserName:   service.UserName,
			Nickname:   service.Nickname,
			Avatar:     service.Avatar,
			Tell:       service.Tell,
			DrawTag:    service.DrawTag,
			OpenId:     service.OpenId,
			UserPoints: Point,
		}
		// 新增数据
		if err := model.DB.Create(&userService).Error; err != nil {
			return serializer.ParamErr("新增失败", err)
		}
	}
	fmt.Printf("UserID  : %v \n", UserInfo.ID)
	//db.Last(&user)
	user := model.User{}
	model.DB.Table("users").Last(&user)

	return serializer.BuildMineUserinfoResponse(user)
}

const (
	//code2sessionURL = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	code2sessionURL = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	appID           = "wx6902b88cb7e7ea61"
	//appSecret       = "1f568972901352eb6814e4dc1b9d50e1" //1f568972901352eb6814e4dc1b9d50e1
	appSecret = "85d7cc18ef1845aacd98235ecc1c2df6" //85d7cc18ef1845aacd98235ecc1c2df6

)

type SnsOauth2 struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	Openid       string `json:"openid"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

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
	//调用auth.code2Session接口获取openid
	url := fmt.Sprintf(code2sessionURL, appID, appSecret, service.Code)
	fmt.Printf("service.Code  : %v \n", service.Code)
	resp, err := http.Get(url)
	if err != nil {
		return serializer.Response{
			Error: "获取微信openid失败",
		}
	}
	fmt.Printf("resp  : %v \n", resp)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	atr := SnsOauth2{}
	err = json.Unmarshal(body, &atr)
	if err != nil {
		fmt.Println("发送get请求获取 openid 返回数据json解析错误", err)
	}
	fmt.Printf("atr  : %v \n", atr)
	json := gojsonq.New().FromString(string(body)).Find("openid")
	//微信code  5分钟之后过期
	fmt.Printf("json  : %v \n", json)
	if json == nil {
		return serializer.Response{
			Error: "获取微信openid失败 需要重新code",
		}
	}
	openId := json.(string)
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
