package main

import (
	"example/blog/helper"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	fmt.Printf("Start server")

	routes := httprouter.New()

	routes.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Welcome Home")
	})

	server := http.Server{Addr: "localhost:8080", Handler: routes}

	err := server.ListenAndServe()

	helper.PanicIfErrors(err)
}
