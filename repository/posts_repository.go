package repository

import (
	"context"
	"example/blog/model"
)

type PostRepository interface {
	Save(ctx context.Context, post model.Post)
	Update(ctx context.Context, post model.Post)
	Delete(ctx context.Context, postId int)
	FindById(ctx context.Context, postId int) (model.Post, error)
	FindAll(ctx context.Context) []model.Post
}
