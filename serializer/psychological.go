package serializer

import (
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
	ID         uint   `json:"id"`
	Content    string `gorm:"not null;default:内容;comment:内容"`
	Img        string `gorm:"not null;default:Img;comment:论坛图片地址"`
	UserId     string `gorm:"int(10);not null;default:100001;comment:用户ID"`
	Nickname   string `gorm:"size:255;type:char(255);not null;default:张三;comment:用户昵称"`
	Status     int64  `gorm:"type:int(1);not null;default:0;comment:状态值 1正常 2异常"`
	Sort       int64  `gorm:"type:int(1);not null;default:2;comment:是否置顶 1 是 2否"`
	Zan        int64  `gorm:"type:int(11);not null;default:1;comment:点赞数量"`
	CommentNum int64  `gorm:"type:int(11);default:0;comment:评论数量"`
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

//序列化 论坛 响应
func BuildLuntanResponses(Luntan []model.Luntan) Response {
	return Response{
		//Data: BuildpSychological(psychological),
		Data: Luntan,
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

func BuildForumCommentListResponse(Craft []model.LuntanComment) Response {
	return Response{
		Data: Craft,
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
		//LuntanId   int64      `gorm:"int(11);not null;default:99999999;comment:帖子ID"`
		//Content    string     `gorm:"not null;default:内容;comment:评论内容"`
		//UserId     int64      `gorm:"int(10);not null;default:100001;comment:用户ID"`
		//Nickname   string     `gorm:"size:255;type:char(255);not null;default:张三;comment:用户昵称"`
		//Status     int64      `gorm:"type:int(1);not null;default:0;comment:状态值 1正常 2异常"`
		LuntanId: LuntanComment.LuntanId,
		Content:  LuntanComment.Content,
		UserId:   LuntanComment.UserId,
		Nickname: LuntanComment.Nickname,
		Status:   LuntanComment.Status,
	}
}

//序列化响应 9 论坛 点赞
func BuildForumZanResponse(Luntan model.Luntan) Response {
	return Response{
		Data: BuildForumZan(Luntan),
	}
}

//序列化内容 9 论坛 点赞
func BuildForumZan(Luntan model.Luntan) LunTans {
	return LunTans{
		//LuntanId   int64      `gorm:"int(11);not null;default:99999999;comment:帖子ID"`
		//Content    string     `gorm:"not null;default:内容;comment:评论内容"`
		//UserId     int64      `gorm:"int(10);not null;default:100001;comment:用户ID"`
		//Nickname   string     `gorm:"size:255;type:char(255);not null;default:张三;comment:用户昵称"`
		//Status     int64      `gorm:"type:int(1);not null;default:0;comment:状态值 1正常 2异常"`
		Zan: Luntan.Zan,
	}
}

//序列化响应 10 测评 get 性格测试
func BuildEvaluationXinggeResponse(Psychological []model.Psychological) Response {
	return Response{
		Data: Psychological,
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
