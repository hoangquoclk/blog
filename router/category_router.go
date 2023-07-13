package router

import (
	"example/blog/controller"
	"github.com/julienschmidt/httprouter"
)

func CategoryRouter(controller controller.CategoryController, router *httprouter.Router) *httprouter.Router {
	router.GET("/categories", controller.FindAll)
	router.GET("/categories/:categoryId", controller.FindById)
	router.POST("/categories", controller.Create)
	router.PUT("/categories/:categoryId", controller.Update)
	router.DELETE("/categories/:categoryId", controller.Delete)

	return router
}
