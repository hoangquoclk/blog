package router

import (
	"example/blog/controller"
	"github.com/julienschmidt/httprouter"
)

func PostRouter(controller controller.PostController, router *httprouter.Router) *httprouter.Router {
	router.GET("/posts", controller.FindAll)
	router.GET("/posts/:postId", controller.FindById)
	router.POST("/posts", controller.Create)
	router.PUT("/posts/:postId", controller.Update)
	router.DELETE("/posts/:postId", controller.Delete)

	return router
}
