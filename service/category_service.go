package service

import (
	"context"
	"example/blog/data/request"
	"example/blog/data/response"
)

type CategoryService interface {
	Create(ctx context.Context, request request.CategoryCreateRequest)
	Update(ctx context.Context, request request.CategoryUpdateRequest)
	Delete(ctx context.Context, categoryId string)
	FindById(ctx context.Context, categoryId string) response.CategoryResponse
	FindAll(ctx context.Context) []response.CategoryResponse
}
