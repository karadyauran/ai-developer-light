package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"./handlers"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	r.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")

	r.HandleFunc("/time-logs", handlers.GetTimeLogs).Methods("GET")
	r.HandleFunc("/time-logs", handlers.CreateTimeLog).Methods("POST")
	r.HandleFunc("/time-logs/{id}", handlers.UpdateTimeLog).Methods("PUT")
	r.HandleFunc("/time-logs/{id}", handlers.DeleteTimeLog).Methods("DELETE")

	return r
}