package dbx

import (
	"github.com/yuuki798/TimerMe3/internal/app/ping"
	"github.com/yuuki798/TimerMe3/internal/app/task/task_entity"
	"github.com/yuuki798/TimerMe3/internal/app/user/user_entity"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&ping.TestModel{},
		&user_entity.User{},
		&user_entity.EmailVerification{},
		&task_entity.Task{},
	)
	return err
}
