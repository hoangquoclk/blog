package request

type CommentCreateRequest struct {
	Content string `validate:"required min=1,max=100" json:"content"`
	UserId  int    `validate:"require" json:"user_id"`
	PostId  int    `validate:"require" json:"post_id"`
}
