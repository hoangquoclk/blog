package repository

import (
	"context"
	"database/sql"
	"errors"
	"example/blog/helper"
	"example/blog/model"
	uuid2 "github.com/google/uuid"
)

type CommentRepositoryImpl struct {
	Db *sql.DB
}

func NewCommentRepository(Db *sql.DB) CommentRepository {
	return &CommentRepositoryImpl{Db: Db}
}

// Save implement CommentsRepository
func (c *CommentRepositoryImpl) Save(ctx context.Context, comment model.Comment) {
	uuid := uuid2.New()
	tx, err := c.Db.Begin()
	helper.PanicIfErrors(err)
	defer helper.CommitOrRollback(tx)

	SQL := "insert into comments(id, postId, userId, content) values(?, ?, ? ,?)"

	_, errQuery := tx.ExecContext(ctx, SQL, uuid, comment.PostId, comment.UserId, comment.Content)

	helper.PanicIfErrors(errQuery)
}

// Update implements CommentsRepository
func (c *CommentRepositoryImpl) Update(ctx context.Context, comment model.Comment) {
	tx, err := c.Db.Begin()
	helper.PanicIfErrors(err)
	defer helper.CommitOrRollback(tx)

	SQL := "update comments set postId=:postId, userId=:userId, content=:content where id=:id"

	_, errQuery := tx.ExecContext(ctx, SQL, comment)
	helper.PanicIfErrors(errQuery)
}

// Delete implements CommentsRepository
func (c *CommentRepositoryImpl) Delete(ctx context.Context, commentId string) {
	tx, err := c.Db.Begin()
	helper.PanicIfErrors(err)
	defer helper.CommitOrRollback(tx)

	SQL := "delete from comments where id=?"

	_, errQuery := tx.ExecContext(ctx, SQL, commentId)
	helper.PanicIfErrors(errQuery)
}

// FindById implements CommentsRepository
func (c *CommentRepositoryImpl) FindById(ctx context.Context, commentId string) (model.Comment, error) {
	tx, err := c.Db.Begin()
	helper.PanicIfErrors(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select id, postId, userId, content from comments where id=?"
	result, errExec := tx.QueryContext(ctx, SQL, commentId)
	helper.PanicIfErrors(errExec)
	defer result.Close()

	comment := model.Comment{}

	if result.Next() {
		err := result.Scan(&comment.Id, &comment.PostId, &comment.UserId, &comment.Content)
		helper.PanicIfErrors(err)
		return comment, nil
	} else {
		return comment, errors.New("comment id not found")
	}
}

// FindAll implements UsersRepository
func (c *CommentRepositoryImpl) FindAll(ctx context.Context) []model.Comment {
	tx, err := c.Db.Begin()
	helper.PanicIfErrors(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select id, postId, userId, content from comments"
	result, errQuery := tx.QueryContext(ctx, SQL)
	helper.PanicIfErrors(errQuery)
	defer result.Close()

	var comments []model.Comment

	for result.Next() {
		comment := model.Comment{}
		err := result.Scan(&comment.Id, &comment.PostId, &comment.UserId, &comment.Content)
		helper.PanicIfErrors(err)
		comments = append(comments, comment)
	}
	return comments
}
