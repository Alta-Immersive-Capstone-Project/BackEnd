package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Name     string
	Password string
	Gender   string
	Phone    string
	Avatar   string
	Role     string
}

type CreateInternalRequest struct {
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required"`
	Name     string `form:"name" validate:"required"`
	Gender   string `form:"gender" validate:"required"`
	Phone    string `form:"phone" validate:"required"`
	Role     string `form:"role" validate:"required"`
}

type UpdateInternalRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
	Name     string `form:"name"`
	Gender   string `form:"gender"`
	Phone    string `form:"phone" validate:"required"`
	Role     string `form:"role" validate:"required"`
}
type InternalResponse struct {
	ID          uint      `json:"id"`
	Email       string    `json:"email"`
	Name        string    `json:"name"`
	Gender      string    `json:"gender"`
	PhoneNumber string    `json:"phone"`
	Avatar      string    `json:"avatar"`
	Role        string    `json:"role"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
