package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	TaskName string `gorm:"size:64;not null"`
	IsDone   bool   `gorm:"default:false"`
}
