package service

import (
	"context"
	"example/blog/data/request"
	"example/blog/data/response"
)

type CommentService interface {
	Create(ctx context.Context, request request.CommentCreateRequest)
	Update(ctx context.Context, request request.CommentUpdateRequest)
	Delete(ctx context.Context, userId string)
	FindById(ctx context.Context, userId string) response.CommentResponse
	FindAll(ctx context.Context) []response.CommentResponse
}
