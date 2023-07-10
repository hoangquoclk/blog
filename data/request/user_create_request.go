package request

type UserCreateRequest struct {
	Username string `validate:"required min=1,max=100" json:"username"`
	Password string `validate:"required min=1,max=100" json:"password"`
	Email    string `validate:"required min=1,max=100" json:"email"`
}
