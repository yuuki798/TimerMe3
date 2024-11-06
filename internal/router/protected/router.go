package protected

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuki798/TimerMe3/internal/app/ping"
	"github.com/yuuki798/TimerMe3/internal/app/task"
)

type Entity struct {
}

func (r Entity) Router(g *gin.RouterGroup) {
	g.GET("/ping", ping.Handler)
	g.GET("/tasks", task.GetTasks)
	g.POST("/tasks", task.CreateTask)
	g.PUT("/tasks/:id", task.UpdateTask)
	g.DELETE("/tasks/:id", task.DeleteTask)
	g.PUT("/tasks/:id/start", task.StartTask)
	g.PUT("/tasks/:id/pause", task.PauseTask)
	g.PUT("/tasks/:id/complete", task.CompleteTask)
	g.PUT("/tasks/:id/reset", task.ResetTask)
}
