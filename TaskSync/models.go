```go
package models

import (
	"database/sql"
	"errors"
)

// Task represents a task entity
type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

// GetAllTasks retrieves all tasks from the database
func GetAllTasks(db *sql.DB) ([]Task, error) {
	rows, err := db.Query("SELECT id, title, description, completed FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

// CreateTask inserts a new task into the database
func CreateTask(db *sql.DB, task Task) error {
	result, err := db.Exec("INSERT INTO tasks (title, description, completed) VALUES (?, ?, ?)", task.Title, task.Description, task.Completed)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows affected")
	}

	return nil
}
```