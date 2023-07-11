package service

import (
	"context"
	"example/blog/data/request"
	"example/blog/data/response"
)

type CategoryService interface {
	Create(ctx context.Context, request request.CategoryCreateRequest)
	Update(ctx context.Context, request request.CategoryUpdateRequest)
	Delete(ctx context.Context, userId int)
	FindById(ctx context.Context, userId int) response.CategoryResponse
	FindAll(ctx context.Context) []response.CategoryResponse
}
