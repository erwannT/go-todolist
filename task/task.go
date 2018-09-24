package task

import (
	"time"

	"github.com/erwannT/go-todolist/connection"
)

type Task struct {
	ID   int       `json:"id" gorm:"primary_key"`
	Task string    `json:"task"`
	Date time.Time `json:"date"`
	Done bool      `json:"done"`
}

func AddTask(task Task) Task {
	db := connection.DB()
	db.Create(&task)
	return task
}

func UpdateTask(task Task) {
	db := connection.DB()
	db.Save(&task)
}

func GetTasks() []Task {
	values := []Task{}
	db := connection.DB()
	db.Find(&values)
	return values
}
