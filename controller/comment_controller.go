package controller

import (
	"example/blog/data/request"
	"example/blog/data/response"
	"example/blog/helper"
	"example/blog/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type CommentController struct {
	CommentService service.CommentService
}

func NewCommentController(commentService service.CommentService) *CommentController {
	return &CommentController{CommentService: commentService}
}

func (controller *CommentController) Create(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	commentCreateRequest := request.CommentCreateRequest{}
	helper.ReadRequestBody(requests, &commentCreateRequest)

	controller.CommentService.Create(requests.Context(), commentCreateRequest)
	webResponse := response.WebResponse{Code: 200, Status: "Ok", Data: nil}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *CommentController) Update(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	commentUpdateRequest := request.CommentUpdateRequest{}
	helper.ReadRequestBody(requests, &commentUpdateRequest)

	commentId := params.ByName("commentId")

	commentUpdateRequest.Id = commentId

	controller.CommentService.Update(requests.Context(), commentUpdateRequest)
	webResponse := response.WebResponse{Code: 200, Status: "Ok", Data: nil}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *CommentController) Delete(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	commentId := params.ByName("commentId")

	controller.CommentService.Delete(requests.Context(), commentId)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *CommentController) FindById(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	commentId := params.ByName("commentId")

	result := controller.CommentService.FindById(requests.Context(), commentId)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *CommentController) FindAll(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	result := controller.CommentService.FindAll(requests.Context())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}
