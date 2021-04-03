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

		Tmpl.ExecuteTemplate(rw, "index.html", NewTaskViewModel(tasks, "", err.Error()))

		return
	}

	Tmpl.ExecuteTemplate(rw, "index.html", NewTaskViewModel(tasks, "", ""))
}

func Store(rw http.ResponseWriter, r *http.Request) {
	description := r.FormValue("description")

	tasks, err := models.GetTasks()

	if err != nil {
		fmt.Printf("could not get the tasks: %s\n", err.Error())

		Tmpl.ExecuteTemplate(rw, "index.html", NewTaskViewModel(tasks, description, err.Error()))

		return
	}

	task := models.NewTask(0, description, false)

	err = task.Save()

	if err != nil {
		fmt.Printf("could not save the task: %s\n", err.Error())

		Tmpl.ExecuteTemplate(rw, "index.html", NewTaskViewModel(tasks, description, err.Error()))

		return
	}

	http.Redirect(rw, r, "/", http.StatusMovedPermanently)
}
