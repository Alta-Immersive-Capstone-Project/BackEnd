package entities

import (
	"time"

	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	UserID  uint    `gorm:"type:int;not null"`
	HouseID uint    `json:"house_id" gorm:"type:int;not null"`
	Comment string  `json:"comment" gorm:"type:varchar(100);not null"`
	Rating  float32 `json:"rating" gorm:"type:int;not null"`
}

// Request
type ReviewRequest struct {
	HouseID uint    `json:"house_id" form:"house_id" validate:"required"`
	Comment string  `json:"comment" form:"comment" validate:"required"`
	Rating  float32 `json:"rating" form:"rating" validate:"required"`
}

// Response
type ReviewResponse struct {
	ID        uint      `json:"id"`
	HouseID   uint      `json:"house_id"`
	Comment   string    `json:"comment"`
	Rating    float32   `json:"rating"`
	CreatedAt time.Time `json:"created_at"`
}

type ReviewJoin struct {
	Name      string    `json:"name"`
	Comment   string    `json:"comment"`
	Rating    float32   `json:"rating"`
	CreatedAt time.Time `json:"created_at"`
}

type ReviewGetResponse struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Name      string    `json:"name"`
	HouseID   uint      `json:"house_id"`
	Comment   string    `json:"comment"`
	Rating    float32   `json:"rating"`
	CreatedAt time.Time `json:"created_at"`
}
