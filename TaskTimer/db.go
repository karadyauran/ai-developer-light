package db

import (
	"errors"
	"sync"
	"./models"
)

var (
	tasks   = []models.Task{}
	taskID  = 1
	taskMux sync.Mutex
)

func InitDB() {
	tasks = []models.Task{}
	taskID = 1
}

func GetAllTasks() []models.Task {
	return tasks
}

func CreateTask(task models.Task) {
	taskMux.Lock()
	defer taskMux.Unlock()
	task.ID = taskID
	taskID++
	tasks = append(tasks, task)
}

func GetTaskByID(id int) (models.Task, error) {
	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return models.Task{}, errors.New("task not found")
}

func UpdateTask(id int, updatedTask models.Task) error {
	taskMux.Lock()
	defer taskMux.Unlock()
	for i, task := range tasks {
		if task.ID == id {
			updatedTask.ID = id
			tasks[i] = updatedTask
			return nil
		}
	}
	return errors.New("task not found")
}

func DeleteTask(id int) error {
	taskMux.Lock()
	defer taskMux.Unlock()
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}