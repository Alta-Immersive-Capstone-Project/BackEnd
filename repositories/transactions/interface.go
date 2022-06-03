package transactions

import (
	"kost/entities"

	"github.com/midtrans/midtrans-go/snap"
)

type TransactionModel interface {
	Create(transaction entities.Transaction) (entities.Transaction, error)
	Request(booking_id string) (entities.TransactionResponse, error)
	CreateSnap(req *snap.Request) (*snap.Response, error)
	GetAllbyConsultant() []entities.Transaction
	Update(booking_id string, transaction entities.Transaction) (entities.Transaction, error)
	UpdateSnap(booking_id string, status entities.Callback) (entities.Callback, error)
	GetAllbyUser(role string, user uint, status string, city uint, district uint) []entities.TransactionJoin
	GetAllbyKost(duration int, status string, name string) []entities.TransactionKost
	GetTransactionByBookingID(BookingID string) (entities.DataReminder, error)
}
