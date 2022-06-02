package transactions

import "kost/entities"

type TransactionService interface {
	CreateTransaction(customer_id uint, req entities.TransactionRequest) (entities.TransactionResponse, error)
	GetAllTransactionbyConsultant() []entities.TransactionResponse
	UpdateTransaction(customer_id uint, booking_id string, request entities.TransactionUpdateRequest) (entities.TransactionUpdateResponse, error)
	UpdateCallback(req entities.Callback) (entities.Callback, error)
	GetAllTransactionbyCustomer(role string, user uint, status string, city uint, district uint) []entities.TransactionJoin
	GetAllTransactionbyKost(duration int, status string, name string) []entities.TransactionKost
}
