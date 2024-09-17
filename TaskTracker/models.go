package models

import (
    "errors"
)

type Task struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Status      string `json:"status"`
}

var tasks = []Task{
    {ID: 1, Title: "Task 1", Description: "Description 1", Status: "pending"},
    {ID: 2, Title: "Task 2", Description: "Description 2", Status: "completed"},
}

func GetAllTasks() []Task {
    return tasks
}

func AddTask(task Task) {
    task.ID = len(tasks) + 1
    tasks = append(tasks, task)
}

func UpdateTask(id int, updatedTask Task) (Task, error) {
    for i, task := range tasks {
        if task.ID == id {
            tasks[i] = updatedTask
            tasks[i].ID = id
            return tasks[i], nil
        }
    }
    return Task{}, errors.New("task not found")
}

func DeleteTask(id int) error {
    for i, task := range tasks {
        if task.ID == id {
            tasks = append(tasks[:i], tasks[i+1:]...)
            return nil
        }
    }
    return errors.New("task not found")
}