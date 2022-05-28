package transactions

import "kost/entities"

type TransactionService interface {
	AddTransaction(customer_id uint, request *entities.TransactionRequest) (*entities.TransactionResponse, error)
	GetTransaction(booking_id string) (*entities.TransactionResponse, error)
	GetAllTransactionbyCustomer(customer_id uint, status string) []entities.TransactionResponse
	GetAllTransactionbyConsultant(consultant_id uint, status string) []entities.TransactionResponse
	UpdateTransaction(customer_id uint, booking_id string, request *entities.TransactionUpdateRequest) (*entities.TransactionUpdateResponse, error)
}
