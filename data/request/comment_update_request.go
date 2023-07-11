package request

type CommentUpdateRequest struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
}
