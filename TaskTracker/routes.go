package routes

import (
    "net/http"
    "github.com/gorilla/mux"
    "project/handlers"
)

func InitRoutes() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
    router.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
    router.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")
    router.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")

    return router
}