package main

import (
	"fmt"
	"log"
	"net/http"

	"./handlers"
)

func main() {
	http.HandleFunc("/tasks", handlers.TasksHandler)
	http.HandleFunc("/task", handlers.TaskHandler)

	fmt.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}