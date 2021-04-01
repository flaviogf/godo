package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/flaviogf/godo/api/models"
	"github.com/go-playground/validator/v10"
)

func GetTasks(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	var tasks, err = models.GetTasks()

	encoder := json.NewEncoder(rw)

	if err != nil {
		log.Printf("something went wrong: %s\n", err.Error())

		rw.WriteHeader(http.StatusInternalServerError)

		encoder.Encode(Failure(err.Error()))

		return
	}

	encoder.Encode(Success(tasks))
}

func CreateTask(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	body := struct {
		Description string `json:"description" validate:"required"`
	}{}

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&body)

	encoder := json.NewEncoder(rw)

	if err != nil {
		log.Printf("cannot decode the request body: %s\n", err.Error())

		rw.WriteHeader(http.StatusInternalServerError)

		encoder.Encode(Failure(err.Error()))

		return
	}

	err = validator.New().Struct(&body)

	if err != nil {
		log.Printf("cannot validate the request body: %s", err.Error())

		rw.WriteHeader(http.StatusBadRequest)

		encoder.Encode(Failure(err.Error()))

		return
	}

	task := models.NewTask(body.Description)

	err = task.Save()

	if err != nil {
		log.Printf("something went wrong: %s\n", err.Error())

		rw.WriteHeader(http.StatusInternalServerError)

		encoder.Encode(Failure(err.Error()))

		return
	}

	encoder.Encode(Success(task))
}
