package response

type PostResponse struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	CategoryId int    `json:"category_id"`
	UserId     int    `json:"user_id"`
	Content    string `json:"content"`
}
