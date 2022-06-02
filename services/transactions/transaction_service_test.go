package transactions_test

// import (
// 	"errors"
// 	"kost/entities"
// 	mockHouse "kost/mocks/repositories/house"
// 	repoMock "kost/mocks/repositories/transactions"
// 	mockUser "kost/mocks/repositories/user"
// 	"kost/services/transactions"
// 	"testing"
// 	"time"

// 	"github.com/midtrans/midtrans-go/snap"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// 	"gorm.io/gorm"
// )

// func TestAddComment(t *testing.T) {
// 	repo, repoUser, repoHouse := repoMock.NewTransactionModel(t), mockUser.NewUserRepositoryInterface(t), mockHouse.NewIRepoHouse(t)

// 	insertData := entities.TransactionRequest{RoomID: 1, HouseID: 1, CheckinDate: 1653825446724, RentDuration: 7, TotalBill: 100000}
// 	returnUser := entities.User{Model: gorm.Model{ID: uint(1)}, Name: "mock", Email: "mock@gmail.com", Phone: "081234567890", Role: "customer"}
// 	returnHouse := entities.House{Model: gorm.Model{ID: uint(1)}, Title: "Kacau", Brief: "Kecamatan Kacau"}
// 	returnSnap := &snap.Response{Token: "token", RedirectURL: "http://localhost:8080/transactions/callback", StatusCode: "200"}
// 	insert := entities.Transaction{UserID: uint(1), RoomID: uint(1), HouseID: uint(1), CheckinDate: time.Unix(0, 1653825446724*int64(time.Millisecond)), RentDuration: 7, BookingID: "DM-1653825446724", TotalBill: 100000, TransactionStatus: "processing", Token: "token"}
// 	returnData := entities.Transaction{Model: gorm.Model{ID: uint(1)}, UserID: uint(1), RoomID: 1, HouseID: uint(1), CheckinDate: time.Now(), RentDuration: 7, BookingID: "DM-1653825446724", TotalBill: 100000, TransactionStatus: "processing", Token: "token"}

// 	t.Run("Success Insert", func(t *testing.T) {
// 		repoUser.On("GetUserID", uint(1)).Return(returnUser, nil).Once()
// 		repoHouse.On("GetHouseID", uint(1)).Return(returnHouse, nil).Once()
// 		repo.On("CreateSnap", mock.Anything).Return(returnSnap, nil).Once()
// 		repo.On("Create", insert).Return(returnData, nil).Once()
// 		srv := transactions.NewTransactionService(repo, repoUser, repoHouse)

// 		res, err := srv.AddTransaction(uint(1), insertData)
// 		assert.NoError(t, err)
// 		assert.Equal(t, returnUser.Name, res.Name)
// 		assert.Equal(t, returnHouse.Title, res.Title)
// 		assert.Equal(t, returnData.ID, res.ID)
// 		assert.Equal(t, returnData.RoomID, res.RoomID)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("Error Insert", func(t *testing.T) {
// 		repoUser.On("GetUserID", uint(1)).Return(returnUser, nil).Once()
// 		repoHouse.On("GetHouseID", uint(1)).Return(returnHouse, nil).Once()
// 		repo.On("CreateSnap", mock.Anything).Return(nil, errors.New("there is some error")).Once()
// 		// repo.On("Create", mock.Anything).Return(entities.TransactionResponse{}, errors.New("there is some error")).Once()
// 		srv := transactions.NewTransactionService(repo, repoUser, repoHouse)

// 		_, err := srv.AddTransaction(1, insertData)
// 		assert.Error(t, err)
// 		assert.EqualError(t, err, "there is some error")
// 		repo.AssertExpectations(t)
// 	})
// }

// func TestGetTransaction(t *testing.T) {
// 	repo, repoUser, repoHouse := repoMock.NewTransactionModel(t), mockUser.NewUserRepositoryInterface(t), mockHouse.NewIRepoHouse(t)

// 	insert := "DM-1653825446724"
// 	returnData := entities.Transaction{Model: gorm.Model{ID: uint(1)}, BookingID: "DM-1653825446724", UserID: 1, RoomID: 1, CheckinDate: time.Now(), RentDuration: 7, TotalBill: 100000, TransactionStatus: "processing", Token: "token"}

// 	t.Run("Success Get", func(t *testing.T) {
// 		repo.On("Get", insert).Return(returnData, nil).Once()
// 		srv := transactions.NewTransactionService(repo, repoUser, repoHouse)

// 		res, err := srv.GetTransaction(insert)
// 		assert.NoError(t, err)
// 		assert.Equal(t, returnData.ID, res.ID)
// 		assert.Equal(t, returnData.BookingID, res.BookingID)
// 		assert.Equal(t, returnData.RoomID, res.RoomID)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("Error Get", func(t *testing.T) {
// 		repo.On("Get", insert).Return(entities.Transaction{}, errors.New("data not found")).Once()
// 		srv := transactions.NewTransactionService(repo, repoUser, repoHouse)

// 		res, err := srv.GetTransaction(insert)
// 		assert.Error(t, err)
// 		assert.EqualError(t, err, "data not found")
// 		assert.Equal(t, entities.TransactionResponse{}, res)
// 		repo.AssertExpectations(t)
// 	})
// }

