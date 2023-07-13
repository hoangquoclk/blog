package router

import (
	"example/blog/controller"
	"github.com/julienschmidt/httprouter"
)

func CommentRouter(controller controller.CommentController, router *httprouter.Router) *httprouter.Router {
	router.GET("/comments", controller.FindAll)
	router.GET("/comments/:commentId", controller.FindById)
	router.POST("/comments", controller.Create)
	router.PUT("/comments/:commentId", controller.Update)
	router.DELETE("/comments/:commentId", controller.Delete)

	return router
}
