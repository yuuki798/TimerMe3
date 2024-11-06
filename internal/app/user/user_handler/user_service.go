package user_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuki798/TimerMe3/core/libx"
)

func GetProfile(c *gin.Context) {
	uid := libx.Uid(c)
	role := libx.GetRole(c)
	username := libx.GetUsername(c)
	libx.Ok(c, "获取用户信息成功", gin.H{
		"uid":      uid,
		"username": username,
		"role":     role,
	})
}
