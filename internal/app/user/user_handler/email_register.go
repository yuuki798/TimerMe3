package user_handler

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yuuki798/TimerMe3/core/database"
	"github.com/yuuki798/TimerMe3/core/libx"
	"github.com/yuuki798/TimerMe3/internal/app/user/user_dto"
	"github.com/yuuki798/TimerMe3/internal/app/user/user_entity"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
	"log"
	"net/http"
	"regexp"
	"time"
)

// EmailRegister 处理邮箱注册请求
// 每次请求共享一个上下文Context，指针传递，所以不用担心并发问题
func EmailRegister(c *gin.Context) {
	var db = database.GetDb("MainMysql")

	var registerInput user_dto.RegisterInput
	// 从请求中解析JSON数据到registerInput结构体
	// ShouldBindJSON必须绑定的是结构体指针
	if err := c.ShouldBindJSON(&registerInput); err != nil {
		libx.Err(c, http.StatusBadRequest, "不合法的输入", nil)
		return
	}

	// 检查用户名是否重复
	var existingUser2 user_entity.User
	var verification user_entity.EmailVerification
	res2 := db.Where("username = ?", registerInput.Username).First(&existingUser2)

	// 过期重发
	if db.Where("user_id=?", existingUser2.ID).First(&verification).RowsAffected > 0 &&
		verification.ExpiresAt.Before(time.Now()) {
		db.Where("user_id=?", existingUser2.ID).Delete(&user_entity.EmailVerification{})
		if existingUser2.Valid == false {
			db.Where("id = ? and valid= 0", existingUser2.ID).Delete(&user_entity.User{})
		}
	} else if res2.RowsAffected > 0 {
		libx.Registered(c, "用户名已被注册")
		return
	}

	// 校验邮箱格式
	if !isValidEmail(registerInput.Email) {
		libx.Err(c, http.StatusBadRequest, "不合法的邮箱格式", nil)
		return
	}

	// 检查邮箱是否已被注册
	var existingUser user_entity.User
	result := db.Where("email = ?", registerInput.Email).First(&existingUser)
	if result.RowsAffected > 0 {
		libx.Registered(c, "邮箱已被注册")
		return
	}
	var newUser user_entity.User
	// 密码哈希处理，cost是哈希计算的成本，越高越安全，但是也越耗时，默认为10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerInput.Password), bcrypt.DefaultCost)
	if err != nil {
		libx.Err(c, http.StatusInternalServerError, "密码哈希失败", err)
		return
	}
	newUser.PasswordHash = string(hashedPassword)
	newUser.Role = "user" // 默认为普通用户，如果要管理员直接在数据库修改
	newUser.Username = registerInput.Username
	newUser.Email = registerInput.Email
	//newUser.DateOfBirth = time.Now()
	newUser.LastLogin = time.Now()
	//newUser.VipExpireAt = time.Now() // 不赠送VIP，所以注册时间就是过期时间
	//newUser.InterviewPoint = 2 * 60  // 注册送2*60个面试点（分钟），也就是2小时
	//newUser.InterviewPoint = 999 // 改回

	// 验证之后才可以true
	newUser.Valid = false

	// 保存用户信息到数据库
	if err := db.Create(&newUser).Error; err != nil {
		libx.Err(c, http.StatusInternalServerError, "建立用户信息失败", err)
		return
	}

	// 生成验证令牌并保存到数据库
	token, err := generateVerificationToken()
	if err != nil {
		libx.Err(c, http.StatusInternalServerError, "生成验证令牌失败", err)
		return
	}
	verification = user_entity.EmailVerification{
		UserID:    newUser.ID,
		Token:     token,
		ExpiresAt: time.Now().Add(24 * time.Hour), // 令牌24小时后过期
	}
	if err := db.Create(&verification).Error; err != nil {
		libx.Err(c, http.StatusInternalServerError, "创建验证信息失败", err)
		return
	}

	// 发送验证邮件
	err = sendVerificationEmail(newUser.ID, newUser.Email, token)
	if err != nil {
		libx.Err(c, http.StatusInternalServerError, "发送验证邮件失败", err)
		return
	}

	libx.Ok(c, "请查看邮箱，点击校验链接以继续完成注册", user_dto.RegisterOutput{
		Email:    newUser.Email,
		Username: newUser.Username,
		Valid:    newUser.Valid,
	})
}

// 验证邮箱格式是否正确
func isValidEmail(email string) bool {
	regex := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}

// 生成随机的验证令牌
func generateVerificationToken() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %w", err)
	}
	return hex.EncodeToString(bytes), nil
}

func sendVerificationEmail(userID uint, email, token string) error {
	var db = database.GetDb("MainMysql")

	// 定义SMTP服务器信息
	smtpHost := "smtp.qq.com"
	smtpPort := 465 // 使用465端口
	senderEmail := "1195396626@qq.com"
	senderPassword := "izsyvpvqegeyjaia" // QQ邮箱授权码

	// 打印发送邮件前的基本信息
	log.Printf("Starting to send verification email to: %s for user ID: %d", email, userID)
	log.Printf("SMTP Host: %s, SMTP Port: %d, Sender Email: %s", smtpHost, smtpPort, senderEmail)

	// 获取用户信息
	var newUser user_entity.User
	err2 := db.Where("id = ?", userID).First(&newUser).Error
	if err2 != nil {
		log.Printf("Error retrieving user with ID %d: %v", userID, err2)
		return fmt.Errorf("failed to get user: %w", err2)
	}
	username := newUser.Username

	// 构建验证链接
	//verificationLink := fmt.Sprintf("https://altar-echo.top/api/verify?token=%s", token)
	verificationLink := fmt.Sprintf("https://127.0.0.1:12349/verify?token=%s", token)
	subject := "【OfferCat】尊敬的" + username + "，请验证您的邮箱"
	body := fmt.Sprintf("尊敬的%s:\n您的邮箱被用于在OfferCat上注册了一个新账号，请验证您的邮箱。\n请点击以下链接以继续完成注册:\n %s", username, verificationLink)

	// 设置邮件信息
	m := gomail.NewMessage()
	m.SetHeader("From", senderEmail)
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	// 打印邮件内容，确保内容格式无误
	log.Printf("Email subject: %s", subject)
	log.Printf("Email body: %s", body)
	log.Printf("Verification link: %s", verificationLink)

	// 设置SMTP的认证信息
	d := gomail.NewDialer(smtpHost, smtpPort, senderEmail, senderPassword)
	d.SSL = true // 启用SSL

	// 捕捉连接SMTP时的详细错误信息
	log.Println("Starting email send process...")
	if err := d.DialAndSend(m); err != nil {
		// 直接打印错误信息
		log.Printf("Failed to send email: %v", err)
		return fmt.Errorf("failed to send email: %w", err)
	}

	// 邮件发送成功日志
	log.Printf("Email successfully sent to %s", email)
	return nil
}
