package transactions_test

// import (
// 	"errors"
// 	"kost/entities"
// 	repoMock "kost/mocks/repositories/transactions"
// 	"kost/services/transactions"
// 	"testing"
// 	"time"

// 	"github.com/midtrans/midtrans-go/snap"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// 	"gorm.io/gorm"
// )

// func TestAddComment(t *testing.T) {
// 	repo := repoMock.NewTransactionModel(t)

// 	requestData := entities.TransactionRequest{HouseID: uint(1), RoomID: uint(1), CheckIn: 1653825446724, Duration: 7, Price: 100000}
// 	insertData := entities.Transaction{BookingID: "LK-1117-1653825446724", UserID: uint(1), HouseID: uint(1), RoomID: uint(1), CheckIn: time.Unix(0, 1653825446724*int64(time.Millisecond)), Duration: 7, Price: 100000, TransactionStatus: "processing"}
// 	returnData := entities.TransactionResponse{BookingID: "LK-1117-1653825446724", Name: "test", Email: "test@test", Phone: "0812", Title: "testing", Address: "test", Price: 100000, CheckIn: time.Now(), Duration: 7, CreatedAt: time.Now()}

// 	t.Run("Success Insert", func(t *testing.T) {
// 		repo.On("Create", insertData).Return(insertData, nil).Once()
// 		repo.On("Request", "LK-1117-1653825446724").Return(returnData, nil).Once()
// 		srv := transactions.NewTransactionService(repo)

// 		res, err := srv.CreateTransaction(uint(1), requestData)
// 		assert.NoError(t, err)
// 		assert.Equal(t, returnData.BookingID, res.BookingID)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("Error Insert Create", func(t *testing.T) {
// 		repo.On("Create", insertData).Return(entities.Transaction{}, errors.New("there is some error")).Once()
// 		repo.On("Request", "LK-1117-1653825446724").Return(entities.TransactionResponse{}, errors.New("there are some error"))
// 		srv := transactions.NewTransactionService(repo)

