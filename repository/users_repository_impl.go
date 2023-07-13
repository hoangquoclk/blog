package repository

import (
	"context"
	"database/sql"
	"errors"
	"example/blog/helper"
	"example/blog/model"
)

type UserRepositoryImpl struct {
	Db *sql.DB
}

func NewUserRepository(Db *sql.DB) UserRepository {
	return &UserRepositoryImpl{Db: Db}
}

// Save implement UsersRepository
func (u *UserRepositoryImpl) Save(ctx context.Context, user model.User) {
	tx, err := u.Db.Begin()
	helper.PanicIfErrors(err)
	defer helper.CommitOrRollback(tx)

	SQL := "insert into users(username, password, email) values (?, ?, ?)"

	_, errQuery := tx.ExecContext(ctx, SQL, user.Username, user.Password, user.Email)

	helper.PanicIfErrors(errQuery)
}

// Update implements UsersRepository
func (u *UserRepositoryImpl) Update(ctx context.Context, user model.User) {
	tx, err := u.Db.Begin()
	helper.PanicIfErrors(err)
	defer helper.CommitOrRollback(tx)

	SQL := "update users set username=:username, password=:password, email=:email where id=:id"
	//SQL := "update users set username=$1, password=$2, email=$3 where id=$4"
	_, errQuery := tx.ExecContext(ctx, SQL, user)
	helper.PanicIfErrors(errQuery)
}

// Delete implements UsersRepository
func (u *UserRepositoryImpl) Delete(ctx context.Context, userId int) {
	tx, err := u.Db.Begin()
	helper.PanicIfErrors(err)
	defer helper.CommitOrRollback(tx)

	SQL := "delete from users where id=?"

	_, errQuery := tx.ExecContext(ctx, SQL, userId)
	helper.PanicIfErrors(errQuery)
}

// FindById implements UsersRepository
func (u *UserRepositoryImpl) FindById(ctx context.Context, userId int) (model.User, error) {
	tx, err := u.Db.Begin()
	helper.PanicIfErrors(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select id, username from users where id=?"
	result, errExec := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfErrors(errExec)
	defer result.Close()

	user := model.User{}

	if result.Next() {
		err := result.Scan(&user.Id, &user.Username)
		helper.PanicIfErrors(err)
		return user, nil
	} else {
		return user, errors.New("user id not found")
	}
}

// FindAll implements UsersRepository
func (u *UserRepositoryImpl) FindAll(ctx context.Context) []model.User {
	tx, err := u.Db.Begin()
	helper.PanicIfErrors(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select id, username from users"
	result, errQuery := tx.QueryContext(ctx, SQL)
	helper.PanicIfErrors(errQuery)
	defer result.Close()

	var users []model.User

	for result.Next() {
		user := model.User{}
		err := result.Scan(&user.Id, &user.Username)
		helper.PanicIfErrors(err)
		users = append(users, user)
	}
	return users
}
