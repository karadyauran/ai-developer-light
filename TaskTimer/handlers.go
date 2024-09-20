package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"./models"
	"./db"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks := db.GetAllTasks()
	json.NewEncoder(w).Encode(tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	db.CreateTask(task)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	task, err := db.GetTaskByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Task not found")
		return
	}
	json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var updatedTask models.Task
	_ = json.NewDecoder(r.Body).Decode(&updatedTask)

	err := db.UpdateTask(id, updatedTask)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Task not found")
		return
	}
	json.NewEncoder(w).Encode(updatedTask)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	err := db.DeleteTask(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Task not found")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}