package service

import (
	"context"
	"example/blog/data/request"
	"example/blog/data/response"
)

type CommentService interface {
	Create(ctx context.Context, request request.CommentCreateRequest)
	Update(ctx context.Context, request request.CommentUpdateRequest)
	Delete(ctx context.Context, userId int)
	FindById(ctx context.Context, userId int) response.CommentResponse
	FindAll(ctx context.Context) []response.CommentResponse
}
