package main

import (
	"log"
	"net/http"
	"./handlers"
	"./routes"
)

func main() {
	router := routes.NewRouter()
	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}