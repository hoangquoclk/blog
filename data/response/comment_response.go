package response

type CommentResponse struct {
	Id      string `json:"id"`
	PostId  string `json:"postId"`
	UserId  string `json:"userId"`
	Content string `json:"content"`
}
