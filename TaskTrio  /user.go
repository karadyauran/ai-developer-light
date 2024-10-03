```go
package user

import (
	"fmt"
	"./task"
)

type User struct {
	Name  string
	Tasks []task.Task
}

// AssignTask assigns a task to a user
func AssignTask(userName string, newTask task.Task) {
	// Dummy user for the purpose of the example
	user := User{
		Name: userName,
	}

	// Assign the task to the user
	user.Tasks = append(user.Tasks, newTask)

	// Notify the user
	fmt.Printf("Task %s assigned to user %s\n", newTask.Name, user.Name)
}
```