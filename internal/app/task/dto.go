package task

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	//ID          uint      `gorm:"primary_key" json:"id"`
	gorm.Model
	Name        string    `json:"name"`
	Duration    int       `json:"duration"`
	IsCompleted bool      `json:"is_completed"`
	StartTime   time.Time `json:"start_time"`
	Status      string    `json:"status"`
	TotalTime   int       `json:"total_time"`
}
