package transactions_test

import (
	"errors"
	"kost/entities"
	invMock "kost/mocks/repositories/invoice"
	repoMock "kost/mocks/repositories/transactions"
	s3Mock "kost/mocks/utils/s3"
	"kost/services/transactions"
	"testing"
	"time"

	"github.com/midtrans/midtrans-go/snap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestAddComment(t *testing.T) {
	repo := repoMock.NewTransactionModel(t)
	invoice := invMock.NewInvoiceModel(t)
	s3 := s3Mock.NewS3Control(t)
	requestData := entities.TransactionRequest{HouseID: uint(1), RoomID: uint(1), CheckIn: 1653825446724, Duration: 7, Price: 100000}
	insertData := entities.Transaction{BookingID: "LK-1117-1653825446724", UserID: uint(1), HouseID: uint(1), RoomID: uint(1), CheckIn: time.Unix(0, 1653825446724*int64(time.Millisecond)), Duration: 7, Price: 100000, TransactionStatus: "processing"}
	returnData := entities.TransactionResponse{BookingID: "LK-1117-1653825446724", Name: "test", Email: "test@test", Phone: "0812", Title: "testing", Address: "test", Price: 100000, CheckIn: time.Now(), Duration: 7, CreatedAt: time.Now()}

	t.Run("Success Insert", func(t *testing.T) {
		repo.On("Create", insertData).Return(insertData, nil).Once()
		repo.On("Request", "LK-1117-1653825446724").Return(returnData, nil).Once()
		srv := transactions.NewTransactionService(repo, invoice, s3)

		res, err := srv.CreateTransaction(uint(1), requestData)
		assert.NoError(t, err)
		assert.Equal(t, returnData.BookingID, res.BookingID)
		repo.AssertExpectations(t)
		invoice.AssertExpectations(t)
		s3.AssertExpectations(t)
	})

	t.Run("Error Insert Create", func(t *testing.T) {
		repo.On("Create", insertData).Return(entities.Transaction{}, errors.New("there is some error")).Once()
		repo.On("Request", "LK-1117-1653825446724").Return(entities.TransactionResponse{}, errors.New("there are some error"))
		srv := transactions.NewTransactionService(repo, invoice, s3)

		_, err := srv.CreateTransaction(uint(1), requestData)
		assert.Error(t, err)
		assert.EqualError(t, err, "there is some error")
		repo.AssertExpectations(t)
		invoice.AssertExpectations(t)
		s3.AssertExpectations(t)
	})

	t.Run("Error Insert Request", func(t *testing.T) {
		repo.On("Create", insertData).Return(insertData, nil).Once()
		repo.On("Request", "LK-1117-1653825446724").Return(entities.TransactionResponse{}, errors.New("there are some error"))
		srv := transactions.NewTransactionService(repo, invoice, s3)

		_, err := srv.CreateTransaction(uint(1), requestData)
		assert.Error(t, err)
		assert.EqualError(t, err, "there are some error")
		repo.AssertExpectations(t)
		invoice.AssertExpectations(t)
		s3.AssertExpectations(t)
	})
}

