package serializer

import "singo/model"

// Content 内容序列化器
type Content struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Content   string `json:"content"`
	Status    string `json:"status"`
	CreatedAt int64  `json:"created_at"`
}

// BuildUserResponse 序列化内容响应
func BuildUserResponses(content model.Content) Response {
	return Response{
		Data: BuildContent(content),
	}
}

// BuildUser 序列化内容
func BuildContent(content model.Content) Content {
	return Content{
		ID:        content.ID,
		Content:   content.Content,
		UserName:  content.UserName,
		Status:    content.Status,
		CreatedAt: content.CreatedAt.Unix(),
	}
}
