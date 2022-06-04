package transactions

import (
	"fmt"
	"kost/entities"

	"github.com/labstack/gommon/log"
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

func (m *transactionModel) Request(BookingID string) (entities.TransactionResponse, error) {
	var response entities.TransactionResponse

	if err := m.db.Raw("select t.booking_id, t.check_in, t.duration, u.name, u.email, h.title, u.phone, t.price, h.address from transactions as t left join houses h on h.id = t.house_id left join users u on u.id = t.user_id where t.booking_id = ?", BookingID).First(&response).Error; err != nil {
		log.Warn(err)
		return response, err
	}
	return response, nil

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
		log.Warn(record.RowsAffected)
		return entities.Transaction{}, record.Error
	}

	return transaction, nil
}

func (m *transactionModel) UpdateSnap(booking_id string, status entities.Callback) (entities.Callback, error) {

	record := m.db.Model(entities.Transaction{}).Where("booking_id = ?", booking_id).Updates(&status)
	if record.RowsAffected == 0 {
		log.Warn(record.RowsAffected)
		return entities.Callback{}, record.Error
	}

	return status, nil
}

func (m *transactionModel) GetAllbyUser(role string, user uint, status string, city uint, district uint) []entities.TransactionJoin {
	var transactions []entities.TransactionJoin
	var filter string
	var query string
	if role == "customer" {
		filter += fmt.Sprintf(" t.user_id = %d", user)
	} else if role == "t.consultant" {
		filter += fmt.Sprintf(" t.user_id = %d", user)
	}
	if role != "supervisor" && role != "admin" {
		if status != "" {
			filter += fmt.Sprintf(" AND t.transaction_status = '%s'", status)
		}
		if city != 0 {
			filter += fmt.Sprintf(" AND c.id = %d", city)
		}
		if district != 0 {
			filter += fmt.Sprintf(" AND h.district_id = %d", district)
		}
		query = "select distinct t.id, t.check_in, t.duration, t.booking_id, t.price, t.transaction_status, t.redirect_url, i.url, h.title from transactions as t left join rooms r on r.id = t.room_id left join images i on i.room_id = t.room_id left join houses h on h.id = r.house_id left join districts d on d.id = h.district_id left join cities c on c.id = d.city_id where" + filter
		m.db.Raw(query).Find(&transactions)

		return transactions
	}
	if status == "" && city == 0 && district == 0 {
		filter = ""
	} else {
		filter = "where"
		if status != "" {
			filter += fmt.Sprintf(" t.transaction_status = '%s'", status)
			if city != 0 {
				filter += fmt.Sprintf(" AND c.id = %d", city)
			}
			if district != 0 {
				filter += fmt.Sprintf(" AND h.district_id = %d", district)
			}
		} else if city != 0 {
			filter += fmt.Sprintf(" c.id = %d", city)
			if district != 0 {
				filter += fmt.Sprintf(" AND h.district_id = %d", district)
			}
		} else {
			filter += fmt.Sprintf(" h.district_id = %d", district)
		}
	}

	query = "select distinct t.id, t.check_in, t.duration, t.booking_id, t.price, t.transaction_status, t.redirect_url,t.pdf_invoices_url, i.url, h.title from transactions as t left join rooms r on r.id = t.room_id left join images i on i.room_id = t.room_id left join houses h on h.id = r.house_id left join districts d on d.id = h.district_id left join cities c on c.id = d.city_id " + filter
	m.db.Raw(query).Find(&transactions)

	return transactions
}

func (m *transactionModel) GetAllbyKost(duration int, status string, name string) []entities.TransactionKost {
	var transactions []entities.TransactionKost

	query := "select distinct t.booking_id, u.name, t.check_in, t.duration, t.price, t.transaction_status, t.redirect_url, t.created_at from transactions as t left join rooms r on r.id = t.room_id left join houses h on h.id = r.house_id left join users u on u.id = t.user_id where t.duration = ? and t.transaction_status = ? OR h.title LIKE ?"

	m.db.Raw(query, duration, status, "%"+name+"%").Scan(&transactions)

	return transactions
}

func (m *transactionModel) GetTransactionByBookingID(BookingID string) (entities.DataReminder, error) {
	var response entities.DataReminder
	if err := m.db.Raw("select t.booking_id, u.name, u.email, h.title, t.duration, t.price, t.redirect_url from transactions as t left join houses h on h.id = t.house_id left join users u on u.id = t.user_id  where t.booking_id = ?", BookingID).First(&response).Error; err != nil {
		log.Warn(err)
		return response, err
	}
	return response, nil
}
