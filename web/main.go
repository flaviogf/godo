package main

import (
	"html/template"
	"net/http"
	"os"

	"github.com/flaviogf/godo/web/handlers"
	"github.com/gorilla/mux"
)

func main() {
	tmpl := template.Must(template.ParseGlob("./templates/*.html"))

	handlers.Tmpl = tmpl

	r := mux.NewRouter()

	r.HandleFunc("/", handlers.Index).Methods("GET")

	r.HandleFunc("/", handlers.Store).Methods("POST")

	http.Handle("/", r)

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	http.ListenAndServe(os.Getenv("GODO_ADDR"), nil)
}
