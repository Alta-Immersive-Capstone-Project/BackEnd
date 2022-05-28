package entities

import (
	"time"

	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Type                  string        `gorm:"type:varchar(100);not null"`
	Price                 int           `gorm:"type:int;not null"`
	AdditionalDescription string        `gorm:"type:varchar(100);not null"`
	UserID                uint          `gorm:"type:int;not null"`
	Transactions          []Transaction `gorm:"foreingkey:RoomID"`
	Reviews               []Review      `gorm:"foreingkey:RoomID"`
}

type Transaction struct {
	gorm.Model
	UserID        uint      `gorm:"type:int;not null"`
	ConsultantID  uint      `gorm:"type:int"`
	RoomID        uint      `json:"room_id" gorm:"type:int;not null"`
	CheckinDate   time.Time `json:"checkin_date" gorm:"type:date;not null"`
	RentDuration  int       `json:"rent_duration" gorm:"type:int;not null"`
	BookingID     string    `gorm:"type:varchar(100);not null"`
	TotalBill     int       `json:"total_bill" gorm:"type:int;not null"`
	PaymentMethod string    `json:"payment_method" gorm:"type:varchar(100);not null"`
	Status        string    `gorm:"type:varchar(100);not null"`
}

// Request
type TransactionRequest struct {
	RoomID        uint   `json:"room_id" form:"room_id" validate:"required"`
	CheckinDate   int64  `json:"checkin_date" form:"checkin_date" validate:"required"`
	RentDuration  int    `json:"rent_duration" form:"rent_duration" validate:"required"`
	TotalBill     int    `json:"total_bill" form:"total_bill" validate:"required"`
	PaymentMethod string `json:"payment_method" form:"payment_method" validate:"required"`
}

type TransactionUpdateRequest struct {
	TotalBill int `json:"total_bill" form:"total_bill" validate:"required"`
}

// Response
type TransactionResponse struct {
	ID            uint      `json:"id"`
	BookingID     string    `json:"booking_id"`
	RoomID        uint      `json:"room_id"`
	CheckinDate   time.Time `json:"checkin_date"`
	RentDuration  int       `json:"rent_duration"`
	TotalBill     int       `json:"total_bill"`
	PaymentMethod string    `json:"payment_method"`
	CreatedAt     time.Time `json:"created_at"`
}

type TransactionUpdateResponse struct {
	TotalBill int       `json:"total_bill"`
	UpdatedAt time.Time `json:"updated_at"`
}
