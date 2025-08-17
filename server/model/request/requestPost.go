package request

type RequestPost struct {
	PostId  uint   `json:"postId"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
