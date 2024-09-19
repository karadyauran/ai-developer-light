package db

import (
	"errors"
	"sync"

	"./models"
)

var (
	tasks   = []models.Task{}
	nextID  = 1
	taskMux sync.Mutex
)

func GetTasks() ([]models.Task, error) {
	taskMux.Lock()
	defer taskMux.Unlock()
	return tasks, nil
}

func CreateTask(task *models.Task) error {
	taskMux.Lock()
	defer taskMux.Unlock()
	task.ID = nextID
	nextID++
	tasks = append(tasks, *task)
	return nil
}

func UpdateTask(updatedTask *models.Task) error {
	taskMux.Lock()
	defer taskMux.Unlock()
	for i, task := range tasks {
		if task.ID == updatedTask.ID {
			tasks[i] = *updatedTask
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