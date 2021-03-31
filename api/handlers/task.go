package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/flaviogf/godo/api/task"
	"github.com/go-playground/validator/v10"
)

func GetTasks(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	var results, err = task.GetTasks()

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)

		log.Printf("something went wrong: %s\n", err.Error())

		return
	}

	encoder := json.NewEncoder(rw)

	encoder.Encode(results)
}

func CreateTask(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	body := struct {
		Description string `json:"description" validate:"required"`
	}{}

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&body)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)

		log.Printf("cannot decode the request body: %s\n", err.Error())

		return
	}

	err = validator.New().Struct(&body)

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)

		log.Printf("cannot validate the request body: %s", err.Error())

		return
	}

	result := task.NewTask(body.Description)

	err = result.Save()

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)

		log.Printf("something went wrong: %s\n", err.Error())

		return
	}

	encoder := json.NewEncoder(rw)

	encoder.Encode(result)
}
