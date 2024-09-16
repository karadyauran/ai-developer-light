package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"io/ioutil"
	"./models"
	"./utils"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Tasks)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range models.Tasks {
		if item.ID == params["id"] {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.NotFound(w, r)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &task)
	task.ID = utils.GenerateID()
	models.Tasks = append(models.Tasks, task)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range models.Tasks {
		if item.ID == params["id"] {
			var updatedTask models.Task
			body, _ := ioutil.ReadAll(r.Body)
			_ = json.Unmarshal(body, &updatedTask)
			updatedTask.ID = item.ID
			models.Tasks[index] = updatedTask
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedTask)
			return
		}
	}
	http.NotFound(w, r)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range models.Tasks {
		if item.ID == params["id"] {
			models.Tasks = append(models.Tasks[:index], models.Tasks[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.NotFound(w, r)
}