package model

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	TaskName  string
	Domain    string
	IsDeleted int64
	State     int64
}

func Create_taskTable() {
	db.AutoMigrate(&Task{})
	//创建Task表
}

func InsertTask(taskName string, domain string) {
	db.Create(&Task{TaskName: taskName, Domain: domain, IsDeleted: 0, State: 1})

}

func GetTaskByName() {
	tasks := []Task{}
	db.Find(&tasks)
	for _, task := range tasks {
		println(task.TaskName)
		println(task.Domain)
	}
}
