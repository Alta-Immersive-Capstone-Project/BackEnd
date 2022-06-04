package entities

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	BookingID         string    `gorm:"type:varchar(100);not null;unique"`
	UserID            uint      `gorm:"type:int;not null"`
	ConsultantID      uint      `gorm:"type:int"`
	HouseID           uint      `json:"house_id" gorm:"type:int;not null"`
	RoomID            uint      `json:"room_id" gorm:"type:int;not null"`
	CheckIn           time.Time `json:"check_in" gorm:"type:timestamp;not null"`
	Duration          int       `json:"duration" gorm:"type:int;not null"`
	Price             int64     `gorm:"type:int;not null"`
	TransactionStatus string    `gorm:"type:varchar(100);not null"`
	TransactionType   string    `gorm:"type:varchar(100)"`
	PaymentType       string    `gorm:"type:varchar(100)"`
	Acquirer          string    `gorm:"type:varchar(100)"`
	RedirectURL       string    `gorm:"type:varchar(100)"`
	PDFInvoicesURL    string    `json:"url_invoices"`
}

type TransactionRequest struct {
	HouseID  uint  `json:"house_id" form:"house_id" validate:"required"`
	RoomID   uint  `json:"room_id" form:"room_id" validate:"required"`
	CheckIn  int64 `json:"check_in" form:"check_in" validate:"required"`
	Duration int   `json:"duration" form:"duration" validate:"required"`
	Price    int64 `json:"price" form:"price" validate:"required"`
}

type TransactionResponse struct {
	BookingID         string    `json:"booking_id"`
	Name              string    `json:"name"`
	Email             string    `json:"email"`
	Phone             string    `json:"phone"`
	Title             string    `json:"title"`
	Address           string    `json:"address"`
	Price             int64     `json:"price"`
	CheckIn           time.Time `json:"check_in"`
	Duration          int       `json:"duration"`
	TransactionStatus string    `json:"transaction_status"`
	RedirectURL       string    `json:"redirect_url"`
	CreatedAt         time.Time `json:"created_at"`
}

type TransactionUpdateRequest struct {
	Price int64 `json:"price" form:"price" validate:"required"`
}

type TransactionUpdateResponse struct {
	BookingID         string    `json:"booking_id"`
	Name              string    `json:"name"`
	Phone             string    `json:"phone"`
	Title             string    `json:"title"`
	Address           string    `json:"address"`
	Price             int64     `json:"price"`
	CheckIn           time.Time `json:"check_in"`
	Duration          int       `json:"duration"`
	TransactionStatus string    `json:"transaction_status"`
	UpdatedAt         time.Time `json:"updated_at"`
	RedirectURL       string    `json:"redirect_url"`
	PDFInvoicesURL    string    `json:"url_invoices"`
}

type Callback struct {
	TransactionType   string `json:"transaction_type"`
	TransactionStatus string `json:"transaction_status"`
	TransactionID     string `json:"transaction_id"`
	StatusCode        string `json:"status_code"`
	SignatureKey      string `json:"signature_key"`
	PaymentType       string `json:"payment_type"`
	OrderID           string `json:"order_id"`
	GrossAmount       string `json:"gross_amount"`
	FraudStatus       string `json:"fraud_status"`
	Acquirer          string `json:"acquirer"`
	ApprovalCode      string `json:"approval_code"`
}

type TransactionUpdate struct {
	BookingID         string    `json:"booking_id"`
	TransactionStatus string    `json:"transaction_status"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type TransactionJoin struct {
	BookingID         string    `json:"booking_id"`
	Title             string    `json:"title"`
	Url               string    `json:"url"`
	CheckIn           time.Time `json:"check_in"`
	Duration          int       `json:"duration"`
	Price             int64     `json:"price"`
	TransactionStatus string    `json:"transaction_status"`
	RedirectURL       string    `json:"redirect_url"`
	PDFInvoicesURL    string    `json:"url_invoices"`
}

type TransactionKost struct {
	BookingID         string    `json:"booking_id"`
	Name              string    `json:"name"`
	Duration          int       `json:"duration"`
	Price             int64     `json:"price"`
	TransactionStatus string    `json:"transaction_status"`
	RedirectURL       string    `json:"redirect_url"`
	CreatedAt         time.Time `json:"created_at"`
}

type AddReminderPay struct {
	BookingID string `json:"booking_id"`
}

type DataReminder struct {
	BookingID   string `json:"booking_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Title       string `json:"title"`
	Price       int64  `json:"price"`
	RedirectURL string `json:"redirect_url"`
}
