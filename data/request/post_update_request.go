package request

type PostUpdateRequest struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
