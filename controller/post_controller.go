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

type PostController struct {
	PostService service.PostService
}

func NewPostController(postService service.PostService) *PostController {
	return &PostController{PostService: postService}
}

func (controller *PostController) Create(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	postCreateRequest := request.PostCreateRequest{}
	helper.ReadRequestBody(requests, &postCreateRequest)

	controller.PostService.Create(requests.Context(), postCreateRequest)
	webResponse := response.WebResponse{Code: 200, Status: "Ok", Data: nil}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *PostController) Update(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	postUpdateRequest := request.PostUpdateRequest{}
	helper.ReadRequestBody(requests, &postUpdateRequest)

	postId := params.ByName("postId")

	id, err := strconv.Atoi(postId)

	helper.PanicIfErrors(err)
	postUpdateRequest.Id = id

	controller.PostService.Update(requests.Context(), postUpdateRequest)
	webResponse := response.WebResponse{Code: 200, Status: "Ok", Data: nil}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *PostController) Delete(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	postId := params.ByName("postId")
	id, err := strconv.Atoi(postId)
	helper.PanicIfErrors(err)

	controller.PostService.Delete(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *PostController) FindById(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	postId := params.ByName("postId")
	id, err := strconv.Atoi(postId)
	helper.PanicIfErrors(err)

	result := controller.PostService.FindById(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *PostController) FindAll(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	result := controller.PostService.FindAll(requests.Context())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}
