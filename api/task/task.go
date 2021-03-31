package task

import (
	"time"

	"github.com/flaviogf/godo/api/data"
)

type Task struct {
	ID          int64  `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func NewTask(description string) *Task {
	return &Task{
		ID:          0,
		Description: description,
		Completed:   false,
	}
}

func GetTasks() ([]*Task, error) {
	tasks := []*Task{}

	rows, err := data.DB.Query("SELECT id, description, completed FROM tasks")

	if err != nil {
		return tasks, err
	}

	for rows.Next() {
		task := &Task{}

		rows.Scan(&task.ID, &task.Description, &task.Completed)

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (t *Task) Save() error {
	stmt, err := data.DB.Prepare("INSERT INTO tasks (description, completed, created_at, updated_at) VALUES (?, ?, ?, ?)")

	if err != nil {
		return err
	}

	now := time.Now()

	result, err := stmt.Exec(t.Description, t.Completed, now, now)

	if err != nil {
		return err
	}

	t.ID, err = result.LastInsertId()

	if err != nil {
		return err
	}

	return nil
}
