package models

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
)

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
	resp, err := http.Get(os.Getenv("GODO_API"))

	if err != nil {
		return []*Task{}, err
	}

	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return []*Task{}, err
	}

	body := struct {
		Data   []*Task  `json:"data"`
		Errors []string `json:"errors"`
	}{}

	err = json.Unmarshal(bytes, &body)

	if err != nil {
		return []*Task{}, err
	}

	return body.Data, nil
}

func (t *Task) Save() error {
	data, err := json.Marshal(t)

	if err != nil {
		return err
	}

	resp, err := http.Post(os.Getenv("GODO_API"), "application/json", bytes.NewReader(data))

	if err != nil {
		return err
	}

	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	body := struct {
		Data   *Task    `json:"data"`
		Errors []string `json:"errors"`
	}{}

	err = json.Unmarshal(bytes, &body)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("could not save the task")
	}

	return nil
}
