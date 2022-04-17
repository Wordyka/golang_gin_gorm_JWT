package dto

// struct LoginDTO untuk digunakan user / client ketika melakukan POST dari url /login 
type LoginDTO struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}


