package request

type CommentUpdateRequest struct {
	Id      string `json:"id"`
	Content string `json:"content"`
}
