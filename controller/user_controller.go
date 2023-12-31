package controller

import (
	"example/blog/data/request"
	"example/blog/data/response"
	"example/blog/helper"
	"example/blog/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (controller *UserController) Create(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	userCreateRequest := request.UserCreateRequest{}
	helper.ReadRequestBody(requests, &userCreateRequest)

	controller.UserService.Create(requests.Context(), userCreateRequest)
	webResponse := response.WebResponse{Code: 200, Status: "Ok", Data: nil}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *UserController) Update(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	userUpdateRequest := request.UserUpdateRequest{}
	helper.ReadRequestBody(requests, &userUpdateRequest)

	userId := params.ByName("userId")

	userUpdateRequest.Id = userId

	controller.UserService.Update(requests.Context(), userUpdateRequest)
	webResponse := response.WebResponse{Code: 200, Status: "Ok", Data: nil}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *UserController) Delete(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")

	controller.UserService.Delete(requests.Context(), userId)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *UserController) FindById(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")

	result := controller.UserService.FindById(requests.Context(), userId)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *UserController) FindAll(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	result := controller.UserService.FindAll(requests.Context())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}
