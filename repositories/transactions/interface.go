package transactions

import (
	"kost/entities"

	"github.com/midtrans/midtrans-go/snap"
)

type TransactionModel interface {
	Create(transaction entities.Transaction) (entities.Transaction, error)
	CreateSnap(req *snap.Request) (*snap.Response, error)
	Get(booking_id string) (entities.Transaction, error)
	GetAllbyConsultant() []entities.Transaction
	Update(booking_id string, transaction entities.Transaction) (entities.Transaction, error)
	GetAllbyCustomer(role string, user uint, status string, city uint, district uint) []entities.TransactionJoin
	GetAllbyKost(duration int, status string, name string) []entities.TransactionKost
	UpdateStatus(booking_id string, status entities.Callback) (entities.Callback, error)
}
