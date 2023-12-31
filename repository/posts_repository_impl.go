package repository

import (
	"context"
	"database/sql"
	"errors"
	"example/blog/helper"
	"example/blog/model"
	uuid2 "github.com/google/uuid"
)

type PostRepositoryImpl struct {
	Db *sql.DB
}

func NewPostRepository(Db *sql.DB) PostRepository {
	return &PostRepositoryImpl{Db: Db}
}

// Save implement PostsRepository
func (p *PostRepositoryImpl) Save(ctx context.Context, post model.Post) {
	uuid := uuid2.New()
	tx, err := p.Db.Begin()
	helper.PanicIfErrors(err)
	defer helper.CommitOrRollback(tx)

	SQL := "insert into posts(id, title, categoryId, userId, content) values(?, ?, ?, ?, ?)"

	_, errQuery := tx.ExecContext(ctx, SQL, uuid, post.Title, post.CategoryId, post.UserId, post.Content)

	helper.PanicIfErrors(errQuery)
}

// Update implements PostsRepository
func (p *PostRepositoryImpl) Update(ctx context.Context, post model.Post) {
	tx, err := p.Db.Begin()
	helper.PanicIfErrors(err)
	defer helper.CommitOrRollback(tx)

	SQL := "update posts set title=:title, categoryId=:categoryId, userId=:userId, content=:content where id=:id"

	_, errQuery := tx.ExecContext(ctx, SQL, post)
	helper.PanicIfErrors(errQuery)
}

// Delete implements PostsRepository
func (p *PostRepositoryImpl) Delete(ctx context.Context, postId string) {
	tx, err := p.Db.Begin()
	helper.PanicIfErrors(err)
	defer helper.CommitOrRollback(tx)

	SQL := "delete from posts where id=?"

	_, errQuery := tx.ExecContext(ctx, SQL, postId)
	helper.PanicIfErrors(errQuery)
}

// FindById implements PostsRepository
func (p *PostRepositoryImpl) FindById(ctx context.Context, postId string) (model.Post, error) {
	tx, err := p.Db.Begin()
	helper.PanicIfErrors(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select id, title, categoryId, userId, content from posts where id=?"
	result, errExec := tx.QueryContext(ctx, SQL, postId)
	helper.PanicIfErrors(errExec)
	defer result.Close()

	post := model.Post{}

	if result.Next() {
		err := result.Scan(&post.Id, &post.Title, &post.CategoryId, &post.UserId, &post.Content)
		helper.PanicIfErrors(err)
		return post, nil
	} else {
		return post, errors.New("post id not found")
	}
}

// FindAll implements PostsRepository
func (p *PostRepositoryImpl) FindAll(ctx context.Context) []model.Post {
	tx, err := p.Db.Begin()
	helper.PanicIfErrors(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select id, title, categoryId, userId, content from posts"
	result, errQuery := tx.QueryContext(ctx, SQL)
	helper.PanicIfErrors(errQuery)
	defer result.Close()

	var posts []model.Post

	for result.Next() {
		post := model.Post{}
		err := result.Scan(&post.Id, &post.Title, &post.CategoryId, &post.UserId, &post.Content)
		helper.PanicIfErrors(err)
		posts = append(posts, post)
	}
	return posts
}
