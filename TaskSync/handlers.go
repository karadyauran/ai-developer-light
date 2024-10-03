```go
package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"./models"
)

// TaskHandler struct to hold dependencies
type TaskHandler struct {
	DB *sql.DB
}

// AssignHandlers associates handlers with routes
func AssignHandlers(r *http.ServeMux, db *sql.DB) {
	handler := &TaskHandler{DB: db}
	r.HandleFunc("/tasks", handler.GetTasks)
	r.HandleFunc("/task", handler.CreateTask)
}

// GetTasks handles the retrieval of tasks
func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := models.GetAllTasks(h.DB)
	if err != nil {
		http.Error(w, "Error fetching tasks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// CreateTask handles the creation of a new task
func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid task data", http.StatusBadRequest)
		return
	}

	if err := models.CreateTask(h.DB, task); err != nil {
		http.Error(w, "Error creating task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
```