package routes

import (
	"net/http"

	"./handlers"
)

func RegisterRoutes() {
	http.HandleFunc("/tasks", handlers.TasksHandler)
	http.HandleFunc("/tasks/create", handlers.CreateTaskHandler)
	http.HandleFunc("/tasks/update", handlers.UpdateTaskHandler)
	http.HandleFunc("/tasks/delete", handlers.DeleteTaskHandler)
}