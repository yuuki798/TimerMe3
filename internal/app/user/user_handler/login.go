package user_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuki798/TimerMe3/core/auth"
	"github.com/yuuki798/TimerMe3/core/database"
	"github.com/yuuki798/TimerMe3/core/libx"
	"github.com/yuuki798/TimerMe3/internal/app/user/user_entity"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

func Login(c *gin.Context) {
	var db = database.GetDb("MainMysql")

	var loginVals struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&loginVals); err != nil {
		libx.Err(c, http.StatusBadRequest, "不合法的输入", err)
		return
	}
	var user user_entity.User

	if err := db.Where("email = ?", loginVals.Email).First(&user).Error; err != nil {
		libx.Err(c, http.StatusUnauthorized, "邮箱不存在", err)
		return
	}
	if !user.Valid {
		libx.Err(c, http.StatusUnauthorized, "用户未被激活，请检查邮箱或联系管理员", nil)
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginVals.Password)); err != nil {
		log.Println(err)
		log.Println(user.PasswordHash)
		log.Println(loginVals.Password)
		libx.Err(c, http.StatusUnauthorized, "用户名或密码错误", err)
		return
	}

	// 生成 JWT Token
	token, err := auth.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		libx.Err(c, http.StatusInternalServerError, "生成token失败", err)
		return
	}
	user.LastLogin = time.Now()
	db.Save(&user)

	libx.Ok(c, "登录成功", gin.H{
		"username": user.Username,
		"token":    token,
	})
}
