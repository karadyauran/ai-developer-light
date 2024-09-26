package main

import (
	"log"
	"net/http"

	"./routes"
	"./db"
)

func main() {
	// Initialize the database connection
	err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Register routes
	routes.RegisterRoutes()

	// Start the server
	log.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}