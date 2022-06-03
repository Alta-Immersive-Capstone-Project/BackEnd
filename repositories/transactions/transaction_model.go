package transactions

import (
	"kost/entities"

	"github.com/midtrans/midtrans-go/snap"
	"gorm.io/gorm"
)

type transactionModel struct {
	db   *gorm.DB
	snap snap.Client
}

func NewTransactionModel(db *gorm.DB, snap snap.Client) *transactionModel {
	return &transactionModel{
		db:   db,
		snap: snap,
	}
}

func (m *transactionModel) Create(transaction entities.Transaction) (entities.Transaction, error) {
	record := m.db.Create(&transaction)

	if record.RowsAffected == 0 {
		return entities.Transaction{}, record.Error
	}

	return transaction, nil
}

func (m *transactionModel) Request(booking_id string) (entities.TransactionResponse, error) {
	var transaction entities.TransactionResponse

	record := m.db.Raw("select t.booking_id, t.check_in, t.duration, u.name, u.email, u.phone, h.title, h.address, r.price from transactions t left join users u on u.id = t.user_id left join rooms r on r.id = t.room_id left join houses h on h.id = r.house_id where t.booking_id = ?", booking_id).Scan(&transaction)

	if record.RowsAffected == 0 {
		return entities.TransactionResponse{}, record.Error
	}

	return transaction, nil
}

