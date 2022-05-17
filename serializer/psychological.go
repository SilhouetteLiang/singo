package serializer

import (
	"singo/model"
)

// Psychological 内容序列化器
type Psychological struct {
	ID        uint   `json:"id"`
	Status    uint64 `json:"status"`
	Title     string `json:"title"`
	CreatedAt int64  `json:"created_at"`
	A         string
	B         string
	C         string
	D         string
}

// Craft 内容序列化器
type Craft struct {
	Title string `json:"title"`
	A     string
	B     string
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

type ResPsychological struct {
	Items []Item `json:"items"`
}

type Item struct {
	Title  string `form:"title" json:"title" binding:"required,min=5,max=30"`
	Status uint64 `form:"status" json:"status" binding:"required"`
}

//序列化响应 新增用户
func BuildUsersResponse(user model.User) Response {
	return Response{
		Data: BuildUsers(user),
	}
}

//序列化内容 新增用户
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
		Title: psychological.Title,
		A:     psychological.A,
		B:     psychological.B,
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
