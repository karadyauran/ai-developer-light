package routes

import (
	"net/http"
	"./controllers"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/tasks", controllers.GetTasks).Methods("GET")
	router.HandleFunc("/tasks", controllers.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", controllers.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", controllers.DeleteTask).Methods("DELETE")

	return router
}