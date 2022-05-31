package transactions

import (
	"fmt"
	"kost/entities"
	repo "kost/repositories/transactions"
	"time"

	"github.com/jinzhu/copier"
)

type transactionService struct {
	tm repo.TransactionModel
}

func NewTransactionService(tm repo.TransactionModel) *transactionService {
	return &transactionService{
		tm: tm,
	}
}

func (ts *transactionService) AddTransaction(customer_id uint, request entities.TransactionRequest) (entities.TransactionResponse, error) {
	var response entities.TransactionResponse

	transaction := entities.Transaction{
		UserID:        customer_id,
		RoomID:        request.RoomID,
		CheckinDate:   time.Unix(0, request.CheckinDate*int64(time.Millisecond)),
		RentDuration:  request.RentDuration,
		BookingID:     fmt.Sprintf("DM-%d", request.CheckinDate),
		TotalBill:     request.TotalBill,
		PaymentMethod: request.PaymentMethod,
		Status:        "pending",
	}

	result, err := ts.tm.Create(transaction)
	if err != nil {
		return entities.TransactionResponse{}, err
	}

	copier.Copy(&response, &result)
	return response, nil
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

func (ts *transactionService) GetAllTransactionbyCustomer(customer_id uint, status string) []entities.TransactionResponse {
	var response []entities.TransactionResponse

	results := ts.tm.GetAllbyCustomer(customer_id, status)

	for _, r := range results {
		var transaction entities.TransactionResponse
		copier.Copy(&transaction, &r)
		response = append(response, transaction)
	}

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
