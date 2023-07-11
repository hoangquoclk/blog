package service

import (
	"context"
	"example/blog/data/request"
	"example/blog/data/response"
	"example/blog/helper"
	"example/blog/model"
	"example/blog/repository"
)

type PostServiceImpl struct {
	PostRepository repository.PostRepository
}

func NewPostServiceImpl(postRepository repository.PostRepository) PostService {
	return &PostServiceImpl{PostRepository: postRepository}
}

func (p *PostServiceImpl) Create(ctx context.Context, request request.PostCreateRequest) {
	post := model.Post{
		Title:   request.Title,
		Content: request.Content,
	}
	p.PostRepository.Save(ctx, post)
}

func (p *PostServiceImpl) Update(ctx context.Context, request request.PostUpdateRequest) {
	post, err := p.PostRepository.FindById(ctx, request.Id)
	helper.PanicIfErrors(err)

	post.Title = request.Title
	post.Content = request.Content
	p.PostRepository.Update(ctx, post)
}

func (p *PostServiceImpl) Delete(ctx context.Context, postId int) {
	post, err := p.PostRepository.FindById(ctx, postId)
	helper.PanicIfErrors(err)
	p.PostRepository.Delete(ctx, post.Id)
}

func (p *PostServiceImpl) FindById(ctx context.Context, postId int) response.PostResponse {
	post, err := p.PostRepository.FindById(ctx, postId)
	helper.PanicIfErrors(err)
	return response.PostResponse(post)
}

func (p *PostServiceImpl) FindAll(ctx context.Context) []response.PostResponse {
	posts := p.PostRepository.FindAll(ctx)

	var postResp []response.PostResponse

	for _, value := range posts {
		user := response.PostResponse{Id: value.Id, Title: value.Title, CategoryId: value.CategoryId, UserId: value.UserId, Content: value.Content}
		postResp = append(postResp, user)
	}
	return postResp
}
