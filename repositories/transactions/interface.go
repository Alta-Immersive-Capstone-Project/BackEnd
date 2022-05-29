package transactions

import "kost/entities"

type TransactionModel interface {
	Create(transaction entities.Transaction) (entities.Transaction, error)
	Get(booking_id string) (entities.Transaction, error)
	GetAllbyCustomer(customer_id uint, status string) []entities.Transaction
	GetAllbyConsultant() []entities.Transaction
	Update(booking_id string, transaction entities.Transaction) (entities.Transaction, error)
}
