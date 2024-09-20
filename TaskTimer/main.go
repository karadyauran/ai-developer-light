package main

import (
	"log"
	"net/http"
	"./router"
	"./config"
	"./db"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Initialize database
	db.InitDB()

	// Set up router
	r := router.SetupRouter()

	// Start server
	log.Println("Starting server on port", config.Config.Port)
	if err := http.ListenAndServe(":"+config.Config.Port, r); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}