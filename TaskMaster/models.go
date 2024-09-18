package models

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
	Completed   bool   `json:"completed"`
}

var tasks = []Task{
	{ID: 1, Title: "Sample Task 1", Description: "This is a sample task", Priority: "High", Completed: false},
	{ID: 2, Title: "Sample Task 2", Description: "This is another sample task", Priority: "Medium", Completed: false},
}

func GetAllTasks() []Task {
	return tasks
}

func AddTask(task Task) {
	task.ID = len(tasks) + 1
	tasks = append(tasks, task)
}

func UpdateTask(id int, updatedTask Task) Task {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Title = updatedTask.Title
			tasks[i].Description = updatedTask.Description
			tasks[i].Priority = updatedTask.Priority
			tasks[i].Completed = updatedTask.Completed
			return tasks[i]
		}
	}
	return Task{}
}

func DeleteTask(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}
}