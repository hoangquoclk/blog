package repository

import (
	"context"
	"example/blog/model"
)

type CommentRepository interface {
	Save(ctx context.Context, comment model.Comment)
	Update(ctx context.Context, comment model.Comment)
	Delete(ctx context.Context, commentId int)
	FindById(ctx context.Context, commentId int) (model.Comment, error)
	FindAll(ctx context.Context) []model.Comment
}
