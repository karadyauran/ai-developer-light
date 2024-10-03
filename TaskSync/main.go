```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"./handlers"
	"./router"
	"./utils"
)

func main() {
	// Initialize the router
	r := router.InitRouter()

	// Initialize the database connection
	db, err := utils.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v\n", err)
	}
	defer db.Close()

	// Assign handlers with database dependency
	handlers.AssignHandlers(r, db)

	// Start the server
	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
```