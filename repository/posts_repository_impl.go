package repository

import (
	"context"
	"database/sql"
	"errors"
	"example/blog/helper"
	"example/blog/model"
)

type PostRepositoryImpl struct {
	Db *sql.DB
}

func NewPostRepository(Db *sql.DB) PostRepository {
	return &PostRepositoryImpl{Db: Db}
}

// Save implement PostsRepository
func (p *PostRepositoryImpl) Save(ctx context.Context, post model.Post) {
	tx, err := p.Db.Begin()
	helper.PanicIfErrors(err)
	defer helper.CommitOrRollback(tx)

	SQL := "insert into posts(title, category_id, user_id, content) values($1, $2, $3, $4)"

	_, errQuery := tx.ExecContext(ctx, SQL, post.Title, post.CategoryId, post.UserId, post.Content)

	helper.PanicIfErrors(errQuery)
}

// Update implements PostsRepository
func (p *PostRepositoryImpl) Update(ctx context.Context, post model.Post) {
	tx, err := p.Db.Begin()
	helper.PanicIfErrors(err)
	defer helper.CommitOrRollback(tx)

	SQL := "update posts set title=:title, category_id=:category_id, user_id=:user_id, content=:content where id=:id"
	//SQL := "update users set username=$1, password=$2, email=$3 where id=$4"
	_, errQuery := tx.ExecContext(ctx, SQL, post)
	helper.PanicIfErrors(errQuery)
}

// Delete implements PostsRepository
func (p *PostRepositoryImpl) Delete(ctx context.Context, postId int) {
	tx, err := p.Db.Begin()
	helper.PanicIfErrors(err)
	defer helper.CommitOrRollback(tx)

	SQL := "delete from posts where id=$1"

	_, errQuery := tx.ExecContext(ctx, SQL, postId)
	helper.PanicIfErrors(errQuery)
}

// FindById implements PostsRepository
func (p *PostRepositoryImpl) FindById(ctx context.Context, postId int) (model.Post, error) {
	tx, err := p.Db.Begin()
	helper.PanicIfErrors(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select id, title, category_id, user_id, content from posts where id=$1"
	result, errExec := tx.QueryContext(ctx, SQL, postId)
	helper.PanicIfErrors(errExec)
	defer result.Close()

	post := model.Post{}

	if result.Next() {
		err := result.Scan(&post.Id, &post.Title, post.CategoryId, post.UserId, post.Content)
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

	SQL := "select id, title, category_id, user_id, content from posts"
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
