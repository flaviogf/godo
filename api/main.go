package main

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	http.Handle("/", r)

	http.Handle("/docs", middleware.SwaggerUI(middleware.SwaggerUIOpts{SpecURL: "/swagger.yml"}, nil))

	http.Handle("/swagger.yml", http.FileServer(http.Dir("./")))

	http.ListenAndServe(":3000", nil)
}
