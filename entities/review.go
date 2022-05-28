package entities

import (
	"time"

	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	UserID  uint   `gorm:"type:int;not null"`
	RoomID  uint   `json:"room_id" gorm:"type:int;not null"`
	Comment string `json:"comment" gorm:"type:varchar(100);not null"`
	Rating  int    `json:"rating" gorm:"type:int;not null"`
}

// Request
type ReviewRequest struct {
	RoomID  uint   `json:"room_id" form:"room_id" validate:"required"`
	Comment string `json:"comment" form:"comment" validate:"required"`
	Rating  int    `json:"rating" form:"rating" validate:"required"`
}

// Response
type ReviewResponse struct {
	ID        uint      `json:"id"`
	RoomID    uint      `json:"room_id"`
	Comment   string    `json:"comment"`
	Rating    int       `json:"rating"`
	CreatedAt time.Time `json:"created_at"`
}

type ReviewGetResponse struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Name      string    `json:"name"`
	RoomID    uint      `json:"room_id"`
	Comment   string    `json:"comment"`
	Rating    int       `json:"rating"`
	CreatedAt time.Time `json:"created_at"`
}
