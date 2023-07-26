package request

type PostUpdateRequest struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
