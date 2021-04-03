package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type Task struct {
	ID          int64  `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
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