// func TestGetByCustomer(t *testing.T) {
// 	repo, repoUser, repoHouse := repoMock.NewTransactionModel(t), mockUser.NewUserRepositoryInterface(t), mockHouse.NewIRepoHouse(t)
// 	role, user, status, city, district := "customer", uint(1), "active", uint(1), uint(1)
// 	returnData := []entities.TransactionJoin{{BookingID: "DM-1653825446724", CheckinDate: time.Now(), RentDuration: 7, TotalBill: 100000, TransactionStatus: status, Url: "http://localhost:8080/transactions/1", Title: "Kacau"}}

// 	t.Run("Success Get All", func(t *testing.T) {
// 		repo.On("GetAllbyCustomer", mock.Anything, user, mock.Anything, city, district).Return(returnData, nil).Once()
// 		srv := transactions.NewTransactionService(repo, repoUser, repoHouse)

// 		res := srv.GetAllTransactionbyCustomer(role, user, status, city, district)
// 		assert.Equal(t, returnData[0].BookingID, res[0].BookingID)
// 		assert.Equal(t, returnData[0].CheckinDate, res[0].CheckinDate)
// 		repo.AssertExpectations(t)
// 	})
// }

// func TestGetByConsultant(t *testing.T) {
// 	repo, repoUser, repoHouse := repoMock.NewTransactionModel(t), mockUser.NewUserRepositoryInterface(t), mockHouse.NewIRepoHouse(t)

// 	returnData := []entities.Transaction{{Model: gorm.Model{ID: uint(1)}, BookingID: "DM-1653825446724", UserID: 1, RoomID: 1, CheckinDate: time.Now(), RentDuration: 7, TotalBill: 100000}}

// 	t.Run("Success Get All", func(t *testing.T) {
// 		repo.On("GetAllbyConsultant").Return(returnData, nil).Once()
// 		srv := transactions.NewTransactionService(repo, repoUser, repoHouse)

// 		res := srv.GetAllTransactionbyConsultant()
// 		assert.Equal(t, returnData[0].ID, res[0].ID)
// 		assert.Equal(t, returnData[0].RoomID, res[0].RoomID)
// 		repo.AssertExpectations(t)
// 	})
// }

// // func TestUpdateStatus(t *testing.T) {
// // 	repo, repoUser, repoHouse := repoMock.NewTransactionModel(t), mockUser.NewUserRepositoryInterface(t), mockHouse.NewIRepoHouse(t)

// // 	returnData := entities.Callback{OrderID: "DM-1653825446724", StatusCode: "200", SignatureKey: "signature", GrossAmount: "10000.00", PaymentType: "credit_card", TransactionID: "123456789", TransactionStatus: "success", ApprovalCode: "", FraudStatus: ""}

// // 	t.Run("Success Error", func(t *testing.T) {
// // 		repo.On("UpdateStatus", "DM-1653825446724", returnData).Once()
// // 		srv := transactions.NewTransactionService(repo, repoUser, repoHouse)

// // 		err := srv.UpdateStatus(returnData)
// // 		assert.NoError(t, err)
// // 		repo.AssertExpectations(t)
// // 	})
// // }

// func TestUpdate(t *testing.T) {
// 	repo, repoUser, repoHouse := repoMock.NewTransactionModel(t), mockUser.NewUserRepositoryInterface(t), mockHouse.NewIRepoHouse(t)

// 	insertData := entities.TransactionUpdateRequest{TotalBill: 100000}
// 	returnData := entities.Transaction{Model: gorm.Model{ID: uint(1), UpdatedAt: time.Now()}, UserID: 1, ConsultantID: 1, RoomID: 1, CheckinDate: time.Now(), RentDuration: 7, TotalBill: 100000}

// 	t.Run("Success Update", func(t *testing.T) {
// 		repo.On("Update", "DM-123", mock.Anything).Return(returnData, nil).Once()
// 		srv := transactions.NewTransactionService(repo, repoUser, repoHouse)

// 		res, err := srv.UpdateTransaction(1, "DM-123", insertData)
// 		assert.NoError(t, err)
// 		assert.Equal(t, returnData.TotalBill, res.TotalBill)
// 		assert.Equal(t, returnData.UpdatedAt, res.UpdatedAt)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("Error Insert", func(t *testing.T) {
// 		repo.On("Update", "DM-123", mock.Anything).Return(entities.Transaction{}, errors.New("there is some error")).Once()
// 		srv := transactions.NewTransactionService(repo, repoUser, repoHouse)

// 		_, err := srv.UpdateTransaction(1, "DM-123", insertData)
// 		assert.Error(t, err)
// 		assert.EqualError(t, err, "there is some error")
// 		repo.AssertExpectations(t)
// 	})
// }

// func TestGetByKost(t *testing.T) {
// 	repo, repoUser, repoHouse := repoMock.NewTransactionModel(t), mockUser.NewUserRepositoryInterface(t), mockHouse.NewIRepoHouse(t)

// 	duration, status, name := 7, "active", "mock"
// 	returnData := []entities.TransactionKost{{BookingID: "DM-1653825446724", Name: "mock", RentDuration: 7, TotalBill: 100000, TransactionStatus: "status"}}

// 	t.Run("Success Get All", func(t *testing.T) {
// 		repo.On("GetAllbyKost", duration, mock.Anything, mock.Anything).Return(returnData, nil).Once()
// 		srv := transactions.NewTransactionService(repo, repoUser, repoHouse)

// 		res := srv.GetAllTransactionbyKost(duration, status, name)
// 		assert.Equal(t, returnData[0].BookingID, res[0].BookingID)
// 		assert.Equal(t, returnData[0].Name, res[0].Name)
// 		repo.AssertExpectations(t)
// 	})
// }
