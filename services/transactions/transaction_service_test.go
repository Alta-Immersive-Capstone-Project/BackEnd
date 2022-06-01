package transactions_test

// func TestAddComment(t *testing.T) {
// 	repo := repoMock.NewTransactionModel(t)
// 	insertData := entities.TransactionRequest{RoomID: 1, CheckinDate: 1653825446724, RentDuration: 7, TotalBill: 100000, PaymentMethod: "gopay"}
// 	returnData := entities.Transaction{Model: gorm.Model{ID: uint(1)}, UserID: 1, RoomID: 1, CheckinDate: time.Now(), RentDuration: 7, TotalBill: 100000, PaymentMethod: "gopay"}

// 	t.Run("Success Insert", func(t *testing.T) {
// 		repo.On("Create", mock.Anything).Return(returnData, nil).Once()
// 		srv := transactions.NewTransactionService(repo)

// 		res, err := srv.AddTransaction(1, insertData)
// 		assert.NoError(t, err)
// 		assert.Equal(t, returnData.ID, res.ID)
// 		assert.Equal(t, returnData.RoomID, res.RoomID)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("Error Insert", func(t *testing.T) {
// 		repo.On("Create", mock.Anything).Return(entities.Transaction{}, errors.New("there is some error")).Once()
// 		srv := transactions.NewTransactionService(repo)

// 		_, err := srv.AddTransaction(1, insertData)
// 		assert.Error(t, err)
// 		assert.EqualError(t, err, "there is some error")
// 		repo.AssertExpectations(t)
// 	})
// }

// func TestGetTransaction(t *testing.T) {
// 	repo := new(repoMock.TransactionModel)
// 	insert := "DM-1653825446724"
// 	returnData := entities.Transaction{Model: gorm.Model{ID: uint(1)}, BookingID: "DM-1653825446724", UserID: 1, RoomID: 1, CheckinDate: time.Now(), RentDuration: 7, TotalBill: 100000, PaymentMethod: "gopay"}

// 	t.Run("Success Get", func(t *testing.T) {
// 		repo.On("Get", insert).Return(returnData, nil).Once()
// 		srv := transactions.NewTransactionService(repo)

// 		res, err := srv.GetTransaction(insert)
// 		assert.NoError(t, err)
// 		assert.Equal(t, returnData.ID, res.ID)
// 		assert.Equal(t, returnData.BookingID, res.BookingID)
// 		assert.Equal(t, returnData.RoomID, res.RoomID)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("Error Get", func(t *testing.T) {
// 		repo.On("Get", insert).Return(entities.Transaction{}, errors.New("data not found")).Once()
// 		srv := transactions.NewTransactionService(repo)

// 		res, err := srv.GetTransaction(insert)
// 		assert.Error(t, err)
// 		assert.EqualError(t, err, "data not found")
// 		assert.Equal(t, entities.TransactionResponse{}, res)
// 		repo.AssertExpectations(t)
// 	})
// }

// func TestGetByCustomer(t *testing.T) {
// 	repo := new(repoMock.TransactionModel)
// 	role, user, status, city, district := "customer", uint(1), "active", uint(1), uint(1)
// 	returnData := []entities.TransactionJoin{{BookingID: "DM-1653825446724", CheckinDate: time.Now(), RentDuration: 7, TotalBill: 100000, Status: status, Url: "http://localhost:8080/transactions/1", Title: "Kacau"}}

// 	t.Run("Success Get All", func(t *testing.T) {
// 		repo.On("GetAllbyCustomer", mock.Anything, user, mock.Anything, city, district).Return(returnData, nil).Once()
// 		srv := transactions.NewTransactionService(repo)

// 		res := srv.GetAllTransactionbyCustomer(role, user, status, city, district)
// 		assert.Equal(t, returnData[0].BookingID, res[0].BookingID)
// 		assert.Equal(t, returnData[0].CheckinDate, res[0].CheckinDate)
// 		repo.AssertExpectations(t)
// 	})
// }

// func TestGetByConsultant(t *testing.T) {
// 	repo := new(repoMock.TransactionModel)
// 	returnData := []entities.Transaction{{Model: gorm.Model{ID: uint(1)}, BookingID: "DM-1653825446724", UserID: 1, RoomID: 1, CheckinDate: time.Now(), RentDuration: 7, TotalBill: 100000, PaymentMethod: "gopay"}}

// 	t.Run("Success Get All", func(t *testing.T) {
// 		repo.On("GetAllbyConsultant").Return(returnData, nil).Once()
// 		srv := transactions.NewTransactionService(repo)

// 		res := srv.GetAllTransactionbyConsultant()
// 		assert.Equal(t, returnData[0].ID, res[0].ID)
// 		assert.Equal(t, returnData[0].RoomID, res[0].RoomID)
// 		repo.AssertExpectations(t)
// 	})
// }

// func TestUpdate(t *testing.T) {
// 	repo := repoMock.NewTransactionModel(t)
// 	insertData := entities.TransactionUpdateRequest{TotalBill: 100000}
// 	returnData := entities.Transaction{Model: gorm.Model{ID: uint(1), UpdatedAt: time.Now()}, UserID: 1, ConsultantID: 1, RoomID: 1, CheckinDate: time.Now(), RentDuration: 7, TotalBill: 100000, PaymentMethod: "gopay"}

// 	t.Run("Success Update", func(t *testing.T) {
// 		repo.On("Update", "DM-123", mock.Anything).Return(returnData, nil).Once()
// 		srv := transactions.NewTransactionService(repo)

// 		res, err := srv.UpdateTransaction(1, "DM-123", insertData)
// 		assert.NoError(t, err)
// 		assert.Equal(t, returnData.TotalBill, res.TotalBill)
// 		assert.Equal(t, returnData.UpdatedAt, res.UpdatedAt)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("Error Insert", func(t *testing.T) {
// 		repo.On("Update", "DM-123", mock.Anything).Return(entities.Transaction{}, errors.New("there is some error")).Once()
// 		srv := transactions.NewTransactionService(repo)

// 		_, err := srv.UpdateTransaction(1, "DM-123", insertData)
// 		assert.Error(t, err)
// 		assert.EqualError(t, err, "there is some error")
// 		repo.AssertExpectations(t)
// 	})
// }

// func TestGetByKost(t *testing.T) {
// 	repo := new(repoMock.TransactionModel)
// 	duration, status, name := 7, "active", "mock"
// 	returnData := []entities.TransactionKost{{BookingID: "DM-1653825446724", Name: "mock", RentDuration: 7, TotalBill: 100000, Status: "status"}}

// 	t.Run("Success Get All", func(t *testing.T) {
// 		repo.On("GetAllbyKost", duration, mock.Anything, mock.Anything).Return(returnData, nil).Once()
// 		srv := transactions.NewTransactionService(repo)

// 		res := srv.GetAllTransactionbyKost(duration, status, name)
// 		assert.Equal(t, returnData[0].BookingID, res[0].BookingID)
// 		assert.Equal(t, returnData[0].Name, res[0].Name)
// 		repo.AssertExpectations(t)
// 	})
// }
