package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/flaviogf/godo/api/handlers"
	"github.com/flaviogf/godo/api/models"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "godo.sqlite3")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	models.DB = db

	r := mux.NewRouter()

	r.HandleFunc("/", handlers.GetTasks).Methods("GET")

	r.HandleFunc("/", handlers.CreateTask).Methods("POST")

	http.Handle("/", r)

	http.Handle("/docs", middleware.SwaggerUI(middleware.SwaggerUIOpts{SpecURL: "/swagger.yml"}, nil))

	http.Handle("/swagger.yml", http.FileServer(http.Dir("./")))

	http.ListenAndServe(":3000", nil)
}
