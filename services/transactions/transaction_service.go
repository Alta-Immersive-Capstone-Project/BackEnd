package transactions

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"kost/configs"
	"kost/entities"
	"kost/repositories/house"
	repo "kost/repositories/transactions"
	"kost/repositories/user"
	"strconv"
	"time"

	"github.com/jinzhu/copier"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type transactionService struct {
	tm repo.TransactionModel
	um user.UserRepositoryInterface
	hm house.IRepoHouse
}

func NewTransactionService(tm repo.TransactionModel, um user.UserRepositoryInterface, hm house.IRepoHouse) *transactionService {
	return &transactionService{
		tm: tm,
		um: um,
		hm: hm,
	}
}

func (ts *transactionService) AddTransaction(customer_id uint, request entities.TransactionRequest) (entities.TransactionResponse, error) {
	var response entities.TransactionResponse
	booking_id := fmt.Sprintf("DM-%d", request.CheckinDate)

	user, _ := ts.um.GetUserID(customer_id)
	house, _ := ts.hm.GetHouseID(request.HouseID)
	snapRequest := &snap.Request{
		CustomerDetail: &midtrans.CustomerDetails{
			FName: user.Name,
			Email: user.Email,
			Phone: user.Phone,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  booking_id,
			GrossAmt: request.TotalBill,
		},
		Items: &[]midtrans.ItemDetails{
			{
				ID:    strconv.Itoa(int(house.ID)),
				Name:  house.Title,
				Price: request.TotalBill,
				Qty:   1,
			},
		},
		Callbacks: &snap.Callbacks{
			Finish: configs.Get().App.BaseURL + "/transactions/callback",
		},
	}

	snap, err := ts.tm.CreateSnap(snapRequest)
	if err != nil {
		return entities.TransactionResponse{}, err
	}

	transaction := entities.Transaction{
		UserID:       customer_id,
		RoomID:       request.RoomID,
		HouseID:      request.HouseID,
		CheckinDate:  time.Unix(0, request.CheckinDate*int64(time.Millisecond)),
		RentDuration: request.RentDuration,
		BookingID:    booking_id,
		TotalBill:    request.TotalBill,
		Status:       "pending",
		Token:        snap.Token,
	}

	result, err := ts.tm.Create(transaction)
	if err != nil {
		return entities.TransactionResponse{}, err
	}

	copier.Copy(&response, &result)
	copier.Copy(&response, &house)
	copier.Copy(&response, &user)
	copier.Copy(&response, &snap)
	return response, nil
}

func (ts *transactionService) UpdateStatus(req entities.Callback) error {
	key := req.OrderID + req.StatusCode + req.GrossAmount + configs.Get().Payment.MidtransServerKey
	hash := sha512.New()
	hash.Write([]byte(key))
	SignatureKey := hex.EncodeToString(hash.Sum(nil))

	fmt.Println(req.OrderID)
	fmt.Println(req.StatusCode)
	fmt.Println(req.GrossAmount)
	fmt.Println(configs.Get().Payment.MidtransServerKey)
	fmt.Println(SignatureKey)

	if req.SignatureKey != SignatureKey {
		return errors.New("you are not allowed to access this resource")
	}

	transaction := entities.Callback{
		TransactionStatus: req.TransactionStatus,
		PaymentType:       req.PaymentType,
	}

	ts.tm.UpdateStatus(req.OrderID, transaction)

	return nil
}

func (ts *transactionService) GetTransaction(booking_id string) (entities.TransactionResponse, error) {
	var response entities.TransactionResponse

	result, err := ts.tm.Get(booking_id)
	if err != nil {
		return entities.TransactionResponse{}, err
	}

	copier.Copy(&response, &result)
	return response, nil
}

func (ts *transactionService) GetAllTransactionbyCustomer(role string, user uint, status string, city uint, district uint) []entities.TransactionJoin {
	response := ts.tm.GetAllbyCustomer(role, user, status, city, district)
	return response
}

func (ts *transactionService) GetAllTransactionbyConsultant() []entities.TransactionResponse {
	var response []entities.TransactionResponse

	results := ts.tm.GetAllbyConsultant()

	for _, r := range results {
		var transaction entities.TransactionResponse
		copier.Copy(&transaction, &r)
		response = append(response, transaction)
	}

	return response
}

func (ts *transactionService) UpdateTransaction(customer_id uint, booking_id string, request entities.TransactionUpdateRequest) (entities.TransactionUpdateResponse, error) {
	var response entities.TransactionUpdateResponse

	transaction := entities.Transaction{
		ConsultantID: customer_id,
		TotalBill:    request.TotalBill,
	}

	result, err := ts.tm.Update(booking_id, transaction)
	if err != nil {
		return entities.TransactionUpdateResponse{}, err
	}

	copier.Copy(&response, &result)
	return response, nil
}

func (ts *transactionService) GetAllTransactionbyKost(duration int, status string, name string) []entities.TransactionKost {
	response := ts.tm.GetAllbyKost(duration, status, name)
	return response
}
