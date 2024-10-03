```go
package reminder

import (
	"fmt"
	"time"
	"./task"
)

// ScheduleReminder sets up a reminder for a given task
func ScheduleReminder(t task.Task) {
	// Parse the deadline date
	deadline, err := time.Parse("2006-01-02", t.Deadline)
	if err != nil {
		fmt.Printf("Error parsing deadline: %v\n", err)
		return
	}

	// Calculate the time remaining until the deadline
	timeUntilDeadline := time.Until(deadline)

	// Set a reminder 1 day before the deadline
	reminderTime := timeUntilDeadline - 24*time.Hour

	if reminderTime > 0 {
		time.AfterFunc(reminderTime, func() {
			fmt.Printf("Reminder: Task %s is due in 24 hours!\n", t.Name)
		})
	} else {
		fmt.Printf("Task %s is already overdue!\n", t.Name)
	}
}
```