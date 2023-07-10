package service

import (
	"context"
	"example/blog/data/request"
	"example/blog/data/response"
)

type UserService interface {
	Create(ctx context.Context, request request.UserCreateRequest)
	Update(ctx context.Context, request request.UserUpdateRequest)
	Delete(ctx context.Context, userId int)
	FindById(ctx context.Context, userId int) response.UserResponse
	FindAll(ctx context.Context) []response.UserResponse
}
