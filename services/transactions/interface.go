package transactions

import "kost/entities"

type TransactionService interface {
	AddTransaction(customer_id uint, request entities.TransactionRequest) (entities.TransactionResponse, error)
	GetTransaction(booking_id string) (entities.TransactionResponse, error)
	GetAllTransactionbyConsultant() []entities.TransactionResponse
	UpdateTransaction(customer_id uint, booking_id string, request entities.TransactionUpdateRequest) (entities.TransactionUpdateResponse, error)
	GetAllTransactionbyCustomer(role string, user uint, status string, city uint, district uint) []entities.TransactionJoin
	GetAllTransactionbyKost(duration int, status string, name string) []entities.TransactionKost
	UpdateStatus(request entities.Callback) error
}