func TestUpdateTransaction(t *testing.T) {
	repo := repoMock.NewTransactionModel(t)
	invoice := invMock.NewInvoiceModel(t)
	s3 := s3Mock.NewS3Control(t)
	updateData := entities.TransactionUpdateRequest{Price: 100000}
	insertData := entities.DataReminder{BookingID: "LK-1117-1653825446724", Name: "test", Email: "test@test", Phone: "0812", Title: "testing", Address: "test", Price: 100000, CheckIn: time.Now(), Duration: 7}
	returnData := entities.Transaction{BookingID: "LK-1117-1653825446724", ConsultantID: 1, Duration: 7, Price: 100000, TransactionStatus: "pending", RedirectURL: "http://transactions/pending"}
	returnSnap := &snap.Response{Token: "token", RedirectURL: "http://transactions/pending", StatusCode: "200"}

	t.Run("Success Update", func(t *testing.T) {
		repo.On("GetTransactionByBookingID", "LK-1117-1653825446724").Return(insertData, nil).Once()
		repo.On("CreateSnap", mock.Anything).Return(returnSnap, nil).Once()
		repo.On("Update", "LK-1117-1653825446724", mock.Anything).Return(returnData, nil).Once()
		invoice.On("CreateInvoice", "logo.png", mock.Anything).Return("namefile").Once()
		s3.On("UploadInvoiceToS3", "LK-1117-1653825446724", mock.Anything).Return("url", nil).Once()
		repo.On("Update", "LK-1117-1653825446724", mock.Anything).Return(returnData, nil).Once()
		srv := transactions.NewTransactionService(repo, invoice, s3)

		res, err := srv.UpdateTransaction(1, "LK-1117-1653825446724", updateData)
		assert.NoError(t, err)
		assert.Equal(t, returnData.Price, res.Price)
		assert.Equal(t, returnData.UpdatedAt, res.UpdatedAt)
		repo.AssertExpectations(t)
		invoice.AssertExpectations(t)
		s3.AssertExpectations(t)
	})

	t.Run("Error Duplicate", func(t *testing.T) {
		Tes := entities.DataReminder{BookingID: "LK-1117-1653825446724", Name: "test", Email: "test@test", Phone: "0812", Title: "testing", Address: "test", Price: 100000, CheckIn: time.Now(), Duration: 7, RedirectURL: "Test.com"}

		repo.On("GetTransactionByBookingID", "LK-1117-1653825446724").Return(Tes, nil).Once()

		srv := transactions.NewTransactionService(repo, invoice, s3)

		_, err := srv.UpdateTransaction(1, "LK-1117-1653825446724", updateData)
		assert.Error(t, err)

		repo.AssertExpectations(t)
		invoice.AssertExpectations(t)
		s3.AssertExpectations(t)
	})
	t.Run("Error Get Data Request", func(t *testing.T) {
		repo.On("GetTransactionByBookingID", "LK-1117-1653825446724").Return(insertData, errors.New("Error Get Data")).Once()

		srv := transactions.NewTransactionService(repo, invoice, s3)

		_, errGet := srv.UpdateTransaction(1, mock.Anything, updateData)
		assert.Error(t, errGet)
		assert.EqualError(t, errGet, "Booking ID Not Found")
		repo.AssertExpectations(t)
		invoice.AssertExpectations(t)
		s3.AssertExpectations(t)
	})

	t.Run("Error CreateSnap", func(t *testing.T) {
		repo.On("GetTransactionByBookingID", "LK-1117-1653825446724").Return(insertData, nil).Once()
		repo.On("CreateSnap", mock.Anything).Return(returnSnap, errors.New("Error Create Snap")).Once()

		srv := transactions.NewTransactionService(repo, invoice, s3)
		_, errSnap := srv.UpdateTransaction(1, mock.Anything, updateData)
		assert.Error(t, errSnap)
		assert.EqualError(t, errSnap, "Error Create Snap")
		repo.AssertExpectations(t)
		invoice.AssertExpectations(t)
		s3.AssertExpectations(t)
	})

	t.Run("Error Update 1", func(t *testing.T) {
		repo.On("GetTransactionByBookingID", "LK-1117-1653825446724").Return(insertData, nil).Once()
		repo.On("CreateSnap", mock.Anything).Return(returnSnap, nil).Once()
		repo.On("Update", "LK-1117-1653825446724", mock.Anything).Return(returnData, errors.New("Error Update 1")).Once()

		srv := transactions.NewTransactionService(repo, invoice, s3)

		_, errUpd := srv.UpdateTransaction(1, mock.Anything, updateData)
		assert.Error(t, errUpd)
		assert.EqualError(t, errUpd, "Error Update 1")
		repo.AssertExpectations(t)
		invoice.AssertExpectations(t)
		s3.AssertExpectations(t)
	})

	t.Run("Error Upload S3", func(t *testing.T) {
		repo.On("GetTransactionByBookingID", "LK-1117-1653825446724").Return(insertData, nil).Once()
		repo.On("CreateSnap", mock.Anything).Return(returnSnap, nil).Once()
		repo.On("Update", "LK-1117-1653825446724", mock.Anything).Return(returnData, nil).Once()
		invoice.On("CreateInvoice", "logo.png", mock.Anything).Return("namefile").Once()
		s3.On("UploadInvoiceToS3", "LK-1117-1653825446724", mock.Anything).Return("", errors.New("Error Upload S3")).Once()

		srv := transactions.NewTransactionService(repo, invoice, s3)

		_, errUpload := srv.UpdateTransaction(1, mock.Anything, updateData)
		assert.Error(t, errUpload)
		assert.EqualError(t, errUpload, "Error Upload S3")
		repo.AssertExpectations(t)
		invoice.AssertExpectations(t)
		s3.AssertExpectations(t)
	})

	t.Run("Error Update 2", func(t *testing.T) {
		repo.On("GetTransactionByBookingID", "LK-1117-1653825446724").Return(insertData, nil).Once()
		repo.On("CreateSnap", mock.Anything).Return(returnSnap, nil).Once()
		repo.On("Update", "LK-1117-1653825446724", mock.Anything).Return(returnData, nil).Once()
		invoice.On("CreateInvoice", "logo.png", mock.Anything).Return("namefile").Once()
		s3.On("UploadInvoiceToS3", "LK-1117-1653825446724", mock.Anything).Return("url", nil).Once()
		repo.On("Update", "LK-1117-1653825446724", mock.Anything).Return(returnData, errors.New("there is some error")).Once()
		srv := transactions.NewTransactionService(repo, invoice, s3)

		_, errUpd2 := srv.UpdateTransaction(1, mock.Anything, updateData)
		assert.Error(t, errUpd2)
		assert.EqualError(t, errUpd2, "there is some error")
		repo.AssertExpectations(t)
		invoice.AssertExpectations(t)
		s3.AssertExpectations(t)
	})
}

