package router

import (
	"example/blog/controller"
	"github.com/julienschmidt/httprouter"
)

func UserRouter(controller controller.UserController, router *httprouter.Router) *httprouter.Router {
	//router := httprouter.New()

	router.GET("/users", controller.FindAll)
	router.GET("/users/:userId", controller.FindById)
	router.POST("/users", controller.Create)
	router.PUT("/users/:userId", controller.Update)
	router.DELETE("/users/:userId", controller.Delete)

	return router
}
