package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/flaviogf/godo/web/models"
	"github.com/gorilla/mux"
)

func GetTasks(rw http.ResponseWriter, r *http.Request) {
	tasks, err := models.GetTasks()

	if err != nil {
		fmt.Printf("could not get the tasks: %s\n", err.Error())

		Tmpl.ExecuteTemplate(rw, "index.html", NewTaskViewModel(tasks, "", err.Error()))

		return
	}

	Tmpl.ExecuteTemplate(rw, "index.html", NewTaskViewModel(tasks, "", ""))
}

func CreateTask(rw http.ResponseWriter, r *http.Request) {
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

func MakeTaskComplete(rw http.ResponseWriter, r *http.Request) {
	tasks, err := models.GetTasks()

	if err != nil {
		fmt.Printf("could not get the tasks: %s\n", err.Error())

		Tmpl.ExecuteTemplate(rw, "index.html", NewTaskViewModel(tasks, "", err.Error()))

		return
	}

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Printf("could not parse the id: %s\n", err.Error())

		Tmpl.ExecuteTemplate(rw, "index.html", NewTaskViewModel(tasks, "", err.Error()))

		return
	}

	_, err = models.MakeTaskComplete(id)

	if err != nil {
		fmt.Printf("could not make the task complete: %s\n", err.Error())

		Tmpl.ExecuteTemplate(rw, "index.html", NewTaskViewModel(tasks, "", err.Error()))

		return
	}

	http.Redirect(rw, r, "/", http.StatusMovedPermanently)
}

func MakeTaskIncomplete(rw http.ResponseWriter, r *http.Request) {
	tasks, err := models.GetTasks()

	if err != nil {
		fmt.Printf("could not get the tasks: %s\n", err.Error())

		Tmpl.ExecuteTemplate(rw, "index.html", NewTaskViewModel(tasks, "", err.Error()))

		return
	}

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Printf("could not parse the id: %s\n", err.Error())

		Tmpl.ExecuteTemplate(rw, "index.html", NewTaskViewModel(tasks, "", err.Error()))

		return
	}

	_, err = models.MakeTaskIncomplete(id)

	if err != nil {
		fmt.Printf("could not make the task incomplete: %s\n", err.Error())

		Tmpl.ExecuteTemplate(rw, "index.html", NewTaskViewModel(tasks, "", err.Error()))

		return
	}

	http.Redirect(rw, r, "/", http.StatusMovedPermanently)
}
