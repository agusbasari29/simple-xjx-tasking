package entity

import (
	"time"

	"gorm.io/gorm"
)

type TaskStatus string

const (
	Progress  TaskStatus = "progress"
	Idle      TaskStatus = "idle"
	Completed TaskStatus = "completed"
)

type Tasks struct {
	gorm.Model
	ID       uint `gorm:"primaryKey;autoIncrement"`
	Task     string
	Assignee string
	Status   TaskStatus
	Deadline time.Time
}