func TestGetByCustomer(t *testing.T) {
	repo := repoMock.NewTransactionModel(t)
	invoice := invMock.NewInvoiceModel(t)
	s3 := s3Mock.NewS3Control(t)
	role, user, status, city, district := "customer", uint(1), "active", uint(1), uint(1)
	returnData := []entities.TransactionJoin{{BookingID: "DM-1653825446724", CheckIn: time.Now(), Duration: 7, Price: 100000, TransactionStatus: status, Url: "http://localhost:8080/transactions/1", Title: "Kacau"}}

	t.Run("Success Get All", func(t *testing.T) {
		repo.On("GetAllbyUser", mock.Anything, user, mock.Anything, city, district).Return(returnData, nil).Once()
		srv := transactions.NewTransactionService(repo, invoice, s3)

		res := srv.GetAllTransactionbyUser(role, user, status, city, district)
		assert.Equal(t, returnData[0].BookingID, res[0].BookingID)
		assert.Equal(t, returnData[0].CheckIn, res[0].CheckIn)
		repo.AssertExpectations(t)
		invoice.AssertExpectations(t)
		s3.AssertExpectations(t)
		invoice.AssertExpectations(t)
		s3.AssertExpectations(t)
	})
}

func TestGetByConsultant(t *testing.T) {
	repo := repoMock.NewTransactionModel(t)
	invoice := invMock.NewInvoiceModel(t)
	s3 := s3Mock.NewS3Control(t)
	returnData := []entities.Transaction{{Model: gorm.Model{ID: uint(1)}, BookingID: "DM-1653825446724", UserID: 1, HouseID: 1, RoomID: 1, CheckIn: time.Now(), Duration: 7, Price: 100000, TransactionStatus: "processing"}}
	returnRequest := entities.TransactionResponse{BookingID: "DM-1653825446724", Name: "test", Email: "test@test", Phone: "0812", Title: "testing", Address: "test", Price: 100000, CheckIn: time.Now(), Duration: 7, CreatedAt: time.Now()}

	t.Run("Success Get All", func(t *testing.T) {
		repo.On("GetAllbyConsultant").Return(returnData, nil).Once()
		repo.On("Request", "DM-1653825446724").Return(returnRequest, nil).Once()
		srv := transactions.NewTransactionService(repo, invoice, s3)

		res := srv.GetAllTransactionbyConsultant()
		assert.Equal(t, returnData[0].BookingID, res[0].BookingID)
		assert.Equal(t, returnData[0].Price, res[0].Price)
		repo.AssertExpectations(t)
		invoice.AssertExpectations(t)
		s3.AssertExpectations(t)
	})
}

