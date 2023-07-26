package request

type PostCreateRequest struct {
	Title      string `validate:"required min=1,max=100" json:"title"`
	Content    string `validate:"required min=1,max=100" json:"content"`
	UserId     string `validate:"require" json:"userId"`
	CategoryId string `validate:"require" json:"categoryId"`
}
