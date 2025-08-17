package request

type RequestComment struct {
	PostId  uint   `json:"postId"`
	Content string `json:"content"`
}
