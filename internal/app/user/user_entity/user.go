package user_entity

import (
	"github.com/yuuki798/TimerMe3/core/model"
	"time"
)

// 定义用户的结构
type User struct {
	//gorm.Model
	model.BaseModel
	Username     string `json:"username" gorm:"unique;not null"`
	PasswordHash string `json:"-"` // 存储加密后的密码,不进行传输
	Email        string `json:"email" gorm:"unique;not null"`
	Role         string `json:"role"` //admin 和 user
	//DateOfBirth    time.Time `json:"date_of_birth,omitempty"`
	ProfilePicture string    `json:"profile_picture,omitempty"` // 用户头像
	LastLogin      time.Time `json:"last_login,omitempty"`
	Valid          bool      `json:"valid,omitempty"` // 用户是否有效
	//VipExpireAt    time.Time `json:"vip_expire_at,omitempty"`   // vip到期时间
	//InterviewPoint int       `json:"interview_point,omitempty"` // 面试点数
	//Money          int64     `json:"money,omitempty"`           // 余额
}

//type SetBirthRequest struct {
//	Year  int `json:"year"`
//	Month int `json:"month"`
//	Day   int `json:"day"`
//}
//
//func SetBirth(c *gin.Context) {
//	uid := libx.Uid(c)
//	var req SetBirthRequest
//	c.ShouldBindJSON(&req)
//	var user User
//	db.Where("id = ?", uid).First(&user)
//
//	user.DateOfBirth = time.Date(req.Year, time.Month(req.Month), req.Day, 0, 0, 0, 0, time.UTC)
//	db.Save(&user)
//	libx.Ok(c, "设置生日成功", gin.H{
//		"date_of_birth": user.DateOfBirth,
//	})
//
//}