// 		_, err := srv.CreateTransaction(uint(1), requestData)
// 		assert.Error(t, err)
// 		assert.EqualError(t, err, "there is some error")
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("Error Insert Request", func(t *testing.T) {
// 		repo.On("Create", insertData).Return(insertData, nil).Once()
// 		repo.On("Request", "LK-1117-1653825446724").Return(entities.TransactionResponse{}, errors.New("there are some error"))
// 		srv := transactions.NewTransactionService(repo)

// 		_, err := srv.CreateTransaction(uint(1), requestData)
// 		assert.Error(t, err)
// 		assert.EqualError(t, err, "there are some error")
// 		repo.AssertExpectations(t)
// 	})
// }

// func TestUpdateTransaction(t *testing.T) {
// 	repo := repoMock.NewTransactionModel(t)

// 	updateData := entities.TransactionUpdateRequest{Price: 100000}
// 	insertData := entities.TransactionResponse{BookingID: "LK-1117-1653825446724", Name: "test", Email: "test@test", Phone: "0812", Title: "testing", Address: "test", Price: 100000, CheckIn: time.Now(), Duration: 7, CreatedAt: time.Now()}
// 	returnData := entities.Transaction{BookingID: "LK-1117-1653825446724", ConsultantID: 1, Duration: 7, Price: 100000, TransactionStatus: "pending", RedirectURL: "http://localhost:8080/transactions/pending"}
// 	returnSnap := &snap.Response{Token: "token", RedirectURL: "http://localhost:8080/transactions/pending", StatusCode: "200"}

// 	t.Run("Success Update", func(t *testing.T) {
// 		repo.On("Request", "LK-1117-1653825446724").Return(insertData, nil).Once()
// 		repo.On("CreateSnap", mock.Anything).Return(returnSnap, nil).Once()
// 		repo.On("Update", "LK-1117-1653825446724", returnData).Return(returnData, nil).Once()
// 		srv := transactions.NewTransactionService(repo)

// 		res, err := srv.UpdateTransaction(1, "LK-1117-1653825446724", updateData)
// 		assert.NoError(t, err)
// 		assert.Equal(t, returnData.Price, res.Price)
// 		assert.Equal(t, returnData.UpdatedAt, res.UpdatedAt)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("Error Update Request", func(t *testing.T) {
// 		repo.On("Request", mock.Anything).Return(entities.TransactionResponse{}, errors.New("there is some error")).Once()
// 		repo.On("CreateSnap", mock.Anything).Return(returnSnap, nil)
// 		repo.On("Update", "LK-1117-1653825446724", returnData).Return(returnData, nil)
// 		srv := transactions.NewTransactionService(repo)

// 		_, err := srv.UpdateTransaction(1, mock.Anything, updateData)
// 		assert.Error(t, err)
// 		assert.EqualError(t, err, "there is some error")
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("Error Update CreateSnap", func(t *testing.T) {
// 		repo.On("Request", "LK-1117-1653825446724").Return(insertData, nil).Once()
// 		repo.On("CreateSnap", mock.Anything).Return(&snap.Response{}, errors.New("there is some error"))
// 		repo.On("Update", "LK-1117-1653825446724", mock.Anything).Return(entities.Transaction{}, errors.New("there is some error"))
// 		srv := transactions.NewTransactionService(repo)

// 		_, err := srv.UpdateTransaction(1, mock.Anything, updateData)
// 		assert.Error(t, err)
// 		assert.EqualError(t, err, "there is some error")
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("Error Update", func(t *testing.T) {
// 		repo.On("Request", "LK-1117-1653825446724").Return(insertData, nil).Once()
// 		repo.On("CreateSnap", mock.Anything).Return(returnSnap, nil)
// 		repo.On("Update", mock.Anything, mock.Anything).Return(entities.Transaction{}, errors.New("there is some error"))
// 		srv := transactions.NewTransactionService(repo)

// 		_, err := srv.UpdateTransaction(1, mock.Anything, updateData)
// 		assert.Error(t, err)
// 		assert.EqualError(t, err, "there is some error")
// 		repo.AssertExpectations(t)
// 	})
// }

// func TestGetByCustomer(t *testing.T) {
// 	repo := repoMock.NewTransactionModel(t)
// 	role, user, status, city, district := "customer", uint(1), "active", uint(1), uint(1)
// 	returnData := []entities.TransactionJoin{{BookingID: "DM-1653825446724", CheckIn: time.Now(), Duration: 7, Price: 100000, TransactionStatus: status, Url: "http://localhost:8080/transactions/1", Title: "Kacau"}}

// 	t.Run("Success Get All", func(t *testing.T) {
// 		repo.On("GetAllbyUser", mock.Anything, user, mock.Anything, city, district).Return(returnData, nil).Once()
// 		srv := transactions.NewTransactionService(repo)

// 		res := srv.GetAllTransactionbyUser(role, user, status, city, district)
// 		assert.Equal(t, returnData[0].BookingID, res[0].BookingID)
// 		assert.Equal(t, returnData[0].CheckIn, res[0].CheckIn)
// 		repo.AssertExpectations(t)
// 	})
// }

// func TestGetByConsultant(t *testing.T) {
// 	repo := repoMock.NewTransactionModel(t)

// 	returnData := []entities.Transaction{{Model: gorm.Model{ID: uint(1)}, BookingID: "DM-1653825446724", UserID: 1, HouseID: 1, RoomID: 1, CheckIn: time.Now(), Duration: 7, Price: 100000, TransactionStatus: "processing"}}
// 	returnRequest := entities.TransactionResponse{BookingID: "DM-1653825446724", Name: "test", Email: "test@test", Phone: "0812", Title: "testing", Address: "test", Price: 100000, CheckIn: time.Now(), Duration: 7, CreatedAt: time.Now()}

// 	t.Run("Success Get All", func(t *testing.T) {
// 		repo.On("GetAllbyConsultant").Return(returnData, nil).Once()
// 		repo.On("Request", "DM-1653825446724").Return(returnRequest, nil).Once()
// 		srv := transactions.NewTransactionService(repo)

// 		res := srv.GetAllTransactionbyConsultant()
// 		assert.Equal(t, returnData[0].BookingID, res[0].BookingID)
// 		assert.Equal(t, returnData[0].Price, res[0].Price)
// 		repo.AssertExpectations(t)
// 	})
// }

// func TestUpdateCallback(t *testing.T) {
// 	repo := repoMock.NewTransactionModel(t)

// 	returnData := entities.Callback{OrderID: "DM-1653825446724", StatusCode: "200", SignatureKey: "signature", GrossAmount: "10000.00", PaymentType: "credit_card", TransactionID: "123456789", TransactionStatus: "success", ApprovalCode: "", FraudStatus: ""}

// 	t.Run("Success Callback", func(t *testing.T) {
// 		repo.On("UpdateSnap", "DM-1653825446724", returnData).Return(returnData, nil).Once()
// 		srv := transactions.NewTransactionService(repo)

// 		res, err := srv.UpdateCallback(returnData)
// 		assert.NoError(t, err)
// 		assert.Equal(t, returnData.OrderID, res.OrderID)
// 		assert.Equal(t, returnData.GrossAmount, res.GrossAmount)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("Error Callback", func(t *testing.T) {
// 		repo.On("UpdateSnap", "DM-1653825446724", returnData).Return(entities.Callback{}, errors.New("there are some error")).Once()
// 		srv := transactions.NewTransactionService(repo)

// 		_, err := srv.UpdateCallback(returnData)
// 		assert.Error(t, err)
// 		assert.EqualError(t, err, "there are some error")
// 		repo.AssertExpectations(t)
// 	})
// }

// func TestGetByKost(t *testing.T) {
// 	repo := repoMock.NewTransactionModel(t)

// 	duration, status, name := 7, "active", "mock"
// 	returnData := []entities.TransactionKost{{BookingID: "DM-1653825446724", Name: "mock", Duration: 7, Price: 100000, TransactionStatus: "status", RedirectURL: "http://localhost:8080/transactions/1", CreatedAt: time.Now()}}

// 	t.Run("Success Get All", func(t *testing.T) {
// 		repo.On("GetAllbyKost", duration, mock.Anything, mock.Anything).Return(returnData, nil).Once()
// 		srv := transactions.NewTransactionService(repo)

// 		res := srv.GetAllTransactionbyKost(duration, status, name)
// 		assert.Equal(t, returnData[0].BookingID, res[0].BookingID)
// 		assert.Equal(t, returnData[0].Name, res[0].Name)
// 		repo.AssertExpectations(t)
// 	})
// }
