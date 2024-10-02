package router

import (
	"net/http"
	"github.com/gorilla/mux"
	"./handlers"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	r.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	r.HandleFunc("/tasks", handlers.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks", handlers.DeleteTask).Methods("DELETE")
	return r
}