package router

import (
	"net/http"
	"github.com/gorilla/mux"
	"./handlers"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	r.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", handlers.GetTask).Methods("GET")
	r.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")

	return r
}