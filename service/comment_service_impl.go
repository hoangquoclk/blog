package service

import (
	"context"
	"example/blog/data/request"
	"example/blog/data/response"
	"example/blog/helper"
	"example/blog/model"
	"example/blog/repository"
)

type CommentServiceImpl struct {
	CommentRepository repository.CommentRepository
}

func NewCommentServiceImpl(commentRepository repository.CommentRepository) CommentService {
	return &CommentServiceImpl{CommentRepository: commentRepository}
}

func (c *CommentServiceImpl) Create(ctx context.Context, request request.CommentCreateRequest) {
	comment := model.Comment{
		Content: request.Content,
		UserId:  request.UserId,
		PostId:  request.PostId,
	}
	c.CommentRepository.Save(ctx, comment)
}

func (c *CommentServiceImpl) Update(ctx context.Context, request request.CommentUpdateRequest) {
	comment, err := c.CommentRepository.FindById(ctx, request.Id)
	helper.PanicIfErrors(err)

	comment.Content = request.Content
	c.CommentRepository.Update(ctx, comment)
}

func (c *CommentServiceImpl) Delete(ctx context.Context, commentId int) {
	comment, err := c.CommentRepository.FindById(ctx, commentId)
	helper.PanicIfErrors(err)
	c.CommentRepository.Delete(ctx, comment.Id)
}

func (c *CommentServiceImpl) FindById(ctx context.Context, commentId int) response.CommentResponse {
	comment, err := c.CommentRepository.FindById(ctx, commentId)
	helper.PanicIfErrors(err)
	return response.CommentResponse(comment)
}

func (c *CommentServiceImpl) FindAll(ctx context.Context) []response.CommentResponse {
	comments := c.CommentRepository.FindAll(ctx)

	var commentResp []response.CommentResponse

	for _, value := range comments {
		comment := response.CommentResponse{Id: value.Id, PostId: value.PostId, UserId: value.UserId, Content: value.Content}
		commentResp = append(commentResp, comment)
	}
	return commentResp
}
