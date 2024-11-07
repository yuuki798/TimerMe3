package protected

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuki798/TimerMe3/internal/app/ping"
	"github.com/yuuki798/TimerMe3/internal/app/task/task_handler"
	"github.com/yuuki798/TimerMe3/internal/app/user/user_handler"
)

type Entity struct {
}

// Router /api
func (r Entity) Router(g *gin.RouterGroup) {
	g.GET("/ping", ping.Handler)
	g.GET("/profile", user_handler.GetProfile)

	g.GET("/tasks", task_handler.GetTasks)
	g.POST("/tasks", task_handler.CreateTask)
	g.PUT("/tasks/:id", task_handler.UpdateTask)
	g.DELETE("/tasks/:id", task_handler.DeleteTask)

	g.PUT("/tasks/:id/start", task_handler.StartTask)
	g.PUT("/tasks/:id/pause", task_handler.PauseTask)
	g.PUT("/tasks/:id/complete", task_handler.CompleteTask)
	g.PUT("/tasks/:id/reset", task_handler.ResetTask)
}
