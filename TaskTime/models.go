package models

type Task struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var tasks []Task
var nextID int

func init() {
	tasks = []Task{}
	nextID = 1
}

func GetAllTasks() []Task {
	return tasks
}

func AddTask(task Task) {
	task.ID = nextID
	nextID++
	tasks = append(tasks, task)
}

func UpdateTask(id int, updatedTask Task) *Task {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Title = updatedTask.Title
			tasks[i].Status = updatedTask.Status
			return &tasks[i]
		}
	}
	return nil
}

func DeleteTask(id int) bool {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return true
		}
	}
	return false
}