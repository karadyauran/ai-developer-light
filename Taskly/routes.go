package routes

import (
    "github.com/gorilla/mux"
    "../handlers"
)

func RegisterRoutes(r *mux.Router) {
    r.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
    r.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
    r.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")
    r.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")
}