package service

import (
	"context"
	"example/blog/data/request"
	"example/blog/data/response"
)

type PostService interface {
	Create(ctx context.Context, request request.PostCreateRequest)
	Update(ctx context.Context, request request.PostUpdateRequest)
	Delete(ctx context.Context, userId string)
	FindById(ctx context.Context, userId string) response.PostResponse
	FindAll(ctx context.Context) []response.PostResponse
}
