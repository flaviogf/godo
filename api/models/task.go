package models

import "time"

type Task struct {
	ID          int64  `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func NewTask(id int64, description string, completed bool) *Task {
	return &Task{
		ID:          id,
		Description: description,
		Completed:   completed,
	}
}

func GetTasks() ([]*Task, error) {
	tasks := []*Task{}

	rows, err := DB.Query("SELECT id, description, completed FROM tasks")

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

func GetTask(id int64) (*Task, error) {
	row := DB.QueryRow("SELECT id, description, completed FROM tasks WHERE id = ?", id)

	err := row.Err()

	if err != nil {
		return nil, err
	}

	task := &Task{}

	err = row.Scan(&task.ID, &task.Description, &task.Completed)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (t *Task) MakeComplete() {
	t.Completed = true
}

func (t *Task) MakeIncomplete() {
	t.Completed = false
}

func (t *Task) Save() error {
	if t.ID <= 0 {
		return t.create()
	}

	return t.update()
}

func (t *Task) create() error {
	stmt, err := DB.Prepare("INSERT INTO tasks (description, completed, created_at, updated_at) VALUES (?, ?, ?, ?)")

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

func (t *Task) update() error {
	stmt, err := DB.Prepare("UPDATE tasks SET description = ?, completed = ?, updated_at = ? WHERE id = ?")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(t.Description, t.Completed, time.Now(), t.ID)

	if err != nil {
		return err
	}

	return nil
}
