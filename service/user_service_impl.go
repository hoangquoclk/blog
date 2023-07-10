package service

import (
	"context"
	"example/blog/data/request"
	"example/blog/data/response"
	"example/blog/helper"
	"example/blog/model"
	"example/blog/repository"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserServiceImpl(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{UserRepository: userRepository}
}

func (u *UserServiceImpl) Create(ctx context.Context, request request.UserCreateRequest) {
	user := model.User{
		Username: request.Username,
		Password: request.Password,
		Email:    request.Email,
	}
	u.UserRepository.Save(ctx, user)
}

func (u *UserServiceImpl) Update(ctx context.Context, request request.UserUpdateRequest) {
	user, err := u.UserRepository.FindById(ctx, request.Id)
	helper.PanicIfErrors(err)

	user.Username = request.Username
	user.Password = request.Password
	user.Email = request.Email
	u.UserRepository.Update(ctx, user)
}

func (u *UserServiceImpl) Delete(ctx context.Context, userId int) {
	user, err := u.UserRepository.FindById(ctx, userId)
	helper.PanicIfErrors(err)
	u.UserRepository.Delete(ctx, user.Id)
}

func (u *UserServiceImpl) FindById(ctx context.Context, userId int) response.UserResponse {
	user, err := u.UserRepository.FindById(ctx, userId)
	helper.PanicIfErrors(err)
	return response.UserResponse(user)
}

func (u *UserServiceImpl) FindAll(ctx context.Context) []response.UserResponse {
	users := u.UserRepository.FindAll(ctx)

	var userResp []response.UserResponse

	for _, value := range users {
		user := response.UserResponse{Id: value.Id, Username: value.Username, Password: value.Password, Email: value.Email}
		userResp = append(userResp, user)
	}
	return userResp
}
