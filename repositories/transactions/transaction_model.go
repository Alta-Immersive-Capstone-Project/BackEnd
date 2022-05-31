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

func (m *transactionModel) Create(transaction entities.Transaction) (entities.Transaction, error) {
	record := m.db.Create(&transaction)

	if record.RowsAffected == 0 {
		return entities.Transaction{}, record.Error
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

func (m *transactionModel) GetAllbyCustomer(role string, user uint, status string, city uint, district uint) []entities.TransactionJoin {
	var transactions []entities.TransactionJoin

	query := "select distinct t.id, t.checkin_date, t.rent_duration, t.booking_id, t.total_bill, t.status, i.url, h.title from transactions as t left join rooms r on r.id = t.room_id left join images i on i.room_id = t.room_id left join houses h on h.id = r.house_id left join districts d on d.city_id = h.district_id left join cities c on c.id = d.city_id where t.status = ? and c.id = ? and h.district_id = ?"
	if role == "customer" {
		query += " and t.user_id = ?"
	} else {
		query += " and t.consultant_id = ?"
	}

	m.db.Raw(query, status, city, district, user).Scan(&transactions)
	return transactions
}

func (m *transactionModel) GetAllbyConsultant() []entities.Transaction {
	var transactions []entities.Transaction

	m.db.Where("consultant_id = ?", 0).Find(&transactions)

	return transactions
}

func (m *transactionModel) Update(booking_id string, transaction entities.Transaction) (entities.Transaction, error) {
	record := m.db.Where("booking_id = ?", booking_id).Updates(&transaction)

	if record.RowsAffected == 0 {
		return entities.Transaction{}, record.Error
	}

	return transaction, nil
}

func (m *transactionModel) GetAllbyKost(duration int, status string, name string) []entities.TransactionKost {
	var transactions []entities.TransactionKost

	query := "select distinct t.booking_id, u.name, t.checkin_date, t.rent_duration, t.total_bill, t.status, t.created_at from transactions as t left join rooms r on r.id = t.room_id left join houses h on h.id = r.house_id left join users u on u.id = t.user_id where t.rent_duration = ? and t.status = ? OR h.title LIKE ?"

	m.db.Raw(query, duration, status, "%"+name+"%").Scan(&transactions)
	return transactions
}
