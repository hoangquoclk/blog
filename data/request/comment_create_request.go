package request

type CommentCreateRequest struct {
	Content string `validate:"required min=1,max=100" json:"content"`
}
