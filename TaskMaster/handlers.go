package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"./models"
	"./db"
	"./utils"
)

func TasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tasks, err := db.GetTasks()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		utils.RespondWithJSON(w, http.StatusOK, tasks)
		return
	}

	http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var task models.Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := db.CreateTask(&task); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		utils.RespondWithJSON(w, http.StatusCreated, task)
	case "PUT":
		var task models.Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := db.UpdateTask(&task); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		utils.RespondWithJSON(w, http.StatusOK, task)
	case "DELETE":
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "Invalid task ID", http.StatusBadRequest)
			return
		}
		if err := db.DeleteTask(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		utils.RespondWithJSON(w, http.StatusNoContent, nil)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}