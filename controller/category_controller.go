package controller

import (
	"example/blog/data/request"
	"example/blog/data/response"
	"example/blog/helper"
	"example/blog/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type CategoryController struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) *CategoryController {
	return &CategoryController{CategoryService: categoryService}
}

func (controller *CategoryController) Create(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	categoryCreateRequest := request.CategoryCreateRequest{}
	helper.ReadRequestBody(requests, &categoryCreateRequest)

	controller.CategoryService.Create(requests.Context(), categoryCreateRequest)
	webResponse := response.WebResponse{Code: 200, Status: "Ok", Data: nil}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *CategoryController) Update(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	categoryUpdateRequest := request.CategoryUpdateRequest{}
	helper.ReadRequestBody(requests, &categoryUpdateRequest)

	categoryId := params.ByName("categoryId")

	id, err := strconv.Atoi(categoryId)

	helper.PanicIfErrors(err)
	categoryUpdateRequest.Id = id

	controller.CategoryService.Update(requests.Context(), categoryUpdateRequest)
	webResponse := response.WebResponse{Code: 200, Status: "Ok", Data: nil}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *CategoryController) Delete(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfErrors(err)

	controller.CategoryService.Delete(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *CategoryController) FindById(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfErrors(err)

	result := controller.CategoryService.FindById(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *CategoryController) FindAll(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	result := controller.CategoryService.FindAll(requests.Context())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}
