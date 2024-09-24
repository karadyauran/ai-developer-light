package main

import (
	"fmt"
	"log"
	"net/http"
	"./handlers"
	"./config"
)

func main() {
	config.LoadConfig()
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/tasks", handlers.TasksHandler)
	http.HandleFunc("/tasks/create", handlers.CreateTaskHandler)
	http.HandleFunc("/tasks/update", handlers.UpdateTaskHandler)
	http.HandleFunc("/tasks/delete", handlers.DeleteTaskHandler)

	fmt.Println("Server starting on port", config.PORT)
	log.Fatal(http.ListenAndServe(":"+config.PORT, nil))
}