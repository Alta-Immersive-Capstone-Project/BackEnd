package transactions

import (
	"fmt"
	"kost/configs"
	"kost/entities"
	repo "kost/repositories/transactions"
	"time"

	"github.com/jinzhu/copier"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type transactionService struct {
	tm repo.TransactionModel
}

func NewTransactionService(tm repo.TransactionModel) *transactionService {
	return &transactionService{
		tm: tm,
	}
}

func (ts *transactionService) CreateTransaction(customer_id uint, req entities.TransactionRequest) (entities.TransactionResponse, error) {
	transaction := entities.Transaction{
		BookingID:         fmt.Sprintf("LK-%d%d%d%d-%d", customer_id, req.HouseID, req.RoomID, req.Duration, req.CheckIn),
		UserID:            customer_id,
		RoomID:            req.RoomID,
		HouseID:           req.HouseID,
		CheckIn:           time.Unix(0, req.CheckIn*int64(time.Millisecond)),
		Duration:          req.Duration,
		Price:             req.Price,
		TransactionStatus: "processing",
	}

	result, err := ts.tm.Create(transaction)
	if err != nil {
		return entities.TransactionResponse{}, err
	}

	response, err := ts.tm.Request(result.BookingID)
	if err != nil {
		return entities.TransactionResponse{}, err
	}

	copier.Copy(&response, &result)
	return response, nil
}

func (ts *transactionService) GetAllTransactionbyConsultant() []entities.TransactionResponse {
	var response []entities.TransactionResponse

	results := ts.tm.GetAllbyConsultant()

	for _, r := range results {
		transaction, _ := ts.tm.Request(r.BookingID)
		copier.Copy(&transaction, &r)
		response = append(response, transaction)
	}

	return response
}

func (ts *transactionService) UpdateTransaction(customer_id uint, booking_id string, request entities.TransactionUpdateRequest) (entities.TransactionUpdateResponse, error) {
	req, err := ts.tm.Request(booking_id)
	if err != nil {
		return entities.TransactionUpdateResponse{}, err
	}

	snapRequest := &snap.Request{
		CustomerDetail: &midtrans.CustomerDetails{
			FName: req.Name,
			Email: req.Email,
			Phone: req.Phone,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  booking_id,
			GrossAmt: request.Price,
		},
		Items: &[]midtrans.ItemDetails{
			{
				Name:  req.Title,
				Price: request.Price,
				Qty:   1,
			},
		},
		EnabledPayments: snap.AllSnapPaymentType,
		Gopay: &snap.GopayDetails{
			EnableCallback: true,
		},
		Callbacks: &snap.Callbacks{
			Finish: configs.Get().App.FrontURL,
		},
	}

	snap, err := ts.tm.CreateSnap(snapRequest)
	if err != nil {
		return entities.TransactionUpdateResponse{}, err
	}

	transaction := entities.Transaction{
		BookingID:         booking_id,
		ConsultantID:      customer_id,
		Duration:          req.Duration,
		Price:             request.Price,
		TransactionStatus: "pending",
		RedirectURL:       snap.RedirectURL,
	}

	result, err := ts.tm.Update(booking_id, transaction)
	if err != nil {
		return entities.TransactionUpdateResponse{}, err
	}

	var response entities.TransactionUpdateResponse
	copier.Copy(&response, &req)
	copier.Copy(&response, &snap)
	copier.Copy(&response, &result)
	return response, nil
}

func (ts *transactionService) UpdateCallback(req entities.Callback) (entities.Callback, error) {
	transaction := entities.Callback{
		TransactionStatus: req.TransactionStatus,
		TransactionID:     req.TransactionID,
		StatusCode:        req.StatusCode,
		SignatureKey:      req.SignatureKey,
		PaymentType:       req.PaymentType,
		OrderID:           req.OrderID,
		GrossAmount:       req.GrossAmount,
		FraudStatus:       req.FraudStatus,
		ApprovalCode:      req.ApprovalCode,
	}

	response, err := ts.tm.UpdateSnap(req.OrderID, transaction)
	if err != nil {
		return entities.Callback{}, err
	}

	return response, nil
}

func (ts *transactionService) GetAllTransactionbyUser(role string, user uint, status string, city uint, district uint) []entities.TransactionJoin {
	response := ts.tm.GetAllbyUser(role, user, status, city, district)
	return response
}

func (ts *transactionService) GetAllTransactionbyKost(duration int, status string, name string) []entities.TransactionKost {
	response := ts.tm.GetAllbyKost(duration, status, name)
	return response
}
