package models

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TimeLog struct {
	ID        int    `json:"id"`
	TaskID    int    `json:"task_id"`
	Duration  int    `json:"duration"` // Duration in minutes
	Timestamp string `json:"timestamp"`
}

var tasks []Task
var timeLogs []TimeLog
var taskIDCounter int
var timeLogIDCounter int

func init() {
	tasks = []Task{}
	timeLogs = []TimeLog{}
	taskIDCounter = 1
	timeLogIDCounter = 1
}

func GetAllTasks() []Task {
	return tasks
}

func CreateTask(task Task) {
	task.ID = taskIDCounter
	taskIDCounter++
	tasks = append(tasks, task)
}

func UpdateTask(id int, updatedTask Task) {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i] = updatedTask
			tasks[i].ID = id
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

func GetAllTimeLogs() []TimeLog {
	return timeLogs
}

func CreateTimeLog(timeLog TimeLog) {
	timeLog.ID = timeLogIDCounter
	timeLogIDCounter++
	timeLogs = append(timeLogs, timeLog)
}

func UpdateTimeLog(id int, updatedTimeLog TimeLog) {
	for i, timeLog := range timeLogs {
		if timeLog.ID == id {
			timeLogs[i] = updatedTimeLog
			timeLogs[i].ID = id
			break
		}
	}
}

func DeleteTimeLog(id int) {
	for i, timeLog := range timeLogs {
		if timeLog.ID == id {
			timeLogs = append(timeLogs[:i], timeLogs[i+1:]...)
			break
		}
	}
}