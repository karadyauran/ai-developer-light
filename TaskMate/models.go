package models

import (
	"errors"

	"./db"
)

type Task struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Priority string `json:"priority"`
	Deadline string `json:"deadline"`
	Status   string `json:"status"`
}

func GetAllTasks() ([]Task, error) {
	var tasks []Task
	rows, err := db.DB.Query("SELECT id, title, priority, deadline, status FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Priority, &task.Deadline, &task.Status); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func CreateTask(task *Task) error {
	stmt, err := db.DB.Prepare("INSERT INTO tasks(title, priority, deadline, status) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(task.Title, task.Priority, task.Deadline, task.Status)
	return err
}

func UpdateTask(task *Task) error {
	stmt, err := db.DB.Prepare("UPDATE tasks SET title=?, priority=?, deadline=?, status=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(task.Title, task.Priority, task.Deadline, task.Status, task.ID)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return errors.New("no rows affected")
	}

	return nil
}

func DeleteTask(task *Task) error {
	stmt, err := db.DB.Prepare("DELETE FROM tasks WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(task.ID)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return errors.New("no rows affected")
	}

	return nil
}