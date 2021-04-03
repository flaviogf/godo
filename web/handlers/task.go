package handlers

import (
	"fmt"
	"net/http"

	"github.com/flaviogf/godo/web/models"
)

func Index(rw http.ResponseWriter, r *http.Request) {
	tasks, err := models.GetTasks()

	if err != nil {
		fmt.Printf("could not get the tasks: %s\n", err.Error())

		Tmpl.ExecuteTemplate(rw, "index.html", Failure(err.Error()))

		return
	}

	Tmpl.ExecuteTemplate(rw, "index.html", Success(tasks))
}
