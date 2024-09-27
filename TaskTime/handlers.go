package handlers

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
	"strconv"
	"github.com/gorilla/mux"
	"./models"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks := models.GetAllTasks()
	json.NewEncoder(w).Encode(tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &task)
	models.CreateTask(task)
	w.WriteHeader(http.StatusCreated)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var task models.Task
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &task)
	models.UpdateTask(id, task)
	w.WriteHeader(http.StatusOK)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	models.DeleteTask(id)
	w.WriteHeader(http.StatusNoContent)
}

func GetTimeLogs(w http.ResponseWriter, r *http.Request) {
	timeLogs := models.GetAllTimeLogs()
	json.NewEncoder(w).Encode(timeLogs)
}

func CreateTimeLog(w http.ResponseWriter, r *http.Request) {
	var timeLog models.TimeLog
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &timeLog)
	models.CreateTimeLog(timeLog)
	w.WriteHeader(http.StatusCreated)
}

func UpdateTimeLog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var timeLog models.TimeLog
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &timeLog)
	models.UpdateTimeLog(id, timeLog)
	w.WriteHeader(http.StatusOK)
}

func DeleteTimeLog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	models.DeleteTimeLog(id)
	w.WriteHeader(http.StatusNoContent)
}