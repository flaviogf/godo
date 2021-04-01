package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/flaviogf/godo/api/models"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
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
		log.Printf("could not decode the request body: %s\n", err.Error())

		rw.WriteHeader(http.StatusInternalServerError)

		encoder.Encode(Failure(err.Error()))

		return
	}

	err = validator.New().Struct(&body)

	if err != nil {
		log.Printf("could not validate the request body: %s\n", err.Error())

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

func GetTask(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)

	encoder := json.NewEncoder(rw)

	if err != nil {
		log.Printf("could not parse the id: %s\n", err.Error())

		rw.WriteHeader(http.StatusBadRequest)

		encoder.Encode(Failure(err.Error()))

		return
	}

	task, err := models.GetTask(id)

	if err != nil {
		log.Printf("could not found the task: %s\n", err.Error())

		rw.WriteHeader(http.StatusNotFound)

		encoder.Encode(Failure(err.Error()))

		return
	}

	encoder.Encode(Success(task))
}

func MakeTaskComplete(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)

	encoder := json.NewEncoder(rw)

	if err != nil {
		log.Printf("could not parse the id: %s\n", err.Error())

		rw.WriteHeader(http.StatusBadRequest)

		encoder.Encode(Failure(err.Error()))

		return
	}

	task, err := models.GetTask(id)

	if err != nil {
		log.Printf("could not found the task: %s\n", err.Error())

		rw.WriteHeader(http.StatusNotFound)

		encoder.Encode(Failure(err.Error()))

		return
	}

	task.MakeComplete()

	err = task.Save()

	if err != nil {
		log.Printf("could not save the task: %s\n", err.Error())

		rw.WriteHeader(http.StatusInternalServerError)

		encoder.Encode(Failure(err.Error()))

		return
	}

	encoder.Encode(Success(task))
}

func MakeTaskIncomplete(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)

	encoder := json.NewEncoder(rw)

	if err != nil {
		log.Printf("could not parse the id: %s\n", err.Error())

		rw.WriteHeader(http.StatusBadRequest)

		encoder.Encode(Failure(err.Error()))

		return
	}

	task, err := models.GetTask(id)

	if err != nil {
		log.Printf("could not found the task: %s\n", err.Error())

		rw.WriteHeader(http.StatusNotFound)

		encoder.Encode(Failure(err.Error()))

		return
	}

	task.MakeIncomplete()

	err = task.Save()

	if err != nil {
		log.Printf("could not save the task: %s\n", err.Error())

		rw.WriteHeader(http.StatusInternalServerError)

		encoder.Encode(Failure(err.Error()))

		return
	}

	encoder.Encode(Success(task))
}
