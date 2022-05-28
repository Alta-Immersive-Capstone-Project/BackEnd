package entities

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name         string `gorm:"type:varchar(100);not null"`
	Email        string `gorm:"type:varchar(100);not null"`
	Phone        string `gorm:"type:varchar(100);not null"`
	Gender       string
	Role         string
	Password     string
	Rooms        []Room        `gorm:"foreingkey:CustomerID"`
	Transactions []Transaction `gorm:"foreingkey:CustomerID"`
	Reviews      []Review      `gorm:"foreingkey:CustomerID"`
}

type Room struct {
	gorm.Model
	Type                  string        `gorm:"type:varchar(100);not null"`
	Price                 int           `gorm:"type:int;not null"`
	AdditionalDescription string        `gorm:"type:varchar(100);not null"`
	CustomerID            uint          `gorm:"type:int;not null"`
	Transactions          []Transaction `gorm:"foreingkey:RoomID"`
	Reviews               []Review      `gorm:"foreingkey:RoomID"`
}

type Transaction struct {
	gorm.Model
	CustomerID    uint      `gorm:"type:int;not null"`
	ConsultantID  uint      `gorm:"type:int"`
	RoomID        uint      `json:"room_id" gorm:"type:int;not null"`
	CheckinDate   time.Time `json:"checkin_date" gorm:"type:date;not null"`
	RentDuration  int       `json:"rent_duration" gorm:"type:int;not null"`
	BookingID     string    `gorm:"type:varchar(100);not null"`
	TotalBill     int       `json:"total_bill" gorm:"type:int;not null"`
	PaymentMethod string    `json:"payment_method" gorm:"type:varchar(100);not null"`
	Status        string    `gorm:"type:varchar(100);not null"`
}

type Review struct {
	gorm.Model
	CustomerID uint   `gorm:"type:int;not null"`
	RoomID     uint   `json:"room_id" gorm:"type:int;not null"`
	Comment    string `json:"comment" gorm:"type:varchar(100);not null"`
	Rating     int    `json:"rating" gorm:"type:int;not null"`
}

// Request
type TransactionRequest struct {
	RoomID        uint   `json:"room_id" validate:"required"`
	CheckinDate   int64  `json:"checkin_date" validate:"required"`
	RentDuration  int    `json:"rent_duration" validate:"required"`
	TotalBill     int    `json:"total_bill" validate:"required"`
	PaymentMethod string `json:"payment_method" validate:"required"`
}

type TransactionUpdateRequest struct {
	TotalBill int `json:"total_bill" validate:"required"`
}

type ReviewRequest struct {
	RoomID  uint   `json:"room_id" validate:"required"`
	Comment string `json:"comment" validate:"required"`
	Rating  int    `json:"rating" validate:"required"`
}

// Response
type TransactionResponse struct {
	ID            uint      `json:"id"`
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

type ReviewResponse struct {
	ID        uint      `json:"id"`
	RoomID    uint      `json:"room_id"`
	Comment   string    `json:"comment"`
	Rating    int       `json:"rating"`
	CreatedAt time.Time `json:"created_at"`
}

type ReviewGetResponse struct {
	ID         uint      `json:"id"`
	CustomerID uint      `json:"customer_id"`
	Name       string    `json:"name"`
	RoomID     uint      `json:"room_id"`
	Comment    string    `json:"comment"`
	Rating     int       `json:"rating"`
	CreatedAt  time.Time `json:"created_at"`
}
