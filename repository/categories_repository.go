package repository

import (
	"context"
	"example/blog/model"
)

type CategoryRepository interface {
	Save(ctx context.Context, category model.Category)
	Update(ctx context.Context, category model.Category)
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) (model.Category, error)
	FindAll(ctx context.Context) []model.Category
}
