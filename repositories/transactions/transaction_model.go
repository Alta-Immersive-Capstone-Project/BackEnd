package transactions

import (
	"kost/entities"

	"gorm.io/gorm"
)

type transactionModel struct {
	db *gorm.DB
}

func NewTransactionModel(db *gorm.DB) *transactionModel {
	return &transactionModel{
		db: db,
	}
}

func (m *transactionModel) Create(transaction *entities.Transaction) (*entities.Transaction, error) {
	record := m.db.Create(&transaction)

	if record.RowsAffected == 0 {
		return &entities.Transaction{}, record.Error
	}

	return transaction, nil
}

func (m *transactionModel) Get(booking_id string) (entities.Transaction, error) {
	var transaction entities.Transaction

	record := m.db.Where("booking_id = ?", booking_id).First(&transaction)

	if record.RowsAffected == 0 {
		return entities.Transaction{}, record.Error
	}

	return transaction, nil
}

func (m *transactionModel) GetAllbyCustomer(customer_id uint, status string) []entities.Transaction {
	var transactions []entities.Transaction

	m.db.Where("customer_id = ? AND status LIKE ?", customer_id, "%"+status+"%").Find(&transactions)

	return transactions
}

func (m *transactionModel) GetAllbyConsultant(consultant_id uint, status string) []entities.Transaction {
	var transactions []entities.Transaction

	m.db.Where("consultant_id = ? AND status LIKE ?", consultant_id, "%"+status+"%").Find(&transactions)

	return transactions
}

func (m *transactionModel) Update(booking_id string, transaction *entities.Transaction) (*entities.Transaction, error) {
	record := m.db.Where("booking_id = ?", booking_id).Updates(&transaction)

	if record.RowsAffected == 0 {
		return &entities.Transaction{}, record.Error
	}

	return transaction, nil
}
