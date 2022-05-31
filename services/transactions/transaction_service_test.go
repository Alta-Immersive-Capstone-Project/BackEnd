package transactions_test

import (
	"errors"
	"kost/entities"
	repoMock "kost/mocks/repositories/transactions"
	"kost/services/transactions"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestAddComment(t *testing.T) {
	repo := repoMock.NewTransactionModel(t)
	insertData := entities.TransactionRequest{RoomID: 1, CheckinDate: 1653825446724, RentDuration: 7, TotalBill: 100000, PaymentMethod: "gopay"}
	returnData := entities.Transaction{Model: gorm.Model{ID: uint(1)}, UserID: 1, RoomID: 1, CheckinDate: time.Now(), RentDuration: 7, TotalBill: 100000, PaymentMethod: "gopay"}

	t.Run("Success Insert", func(t *testing.T) {
		repo.On("Create", mock.Anything).Return(returnData, nil).Once()
		srv := transactions.NewTransactionService(repo)

		res, err := srv.AddTransaction(1, insertData)
		assert.NoError(t, err)
		assert.Equal(t, returnData.ID, res.ID)
		assert.Equal(t, returnData.RoomID, res.RoomID)
		repo.AssertExpectations(t)
	})

	t.Run("Error Insert", func(t *testing.T) {
		repo.On("Create", mock.Anything).Return(entities.Transaction{}, errors.New("there is some error")).Once()
		srv := transactions.NewTransactionService(repo)

		_, err := srv.AddTransaction(1, insertData)
		assert.Error(t, err)
		assert.EqualError(t, err, "there is some error")
		repo.AssertExpectations(t)
	})
}

func TestGetTransaction(t *testing.T) {
	repo := new(repoMock.TransactionModel)
	insert := "DM-1653825446724"
	returnData := entities.Transaction{Model: gorm.Model{ID: uint(1)}, BookingID: "DM-1653825446724", UserID: 1, RoomID: 1, CheckinDate: time.Now(), RentDuration: 7, TotalBill: 100000, PaymentMethod: "gopay"}

	t.Run("Success Get", func(t *testing.T) {
		repo.On("Get", insert).Return(returnData, nil).Once()
		srv := transactions.NewTransactionService(repo)

		res, err := srv.GetTransaction(insert)
		assert.NoError(t, err)
		assert.Equal(t, returnData.ID, res.ID)
		assert.Equal(t, returnData.BookingID, res.BookingID)
		assert.Equal(t, returnData.RoomID, res.RoomID)
		repo.AssertExpectations(t)
	})

	t.Run("Error Get", func(t *testing.T) {
		repo.On("Get", insert).Return(entities.Transaction{}, errors.New("data not found")).Once()
		srv := transactions.NewTransactionService(repo)

		res, err := srv.GetTransaction(insert)
		assert.Error(t, err)
		assert.EqualError(t, err, "data not found")
		assert.Equal(t, entities.TransactionResponse{}, res)
		repo.AssertExpectations(t)
	})
}

func TestGetByCustomer(t *testing.T) {
	repo := new(repoMock.TransactionModel)
	insert, status := uint(1), "pending"
	returnData := []entities.Transaction{{Model: gorm.Model{ID: uint(1)}, BookingID: "DM-1653825446724", UserID: 1, RoomID: 1, CheckinDate: time.Now(), RentDuration: 7, TotalBill: 100000, PaymentMethod: "gopay"}}

	t.Run("Success Get All", func(t *testing.T) {
		repo.On("GetAllbyCustomer", insert, status).Return(returnData, nil).Once()
		srv := transactions.NewTransactionService(repo)

		res := srv.GetAllTransactionbyCustomer(insert, status)
		assert.Equal(t, returnData[0].ID, res[0].ID)
		assert.Equal(t, returnData[0].RoomID, res[0].RoomID)
		repo.AssertExpectations(t)
	})
}

func TestGetByConsultant(t *testing.T) {
	repo := new(repoMock.TransactionModel)
	returnData := []entities.Transaction{{Model: gorm.Model{ID: uint(1)}, BookingID: "DM-1653825446724", UserID: 1, RoomID: 1, CheckinDate: time.Now(), RentDuration: 7, TotalBill: 100000, PaymentMethod: "gopay"}}

	t.Run("Success Get All", func(t *testing.T) {
		repo.On("GetAllbyConsultant").Return(returnData, nil).Once()
		srv := transactions.NewTransactionService(repo)

		res := srv.GetAllTransactionbyConsultant()
		assert.Equal(t, returnData[0].ID, res[0].ID)
		assert.Equal(t, returnData[0].RoomID, res[0].RoomID)
		repo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	repo := repoMock.NewTransactionModel(t)
	insertData := entities.TransactionUpdateRequest{TotalBill: 100000}
	returnData := entities.Transaction{Model: gorm.Model{ID: uint(1), UpdatedAt: time.Now()}, UserID: 1, ConsultantID: 1, RoomID: 1, CheckinDate: time.Now(), RentDuration: 7, TotalBill: 100000, PaymentMethod: "gopay"}

	t.Run("Success Update", func(t *testing.T) {
		repo.On("Update", "DM-123", mock.Anything).Return(returnData, nil).Once()
		srv := transactions.NewTransactionService(repo)

		res, err := srv.UpdateTransaction(1, "DM-123", insertData)
		assert.NoError(t, err)
		assert.Equal(t, returnData.TotalBill, res.TotalBill)
		assert.Equal(t, returnData.UpdatedAt, res.UpdatedAt)
		repo.AssertExpectations(t)
	})

	t.Run("Error Insert", func(t *testing.T) {
		repo.On("Update", "DM-123", mock.Anything).Return(entities.Transaction{}, errors.New("there is some error")).Once()
		srv := transactions.NewTransactionService(repo)

		_, err := srv.UpdateTransaction(1, "DM-123", insertData)
		assert.Error(t, err)
		assert.EqualError(t, err, "there is some error")
		repo.AssertExpectations(t)
	})
}
