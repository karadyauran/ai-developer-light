package models

import (
	"sync"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

var (
	tasks   = []Task{}
	nextID  = 1
	mutex   = &sync.Mutex{}
)

func GetTasks() []Task {
	mutex.Lock()
	defer mutex.Unlock()
	return tasks
}

func AddTask(task Task) {
	mutex.Lock()
	defer mutex.Unlock()
	task.ID = nextID
	nextID++
	tasks = append(tasks, task)
}

func UpdateTask(updatedTask Task) {
	mutex.Lock()
	defer mutex.Unlock()
	for i, t := range tasks {
		if t.ID == updatedTask.ID {
			tasks[i] = updatedTask
			break
		}
	}
}

func DeleteTask(id int) {
	mutex.Lock()
	defer mutex.Unlock()
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}
}