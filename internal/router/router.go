package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuki798/TimerMe3/internal/app/ping"
	"github.com/yuuki798/TimerMe3/internal/app/user/user_handler"
)

type Entity struct {
}

func (r Entity) Router(g *gin.RouterGroup) {
	g.GET("/ping", ping.Handler)
	g.POST("/register", user_handler.EmailRegister)
	// 邮箱验证接口
	g.GET("/verify", user_handler.VerifyEmail)
	// 登录接口
	g.POST("/login", user_handler.Login)
	//g.GET("/mysql", ping.TestMysql)
	//g.GET("/redis", ping.TestRedis)
}
