package user_entity

import (
	"github.com/yuuki798/TimerMe3/core/model"
	"time"
)

// EmailVerification 存储邮箱验证信息
// 为什么是varchar(255)？就是怕token太长，导致数据库存储不下。说实话其实不设置默认也是255，但是加了比没加好
// uniqueIndex的话是primary的阉割版，可以保证token的唯一性。
// autoCreateTime是自动创建时间，按照go应用的时区。
// id一般不是负的，所以用uint，可以节省一半的空间
type EmailVerification struct {
	//gorm.Model
	model.BaseModel
	UserID    uint      `gorm:"not null"`                               // 关联的用户ID
	Token     string    `gorm:"type:varchar(255);uniqueIndex;not null"` // 唯一的验证令牌
	ExpiresAt time.Time `gorm:"not null"`                               // 令牌过期时间
}
