package models

type Task struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Status      string `json:"status"`
}

var tasks []Task
var idCounter int

func init() {
    idCounter = 1
    tasks = []Task{}
}

func GetAllTasks() []Task {
    return tasks
}

func AddTask(task Task) {
    task.ID = idCounter
    idCounter++
    tasks = append(tasks, task)
}

func UpdateTask(id int, updatedTask Task) {
    for i, task := range tasks {
        if task.ID == id {
            tasks[i].Title = updatedTask.Title
            tasks[i].Description = updatedTask.Description
            tasks[i].Status = updatedTask.Status
            break
        }
    }
}

func DeleteTask(id int) {
    for i, task := range tasks {
        if task.ID == id {
            tasks = append(tasks[:i], tasks[i+1:]...)
            break
        }
    }
}