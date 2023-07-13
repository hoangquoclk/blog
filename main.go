package main

import (
	"example/blog/config"
	"example/blog/controller"
	"example/blog/helper"
	"example/blog/repository"
	"example/blog/router"
	"example/blog/service"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	fmt.Printf("Start server")

	// database
	db := config.DatabaseConnection()

	// repository
	userRepository := repository.NewUserRepository(db)
	categoryRepository := repository.NewCategoryRepository(db)
	postRepository := repository.NewPostRepository(db)
	commentRepository := repository.NewCommentRepository(db)

	// service
	userService := service.NewUserServiceImpl(userRepository)
	categoryService := service.NewCategoryServiceImpl(categoryRepository)
	postService := service.NewPostServiceImpl(postRepository)
	commentService := service.NewCommentServiceImpl(commentRepository)

	// controller
	userController := controller.UserController{UserService: userService}
	categoryController := controller.CategoryController{CategoryService: categoryService}
	postController := controller.PostController{PostService: postService}
	commentController := controller.CommentController{CommentService: commentService}

	// router
	newRoute := httprouter.New()

	newRoute = router.UserRouter(userController, newRoute)
	newRoute = router.CategoryRouter(categoryController, newRoute)
	newRoute = router.PostRouter(postController, newRoute)
	newRoute = router.CommentRouter(commentController, newRoute)

	newRoute.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Welcome Home")
	})

	server := http.Server{Addr: "localhost:8080", Handler: newRoute}

	err := server.ListenAndServe()

	helper.PanicIfErrors(err)
}
