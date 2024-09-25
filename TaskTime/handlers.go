package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"./models"
	"./utils"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks := models.GetAllTasks()
	utils.RespondWithJSON(w, http.StatusOK, tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	models.AddTask(task)
	utils.RespondWithJSON(w, http.StatusCreated, task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var task models.Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	updatedTask := models.UpdateTask(id, task)
	if updatedTask == nil {
		utils.RespondWithJSON(w, http.StatusNotFound, "Task not found")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, updatedTask)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	success := models.DeleteTask(id)
	if !success {
		utils.RespondWithJSON(w, http.StatusNotFound, "Task not found")
		return
	}
	utils.RespondWithJSON(w, http.StatusNoContent, nil)
}