package task_entity

import (
	"github.com/yuuki798/TimerMe3/core/model"
	"time"
)

type Task struct {
	model.BaseModel
	Uid         uint      `json:"uid" form:"uid"`
	Name        string    `json:"name" form:"name"`
	Duration    int       `json:"duration" form:"duration"`
	IsCompleted bool      `json:"is_completed" form:"is_completed"`
	StartTime   time.Time `json:"start_time" form:"start_time"`
	Status      string    `json:"status" form:"status"`
	TotalTime   int       `json:"total_time" form:"total_time"`
}
