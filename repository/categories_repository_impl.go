package repository

import (
	"context"
	"database/sql"
	"errors"
	"example/blog/helper"
	"example/blog/model"
	"github.com/google/uuid"
)

type CategoryRepositoryImpl struct {
	Db *sql.DB
}

func NewCategoryRepository(Db *sql.DB) CategoryRepository {
	return &CategoryRepositoryImpl{Db: Db}
}

// Save implement CategoriesRepository
func (c *CategoryRepositoryImpl) Save(ctx context.Context, category model.Category) {
	uuid := uuid.New()
	tx, err := c.Db.Begin()
	helper.PanicIfErrors(err)
	defer helper.CommitOrRollback(tx)

	SQL := "insert into categories(id, name) values (?, ?)"

	_, errQuery := tx.ExecContext(ctx, SQL, uuid, category.Name)

	helper.PanicIfErrors(errQuery)
}

// Update implements CategoriesRepository
func (c *CategoryRepositoryImpl) Update(ctx context.Context, category model.Category) {
	tx, err := c.Db.Begin()
	helper.PanicIfErrors(err)
	defer helper.CommitOrRollback(tx)

	SQL := "update categories set name=:name"

	_, errQuery := tx.ExecContext(ctx, SQL, category)
	helper.PanicIfErrors(errQuery)
}

// Delete implements CategoriesRepository
func (c *CategoryRepositoryImpl) Delete(ctx context.Context, categoryId string) {
	tx, err := c.Db.Begin()
	helper.PanicIfErrors(err)
	defer helper.CommitOrRollback(tx)

	SQL := "delete from categories where id=?"

	_, errQuery := tx.ExecContext(ctx, SQL, categoryId)
	helper.PanicIfErrors(errQuery)
}

// FindById implements CategoriesRepository
func (c *CategoryRepositoryImpl) FindById(ctx context.Context, categoryId string) (model.Category, error) {
	tx, err := c.Db.Begin()
	helper.PanicIfErrors(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select id, name from categories where id=?"
	result, errQuery := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfErrors(errQuery)
	defer result.Close()

	category := model.Category{}

	if result.Next() {
		err := result.Scan(&category.Id, &category.Name)
		helper.PanicIfErrors(err)
		return category, nil
	} else {
		return category, errors.New("category id not found")
	}
}

// FindAll implements CategoriesRepository
func (c *CategoryRepositoryImpl) FindAll(ctx context.Context) []model.Category {
	tx, err := c.Db.Begin()
	helper.PanicIfErrors(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select id, name from categories"
	result, errExec := tx.QueryContext(ctx, SQL)
	helper.PanicIfErrors(errExec)
	defer result.Close()

	var categories []model.Category

	for result.Next() {
		category := model.Category{}
		err := result.Scan(&category.Id, &category.Name)
		helper.PanicIfErrors(err)
		categories = append(categories, category)
	}
	return categories
}
