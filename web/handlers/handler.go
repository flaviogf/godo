package handlers

import (
	"html/template"

	"github.com/flaviogf/godo/web/models"
)

var Tmpl *template.Template

type TaskViewModel struct {
	Tasks       []*models.Task
	Description string
	Err         string
}

func NewTaskViewModel(tasks []*models.Task, description, err string) *TaskViewModel {
	return &TaskViewModel{
		Tasks:       tasks,
		Description: description,
		Err:         err,
	}
}
