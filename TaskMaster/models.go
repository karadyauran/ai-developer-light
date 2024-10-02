package models

import (
	"errors"
	"sync"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Assignee    string `json:"assignee"`
	Deadline    string `json:"deadline"`
	Status      string `json:"status"`
}

var (
	tasks  = make(map[int]Task)
	nextID = 1
	mu     sync.Mutex
)

func AddTask(task Task) {
	mu.Lock()
	defer mu.Unlock()
	task.ID = nextID
	nextID++
	tasks[task.ID] = task
}

func GetAllTasks() []Task {
	mu.Lock()
	defer mu.Unlock()
	taskList := make([]Task, 0, len(tasks))
	for _, task := range tasks {
		taskList = append(taskList, task)
	}
	return taskList
}

func UpdateTask(updatedTask Task) error {
	mu.Lock()
	defer mu.Unlock()
	if _, exists := tasks[updatedTask.ID]; !exists {
		return errors.New("task not found")
	}
	tasks[updatedTask.ID] = updatedTask
	return nil
}

func DeleteTask(id int) error {
	mu.Lock()
	defer mu.Unlock()
	if _, exists := tasks[id]; !exists {
		return errors.New("task not found")
	}
	delete(tasks, id)
	return nil
}