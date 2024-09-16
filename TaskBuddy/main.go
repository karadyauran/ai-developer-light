package main

import (
	"log"
	"net/http"
	"./handlers"
	"./config"
)

func main() {
	config.LoadConfig()
	router := handlers.SetupRoutes()
	log.Fatal(http.ListenAndServe(":8080", router))
}