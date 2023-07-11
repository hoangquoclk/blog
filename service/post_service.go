package service

import (
	"context"
	"example/blog/data/request"
	"example/blog/data/response"
)

type PostService interface {
	Create(ctx context.Context, request request.PostCreateRequest)
	Update(ctx context.Context, request request.PostUpdateRequest)
	Delete(ctx context.Context, userId int)
	FindById(ctx context.Context, userId int) response.PostResponse
	FindAll(ctx context.Context) []response.PostResponse
}
