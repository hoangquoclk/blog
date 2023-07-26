package service

import (
	"context"
	"example/blog/data/request"
	"example/blog/data/response"
	"example/blog/helper"
	"example/blog/model"
	"example/blog/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
}

func NewCategoryServiceImpl(categoryRepository repository.CategoryRepository) CategoryService {
	return &CategoryServiceImpl{CategoryRepository: categoryRepository}
}

func (c *CategoryServiceImpl) Create(ctx context.Context, request request.CategoryCreateRequest) {
	category := model.Category{
		Name: request.Name,
	}
	c.CategoryRepository.Save(ctx, category)
}

func (c *CategoryServiceImpl) Update(ctx context.Context, request request.CategoryUpdateRequest) {
	category, err := c.CategoryRepository.FindById(ctx, request.Id)
	helper.PanicIfErrors(err)

	category.Name = request.Name
	c.CategoryRepository.Update(ctx, category)
}

func (c *CategoryServiceImpl) Delete(ctx context.Context, categoryId string) {
	category, err := c.CategoryRepository.FindById(ctx, categoryId)
	helper.PanicIfErrors(err)
	c.CategoryRepository.Delete(ctx, category.Id)
}

func (c *CategoryServiceImpl) FindById(ctx context.Context, categoryId string) response.CategoryResponse {
	category, err := c.CategoryRepository.FindById(ctx, categoryId)
	helper.PanicIfErrors(err)
	return response.CategoryResponse(category)
}

func (c *CategoryServiceImpl) FindAll(ctx context.Context) []response.CategoryResponse {
	categories := c.CategoryRepository.FindAll(ctx)

	var categoryResp []response.CategoryResponse

	for _, value := range categories {
		category := response.CategoryResponse{Id: value.Id, Name: value.Name}
		categoryResp = append(categoryResp, category)
	}
	return categoryResp
}
