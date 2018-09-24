package task

import "time"

type Task struct {
	ID   int       `json:"id"`
	Task string    `json:"task"`
	Date time.Time `json:"date"`
	Done bool      `json:"done"`
}

var cpt int

var tasks = make(map[int]Task)

func AddTask(task Task) Task {
	cpt = cpt + 1

	task.ID = cpt

	tasks[task.ID] = task

	return task
}

func UpdateTask(task Task) {
	tasks[task.ID] = task
}

func GetTasks() []Task {

	values := []Task{}
	for _, value := range tasks {
		values = append(values, value)
	}
	return values
}
