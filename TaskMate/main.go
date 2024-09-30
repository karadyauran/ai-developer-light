package main

import (
	"log"
	"net/http"

	"./config"
	"./db"
	"./handlers"
	"./routes"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Initialize database
	db.InitDB()

	// Initialize routes
	router := routes.InitRoutes()

	// Start the server
	log.Println("Server running on port", config.Config.Port)
	if err := http.ListenAndServe(":"+config.Config.Port, router); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}