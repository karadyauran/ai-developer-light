package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"./models"
)

var db *sql.DB

func InitDB() error {
	var err error
	db, err = sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		return err
	}

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		description TEXT,
		deadline TEXT,
		assignee TEXT,
		status TEXT
	);`
	_, err = db.Exec(createTableQuery)
	return err
}

func GetAllTasks() ([]models.Task, error) {
	rows, err := db.Query("SELECT id, title, description, deadline, assignee, status FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Deadline, &task.Assignee, &task.Status); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func CreateTask(task models.Task) error {
	_, err := db.Exec("INSERT INTO tasks (title, description, deadline, assignee, status) VALUES (?, ?, ?, ?, ?)",
		task.Title, task.Description, task.Deadline, task.Assignee, task.Status)
	return err
}

func UpdateTask(task models.Task) error {
	_, err := db.Exec("UPDATE tasks SET title = ?, description = ?, deadline = ?, assignee = ?, status = ? WHERE id = ?",
		task.Title, task.Description, task.Deadline, task.Assignee, task.Status, task.ID)
	return err
}

func DeleteTask(id int) error {
	_, err := db.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}