func TestUpdateCallback(t *testing.T) {
	repo := repoMock.NewTransactionModel(t)
	invoice := invMock.NewInvoiceModel(t)
	s3 := s3Mock.NewS3Control(t)
	returnData := entities.Callback{OrderID: "DM-1653825446724", StatusCode: "200", SignatureKey: "signature", GrossAmount: "10000.00", PaymentType: "credit_card", TransactionID: "123456789", TransactionStatus: "success", ApprovalCode: "", FraudStatus: ""}

	t.Run("Success Callback", func(t *testing.T) {
		repo.On("UpdateSnap", "DM-1653825446724", returnData).Return(returnData, nil).Once()
		srv := transactions.NewTransactionService(repo, invoice, s3)

		res, err := srv.UpdateCallback(returnData)
		assert.NoError(t, err)
		assert.Equal(t, returnData.OrderID, res.OrderID)
		assert.Equal(t, returnData.GrossAmount, res.GrossAmount)
		repo.AssertExpectations(t)
		invoice.AssertExpectations(t)
		s3.AssertExpectations(t)
	})

	t.Run("Error Callback", func(t *testing.T) {
		repo.On("UpdateSnap", "DM-1653825446724", returnData).Return(entities.Callback{}, errors.New("there are some error")).Once()
		srv := transactions.NewTransactionService(repo, invoice, s3)

		_, err := srv.UpdateCallback(returnData)
		assert.Error(t, err)
		assert.EqualError(t, err, "there are some error")
		repo.AssertExpectations(t)
		invoice.AssertExpectations(t)
		s3.AssertExpectations(t)
	})
}

func TestGetByKost(t *testing.T) {
	repo := repoMock.NewTransactionModel(t)
	invoice := invMock.NewInvoiceModel(t)
	s3 := s3Mock.NewS3Control(t)
	duration, status, name := 7, "active", "mock"
	returnData := []entities.TransactionKost{{BookingID: "DM-1653825446724", Name: "mock", Duration: 7, Price: 100000, TransactionStatus: "status", RedirectURL: "http:/transactions/1", CreatedAt: time.Now()}}

	t.Run("Success Get All", func(t *testing.T) {
		repo.On("GetAllbyKost", duration, mock.Anything, mock.Anything).Return(returnData, nil).Once()
		srv := transactions.NewTransactionService(repo, invoice, s3)

		res := srv.GetAllTransactionbyKost(duration, status, name)
		assert.Equal(t, returnData[0].BookingID, res[0].BookingID)
		assert.Equal(t, returnData[0].Name, res[0].Name)
		repo.AssertExpectations(t)
		invoice.AssertExpectations(t)
		s3.AssertExpectations(t)
	})
}

func TestGetReport(t *testing.T) {
	repo := repoMock.NewTransactionModel(t)
	invoice := invMock.NewInvoiceModel(t)
	s3 := s3Mock.NewS3Control(t)
	returnData := []entities.TransactionKost{{BookingID: "DM-1653825446724", Name: "mock", Duration: 7, Price: 100000, TransactionStatus: "status", RedirectURL: "http://localhost:8080/transactions/1", CreatedAt: time.Date(2022, 06, 06, 0, 0, 0, 0, time.UTC)}}
	url := "http://transactions/pending"
	t.Run("Success Get Report", func(t *testing.T) {
		invoice.On("CreateReport", "logo.png", mock.Anything).Return("Success").Once()
		s3.On("UploadInvoiceToS3", "Success", "Success.pdf").Return(url, nil).Once()

		srv := transactions.NewTransactionService(repo, invoice, s3)
		res := srv.GetReport(returnData)
		assert.Equal(t, url, res)
		repo.AssertExpectations(t)
		invoice.AssertExpectations(t)
		s3.AssertExpectations(t)
	})
	t.Run("Success Error Generate", func(t *testing.T) {
		invoice.On("CreateReport", "logo.png", mock.Anything).Return("").Once()

		srv := transactions.NewTransactionService(repo, invoice, s3)
		res := srv.GetReport(returnData)
		assert.Equal(t, "GAGAL GENERATE REPORT", res)
		repo.AssertExpectations(t)
		invoice.AssertExpectations(t)
		s3.AssertExpectations(t)
	})
	t.Run("Success Error Upload", func(t *testing.T) {
		invoice.On("CreateReport", "logo.png", mock.Anything).Return("Test")
		s3.On("UploadInvoiceToS3", "Test", "Test.pdf").Return("", errors.New("Error Upload")).Once()

		srv := transactions.NewTransactionService(repo, invoice, s3)
		err := srv.GetReport(returnData)
		assert.Equal(t, "", err)
		repo.AssertExpectations(t)
		invoice.AssertExpectations(t)
		s3.AssertExpectations(t)
	})
}
