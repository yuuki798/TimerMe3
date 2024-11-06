package user_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuki798/TimerMe3/core/database"
	"github.com/yuuki798/TimerMe3/core/libx"
	"github.com/yuuki798/TimerMe3/internal/app/user/user_entity"
	"net/http"
	"time"
)

func VerifyEmail(c *gin.Context) {
	//email := c.Query("email")
	token := c.Query("token")
	var db = database.GetDb("MainMysql")

	var verification user_entity.EmailVerification

	if err := db.Where("token = ? AND expires_at > ?", token, time.Now()).First(&verification).Error; err != nil {
		// 说明校验失败，删除Valid为false的用户
		if verification.UserID != 0 {
			db.Where("valid = ? and user_id=?", false, verification.UserID).Delete(&user_entity.User{})
			db.Where("user_id=?", verification.UserID).Delete(&user_entity.EmailVerification{})
		}
		libx.Err(c, http.StatusBadRequest, "无效或过期的验证令牌", err)
		return
	}
	var user user_entity.User
	if err := db.Where("id = ?", verification.UserID).First(&user).Error; err != nil {
		libx.Err(c, http.StatusInternalServerError, "用户不存在", err)
		return
	}
	user.Valid = true

	if err := db.Save(&user).Error; err != nil {
		libx.Err(c, http.StatusInternalServerError, "无法更新用户验证状态", err)
		return
	}

	// 删除或失效验证令牌
	db.Delete(&verification)

	libx.Ok(c, "邮箱验证成功，您的账号已可以使用")
}
