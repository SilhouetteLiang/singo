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
}

// BuildUserResponse 序列化内容响应
func BuildPsychologicalResponses(psychological model.Psychological) Response {
	return Response{
		Data: BuildpSychological(psychological),
	}
}

// BuildUser 序列化内容
func BuildpSychological(psychological model.Psychological) Psychological {
	return Psychological{
		ID:        psychological.ID,
		Title:     psychological.Title,
		Status:    psychological.Status,
		CreatedAt: psychological.CreatedAt.Unix(),
	}
}
