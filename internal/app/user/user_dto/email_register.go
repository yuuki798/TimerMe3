package user_dto

// RegisterInput 存储注册信息
// required是必填项，如果没有填写，就会在下面ShouldBindJSON的时候返回400错误（bad request）
type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterOutput struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Valid    bool   `json:"valid"`
}
