package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"./handlers"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	router.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	router.HandleFunc("/tasks", handlers.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks", handlers.DeleteTask).Methods("DELETE")

	return router
}