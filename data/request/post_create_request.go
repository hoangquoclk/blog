package request

type PostCreateRequest struct {
	Title      string `validate:"required min=1,max=100" json:"title"`
	Content    string `validate:"required min=1,max=100" json:"content"`
	UserId     int    `validate:"require" json:"user_id"`
	CategoryId int    `validate:"require" json:"category_id"`
}
