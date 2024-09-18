package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"io/ioutil"
	"strconv"
	"./models"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks := models.GetAllTasks()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &task)
	models.AddTask(task)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &task)
	updatedTask := models.UpdateTask(id, task)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTask)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	models.DeleteTask(id)
	w.WriteHeader(http.StatusNoContent)
}