package routes

import (
	"encoding/json"
	"net/http"

	"github.com/fakeblackspectre/api-postgres/db"
	"github.com/fakeblackspectre/api-postgres/models"
	"github.com/gorilla/mux"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	createdTask := db.DB.Create(&task)
	err := createdTask.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode((&task))
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)
	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound) //404
		w.Write([]byte("Task not found"))
		return
	}

	json.NewEncoder(w).Encode(&task)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)
	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound) //404
		w.Write([]byte("Task not found"))
		return
	}

	// db.DB.Delete(&user) // change status to deleted, adding the delete date to the user
	db.DB.Unscoped().Delete(&task) // remove from table
	w.WriteHeader(http.StatusOK)
}