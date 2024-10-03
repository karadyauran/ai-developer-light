```go
package task

type Task struct {
	Name      string
	Assignee  string
	Deadline  string
	Progress  string
}

// NewTask creates a new task with the given details
func NewTask(name, assignee, deadline string) Task {
	return Task{
		Name:     name,
		Assignee: assignee,
		Deadline: deadline,
		Progress: "0%",
	}
}

// UpdateProgress updates the progress of a given task
func UpdateProgress(taskName, progress string) Task {
	// Dummy implementation for updating progress
	// In a real application, this function would search a database or a list
	updatedTask := Task{
		Name:     taskName,
		Progress: progress,
	}
	return updatedTask
}
```