package request

type CommentCreateRequest struct {
	Content string `validate:"required min=1,max=100" json:"content"`
	UserId  string `validate:"require" json:"userId"`
	PostId  string `validate:"require" json:"postId"`
}
