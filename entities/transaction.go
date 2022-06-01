package entities

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserID            uint      `gorm:"type:int;not null"`
	ConsultantID      uint      `gorm:"type:int"`
	RoomID            uint      `json:"room_id" gorm:"type:int;not null"`
	HouseID           uint      `json:"house_id" gorm:"type:int;not null"`
	CheckinDate       time.Time `json:"checkin_date" gorm:"type:date;not null"`
	RentDuration      int       `json:"rent_duration" gorm:"type:int;not null"`
	BookingID         string    `gorm:"type:varchar(100);not null"`
	TotalBill         int64     `json:"total_bill" gorm:"type:int;not null"`
	PaymentType       string    `json:"payment_type" gorm:"type:varchar(100)"`
	TransactionStatus string    `gorm:"type:varchar(100);not null"`
	Token             string    `gorm:"type:varchar(100);not null"`
}

// Request
type TransactionRequest struct {
	RoomID       uint  `json:"room_id" form:"room_id" validate:"required"`
	HouseID      uint  `json:"house_id" form:"house_id" validate:"required"`
	CheckinDate  int64 `json:"checkin_date" form:"checkin_date" validate:"required"`
	RentDuration int   `json:"rent_duration" form:"rent_duration" validate:"required"`
	TotalBill    int64 `json:"total_bill" form:"total_bill" validate:"required"`
}

type Callback struct {
	TransactionStatus string `json:"transaction_status"`
	TransactionID     string `json:"transaction_id"`
	StatusCode        string `json:"status_code"`
	SignatureKey      string `json:"signature_key"`
	PaymentType       string `json:"payment_type"`
	OrderID           string `json:"order_id"`
	GrossAmount       string `json:"gross_amount"`
	FraudStatus       string `json:"fraud_status"`
	ApprovalCode      string `json:"approval_code"`
}

type TransactionUpdateRequest struct {
	TotalBill int64 `json:"total_bill" form:"total_bill" validate:"required"`
}

// Response
type TransactionResponse struct {
	ID                uint      `json:"id"`
	BookingID         string    `json:"booking_id"`
	RoomID            uint      `json:"room_id"`
	Title             string    `json:"title"`
	Address           string    `json:"address"`
	CheckinDate       time.Time `json:"checkin_date"`
	RentDuration      int       `json:"rent_duration"`
	Name              string    `json:"name"`
	Phone             string    `json:"phone"`
	TotalBill         int64     `json:"total_bill"`
	TransactionStatus string    `json:"transaction_status"`
	RedirectURL       string    `json:"redirect_url"`
	CreatedAt         time.Time `json:"created_at"`
}

type TransactionUpdate struct {
	BookingID         string    `json:"booking_id"`
	TransactionStatus string    `json:"transaction_status"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type TransactionJoin struct {
	BookingID         string    `json:"booking_id"`
	CheckinDate       time.Time `json:"checkin_date"`
	RentDuration      int       `json:"rent_duration"`
	TotalBill         int64     `json:"total_bill"`
	TransactionStatus string    `json:"transaction_status"`
	Url               string    `json:"url"`
	Title             string    `json:"title"`
}

type TransactionKost struct {
	BookingID         string    `json:"booking_id"`
	Name              string    `json:"name"`
	RentDuration      int       `json:"rent_duration"`
	TotalBill         int64     `json:"total_bill"`
	TransactionStatus string    `json:"transaction_status"`
	CreatedAt         time.Time `json:"created_at"`
}

type TransactionUpdateResponse struct {
	TotalBill int64     `json:"total_bill"`
	UpdatedAt time.Time `json:"updated_at"`
}
