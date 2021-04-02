package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/flaviogf/godo/api/handlers"
	"github.com/flaviogf/godo/api/models"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", os.Getenv("GODO_DATABASE"))

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	models.DB = db

	r := mux.NewRouter()

	r.HandleFunc("/", handlers.GetTasks).Methods("GET")

	r.HandleFunc("/", handlers.CreateTask).Methods("POST")

	r.HandleFunc("/{id}", handlers.GetTask).Methods("GET")

	r.HandleFunc("/{id}", handlers.UpdateTask).Methods("PUT")

	r.HandleFunc("/{id}/completed", handlers.MakeTaskComplete).Methods("POST")

	r.HandleFunc("/{id}/completed", handlers.MakeTaskIncomplete).Methods("DELETE")

	http.Handle("/", r)

	http.Handle("/docs", middleware.SwaggerUI(middleware.SwaggerUIOpts{SpecURL: "/swagger.yml"}, nil))

	http.Handle("/swagger.yml", http.FileServer(http.Dir("./")))

	http.ListenAndServe(os.Getenv("GODO_ADDR"), nil)
}
