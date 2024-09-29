package main

import (
	"fmt"
	"log"
	"net/http"
	"./handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/tasks", handlers.TasksHandler)
	http.HandleFunc("/create_task", handlers.CreateTaskHandler)

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}