```go
package main

import (
	"fmt"
	"net/http"
	"./task"
	"./user"
	"./reminder"
	"./router"
)

func main() {
	// Initialize the router
	r := router.NewRouter()

	// Define routes
	r.HandleFunc("/assign", func(w http.ResponseWriter, r *http.Request) {
		// Dummy data
		taskName := "Design Logo"
		assignee := "John Doe"
		deadline := "2023-11-30"

		// Create a new task
		newTask := task.NewTask(taskName, assignee, deadline)

		// Assign task to user
		user.AssignTask(assignee, newTask)

		// Schedule a reminder
		reminder.ScheduleReminder(newTask)

		fmt.Fprintf(w, "Task %s assigned to %s with deadline %s", newTask.Name, newTask.Assignee, newTask.Deadline)
	})

	r.HandleFunc("/progress", func(w http.ResponseWriter, r *http.Request) {
		// Update task progress
		taskName := "Design Logo"
		progress := "50%"

		updatedTask := task.UpdateProgress(taskName, progress)

		fmt.Fprintf(w, "Task %s updated to %s progress", updatedTask.Name, updatedTask.Progress)
	})

	// Start the server
	http.Handle("/", r)
	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}
```