func (m *transactionModel) CreateSnap(req *snap.Request) (*snap.Response, error) {
	transaction, err := m.snap.CreateTransaction(req)

	if err != nil {
		return nil, err
	}

	return transaction, nil
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

func (m *transactionModel) UpdateSnap(booking_id string, status entities.Callback) (entities.Callback, error) {

	record := m.db.Model(entities.Transaction{}).Where("booking_id = ?", booking_id).Updates(&status)
	if record.RowsAffected == 0 {
		return entities.Callback{}, record.Error
	}

	return status, nil
}

func (m *transactionModel) GetAllbyUser(role string, user uint, status string, city uint, district uint) []entities.TransactionJoin {
	var transactions []entities.TransactionJoin
	var player string
	var query string
	if role == "customer" {
		player = "t.user_id = ?"
	} else if role == "t.consultant" {
		player = "consultant_id = ?"
	}
	if role != "supervisor" && role != "admin" {
		if status == "" && city == 0 && district == 0 {
			query = "select distinct t.id, t.check_in, t.duration, t.booking_id, t.price, t.transaction_status, t.redirect_url, i.url, h.title from transactions as t left join rooms r on r.id = t.room_id left join images i on i.room_id = t.room_id left join houses h on h.id = r.house_id left join districts d on d.id = h.district_id left join cities c on c.id = d.city_id where " + player
			m.db.Raw(query, user).Scan(&transactions)
		} else if status != "" && city == 0 && district == 0 {
			query = "select distinct t.id, t.check_in, t.duration, t.booking_id, t.price, t.transaction_status, t.redirect_url, i.url, h.title from transactions as t left join rooms r on r.id = t.room_id left join images i on i.room_id = t.room_id left join houses h on h.id = r.house_id left join districts d on d.id = h.district_id left join cities c on c.id = d.city_id where t.transaction_status = ? AND " + player
			m.db.Raw(query, status, user).Scan(&transactions)
		} else if status != "" && city != 0 && district == 0 {
			query = "select distinct t.id, t.check_in, t.duration, t.booking_id, t.price, t.transaction_status, t.redirect_url, i.url, h.title from transactions as t left join rooms r on r.id = t.room_id left join images i on i.room_id = t.room_id left join houses h on h.id = r.house_id left join districts d on d.id = h.district_id left join cities c on c.id = d.city_id where t.transaction_status = ? AND c.id = ? AND " + player
			m.db.Raw(query, status, city, user).Scan(&transactions)
		} else if status != "" && city != 0 && district != 0 {
			query = "select distinct t.id, t.check_in, t.duration, t.booking_id, t.price, t.transaction_status, t.redirect_url, i.url, h.title from transactions as t left join rooms r on r.id = t.room_id left join images i on i.room_id = t.room_id left join houses h on h.id = r.house_id left join districts d on d.id = h.district_id left join cities c on c.id = d.city_id where t.transaction_status = ? AND c.id = ? AND h.district_id = ? AND " + player
			m.db.Raw(query, status, city, district, user).Scan(&transactions)
		} else if status == "" && city != 0 && district == 0 {
			query = "select distinct t.id, t.check_in, t.duration, t.booking_id, t.price, t.transaction_status, t.redirect_url, i.url, h.title from transactions as t left join rooms r on r.id = t.room_id left join images i on i.room_id = t.room_id left join houses h on h.id = r.house_id left join districts d on d.id = h.district_id left join cities c on c.id = d.city_id where c.id = ? AND " + player
			m.db.Raw(query, city, user).Scan(&transactions)
		} else if status == "" && city != 0 && district != 0 {
			query = "select distinct t.id, t.check_in, t.duration, t.booking_id, t.price, t.transaction_status, t.redirect_url, i.url, h.title from transactions as t left join rooms r on r.id = t.room_id left join images i on i.room_id = t.room_id left join houses h on h.id = r.house_id left join districts d on d.id = h.district_id left join cities c on c.id = d.city_id where c.id = ? AND h.district_id = ? AND " + player
			m.db.Raw(query, city, district, user).Scan(&transactions)
		} else if status != "" && city == 0 && district != 0 {
			query = "select distinct t.id, t.check_in, t.duration, t.booking_id, t.price, t.transaction_status, t.redirect_url, i.url, h.title from transactions as t left join rooms r on r.id = t.room_id left join images i on i.room_id = t.room_id left join houses h on h.id = r.house_id left join districts d on d.id = h.district_id left join cities c on c.id = d.city_id where t.transaction_status = ? AND h.district_id = ? AND " + player
			m.db.Raw(query, status, district, user).Scan(&transactions)
		} else {
			query = "select distinct t.id, t.check_in, t.duration, t.booking_id, t.price, t.transaction_status, t.redirect_url, i.url, h.title from transactions as t left join rooms r on r.id = t.room_id left join images i on i.room_id = t.room_id left join houses h on h.id = r.house_id left join districts d on d.id = h.district_id left join cities c on c.id = d.city_id where h.district_id = ? AND " + player
			m.db.Raw(query, district, user).Scan(&transactions)
		}
		return transactions
	}

	if status == "" && city == 0 && district == 0 {
		query = "select distinct t.id, t.check_in, t.duration, t.booking_id, t.price, t.transaction_status, t.redirect_url, i.url, h.title from transactions as t left join rooms r on r.id = t.room_id left join images i on i.room_id = t.room_id left join houses h on h.id = r.house_id left join districts d on d.id = h.district_id left join cities c on c.id = d.city_id"
		m.db.Raw(query).Scan(&transactions)
	} else if status != "" && city == 0 && district == 0 {
		query = "select distinct t.id, t.check_in, t.duration, t.booking_id, t.price, t.transaction_status, t.redirect_url, i.url, h.title from transactions as t left join rooms r on r.id = t.room_id left join images i on i.room_id = t.room_id left join houses h on h.id = r.house_id left join districts d on d.id = h.district_id left join cities c on c.id = d.city_id where t.transaction_status = ?"
		m.db.Raw(query, status).Scan(&transactions)
	} else if status != "" && city != 0 && district == 0 {
		query = "select distinct t.id, t.check_in, t.duration, t.booking_id, t.price, t.transaction_status, t.redirect_url, i.url, h.title from transactions as t left join rooms r on r.id = t.room_id left join images i on i.room_id = t.room_id left join houses h on h.id = r.house_id left join districts d on d.id = h.district_id left join cities c on c.id = d.city_id where t.transaction_status = ? AND c.id = ?"
		m.db.Raw(query, status, city).Scan(&transactions)
	} else if status != "" && city != 0 && district != 0 {
		query = "select distinct t.id, t.check_in, t.duration, t.booking_id, t.price, t.transaction_status, t.redirect_url, i.url, h.title from transactions as t left join rooms r on r.id = t.room_id left join images i on i.room_id = t.room_id left join houses h on h.id = r.house_id left join districts d on d.id = h.district_id left join cities c on c.id = d.city_id where t.transaction_status = ? AND c.id = ? AND h.district_id = ?"
		m.db.Raw(query, status, city, district).Scan(&transactions)
	} else if status == "" && city != 0 && district == 0 {
		query = "select distinct t.id, t.check_in, t.duration, t.booking_id, t.price, t.transaction_status, t.redirect_url, i.url, h.title from transactions as t left join rooms r on r.id = t.room_id left join images i on i.room_id = t.room_id left join houses h on h.id = r.house_id left join districts d on d.id = h.district_id left join cities c on c.id = d.city_id where c.id = ?"
		m.db.Raw(query, city).Scan(&transactions)
	} else if status == "" && city != 0 && district != 0 {
		query = "select distinct t.id, t.check_in, t.duration, t.booking_id, t.price, t.transaction_status, t.redirect_url, i.url, h.title from transactions as t left join rooms r on r.id = t.room_id left join images i on i.room_id = t.room_id left join houses h on h.id = r.house_id left join districts d on d.id = h.district_id left join cities c on c.id = d.city_id where c.id = ? AND h.district_id = ?"
		m.db.Raw(query, city, district).Scan(&transactions)
	} else if status != "" && city == 0 && district != 0 {
		query = "select distinct t.id, t.check_in, t.duration, t.booking_id, t.price, t.transaction_status, t.redirect_url, i.url, h.title from transactions as t left join rooms r on r.id = t.room_id left join images i on i.room_id = t.room_id left join houses h on h.id = r.house_id left join districts d on d.id = h.district_id left join cities c on c.id = d.city_id where t.transaction_status = ? AND h.district_id = ?"
		m.db.Raw(query, status, district).Scan(&transactions)
	} else {
		query = "select distinct t.id, t.check_in, t.duration, t.booking_id, t.price, t.transaction_status, t.redirect_url, i.url, h.title from transactions as t left join rooms r on r.id = t.room_id left join images i on i.room_id = t.room_id left join houses h on h.id = r.house_id left join districts d on d.id = h.district_id left join cities c on c.id = d.city_id where h.district_id = ?"
		m.db.Raw(query, district).Scan(&transactions)
	}
	return transactions
}

func (m *transactionModel) GetAllbyKost(duration int, status string, name string) []entities.TransactionKost {
	var transactions []entities.TransactionKost

	query := "select distinct t.booking_id, u.name, t.check_in, t.duration, t.price, t.transaction_status, t.redirect_url, t.created_at from transactions as t left join rooms r on r.id = t.room_id left join houses h on h.id = r.house_id left join users u on u.id = t.user_id where t.duration = ? and t.transaction_status = ? OR h.title LIKE ?"

	m.db.Raw(query, duration, status, "%"+name+"%").Scan(&transactions)

	return transactions
}
