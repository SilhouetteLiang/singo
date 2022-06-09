package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
	//Title      string `gorm:"size:255;type:char(255);not null;default:标题;comment:标题"` // 设置字段大小为255
	//Content    string `gorm:"not null;default:内容;comment:内容"`
	//Img        string `gorm:"not null;default:Img;comment:论坛图片地址"`
	//UserId     string `gorm:"int(10);not null;default:100001;comment:用户ID"`
	//Nickname   string `gorm:"size:255;type:char(255);not null;default:张三;comment:用户昵称"`
	//Status     int64  `gorm:"type:int(1);not null;default:0;comment:状态值 1正常 2异常"`
	//Sort       int64  `gorm:"type:int(1);not null;default:2;comment:是否置顶 1 是 2否"`
	//Zan        int64  `gorm:"type:int(11);not null;default:1;comment:点赞数量"`
	//CommentNum int64  `gorm:"type:int(11);default:0;comment:评论数量"`
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
}

//8.发布评论
type ForumPublishCommentService struct {
	LuntanId int64  `form:"luntanId" json:"luntanId"`
	Content  string `form:"content" json:"content"`
	UserId   int64  `form:"userId" json:"userId"`
	Nickname string `form:"nickname" json:"nickname"`
	Status   int64  `form:"status" json:"status"`
}

//8.发布评论
type ForumZanService struct {
	LuntanId int64 `form:"luntanId" json:"luntanId"`
	Type     int64 `form:"type" json:"type"`
}

//LuntanId   int64      `gorm:"int(11);not null;default:99999999;comment:帖子ID"`
//Content    string     `gorm:"not null;default:内容;comment:评论内容"`
//UserId     string     `gorm:"int(10);not null;default:100001;comment:用户ID"`
//Nickname   string     `gorm:"size:255;type:char(255);not null;default:张三;comment:用户昵称"`
//Status     int64      `gorm:"type:int(1);not null;default:0;comment:状态值 1正常 2异常"`

//`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
//`created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
//`updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
//`deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
//`luntan_id` bigint(20) NOT NULL DEFAULT '99999999' COMMENT '帖子ID',
//`content` varchar(191) NOT NULL DEFAULT '内容' COMMENT '评论内容',
//`user_id` varchar(191) NOT NULL DEFAULT '100001' COMMENT '用户ID',
//`nickname` char(255) NOT NULL DEFAULT '张三' COMMENT '用户昵称',
//`status` int(1) NOT NULL DEFAULT '0' COMMENT '状态值 1正常 2异常',

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
	fmt.Printf("1111111  %v  \n")
	luntan := model.Luntan{
		Content:  service.Content,
		Img:      service.Img,
		UserId:   service.UserId,
		Nickname: service.Nickname,
		Status:   service.Status,
	}
	fmt.Printf("luntan  %v  \n", luntan)

	// 新增数据
	if err := model.DB.Create(&luntan).Error; err != nil {
		return serializer.ParamErr("新增失败", err)
	}
	return serializer.BuildLuntanResponse(luntan)
}

//论坛列表
func (service *Luntan) LuntanList(c *gin.Context) serializer.Response {
	Luntan := make([]model.Luntan, 0)
	// 获取全部数据
	if err := model.DB.Find(&Luntan).Error; err != nil {
		return serializer.ParamErr("获取数据失败", err)
	}
	//打印结果
	//fmt.Printf("result : %v \n", psychological)
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
	return serializer.BuildLuntanResponses(Luntan)
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

func (service *IndexSearchService) ForumCommentList(c *gin.Context, id int64) serializer.Response {
	LuntanComment := make([]model.LuntanComment, 0)
	// 获取全部数据
	//db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
	//db.Where("name LIKE ?", "%jin%").Find(&users)
	if err := model.DB.Select("*").Where("luntan_id = ? ", id).Find(&LuntanComment).Error; err != nil {
		return serializer.ParamErr("获取数据失败", err)
	}
	//打印结果
	fmt.Printf("result : %v \n", LuntanComment)
	return serializer.BuildForumCommentListResponse(LuntanComment)
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
	}
	fmt.Printf("user  %v  \n", LuntanComment)
	// 新增数据
	if err := model.DB.Create(&LuntanComment).Error; err != nil {
		return serializer.ParamErr("新增失败", err)
	}
	return serializer.BuildPublishCommentResponse(LuntanComment)
}

// 9.论坛 点赞
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

// 10测评 get 性格测试
func (service *PsychologicalService) EvaluationXingge(c *gin.Context, id int64) serializer.Response {
	Psychological := []model.Psychological{}
	model.DB.Table("psychologicals").Select("*").Where("type = ?", id).Find(&Psychological)
	// 新增数据
	//model.DB.Table("luntans").Where("id = ?", service.LuntanId).Update("zan", 10)
	//if err := model.DB.Create(&Luntan).Error; err != nil {
	//	return serializer.ParamErr("新增失败", err)
	//}
	return serializer.BuildEvaluationXinggeResponse(Psychological)
}
