package handlers

import (
	"html/template"

	"github.com/flaviogf/godo/web/models"
)

var Tmpl *template.Template

type TaskViewModel struct {
	Tasks       []*models.Task
	Today       string
	Description string
	Err         string
}

func NewTaskViewModel(tasks []*models.Task, today, description, err string) *TaskViewModel {
	return &TaskViewModel{
		Tasks:       tasks,
		Today:       today,
		Description: description,
		Err:         err,
	}
}
