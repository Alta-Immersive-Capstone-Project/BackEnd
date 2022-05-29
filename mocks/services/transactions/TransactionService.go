// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	entities "kost/entities"

	mock "github.com/stretchr/testify/mock"
)

// TransactionService is an autogenerated mock type for the TransactionService type
type TransactionService struct {
	mock.Mock
}

// AddTransaction provides a mock function with given fields: customer_id, request
func (_m *TransactionService) AddTransaction(customer_id uint, request *entities.TransactionRequest) (*entities.TransactionResponse, error) {
	ret := _m.Called(customer_id, request)

	var r0 *entities.TransactionResponse
	if rf, ok := ret.Get(0).(func(uint, *entities.TransactionRequest) *entities.TransactionResponse); ok {
		r0 = rf(customer_id, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.TransactionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, *entities.TransactionRequest) error); ok {
		r1 = rf(customer_id, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllTransactionbyConsultant provides a mock function with given fields:
func (_m *TransactionService) GetAllTransactionbyConsultant() []entities.TransactionResponse {
	ret := _m.Called()

	var r0 []entities.TransactionResponse
	if rf, ok := ret.Get(0).(func() []entities.TransactionResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.TransactionResponse)
		}
	}

	return r0
}

// GetAllTransactionbyCustomer provides a mock function with given fields: customer_id, status
func (_m *TransactionService) GetAllTransactionbyCustomer(customer_id uint, status string) []entities.TransactionResponse {
	ret := _m.Called(customer_id, status)

	var r0 []entities.TransactionResponse
	if rf, ok := ret.Get(0).(func(uint, string) []entities.TransactionResponse); ok {
		r0 = rf(customer_id, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.TransactionResponse)
		}
	}

	return r0
}

// GetTransaction provides a mock function with given fields: booking_id
func (_m *TransactionService) GetTransaction(booking_id string) (*entities.TransactionResponse, error) {
	ret := _m.Called(booking_id)

	var r0 *entities.TransactionResponse
	if rf, ok := ret.Get(0).(func(string) *entities.TransactionResponse); ok {
		r0 = rf(booking_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.TransactionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(booking_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTransaction provides a mock function with given fields: customer_id, booking_id, request
func (_m *TransactionService) UpdateTransaction(customer_id uint, booking_id string, request *entities.TransactionUpdateRequest) (*entities.TransactionUpdateResponse, error) {
	ret := _m.Called(customer_id, booking_id, request)

	var r0 *entities.TransactionUpdateResponse
	if rf, ok := ret.Get(0).(func(uint, string, *entities.TransactionUpdateRequest) *entities.TransactionUpdateResponse); ok {
		r0 = rf(customer_id, booking_id, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.TransactionUpdateResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, string, *entities.TransactionUpdateRequest) error); ok {
		r1 = rf(customer_id, booking_id, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewTransactionServiceT interface {
	mock.TestingT
	Cleanup(func())
}

// NewTransactionService creates a new instance of TransactionService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTransactionService(t NewTransactionServiceT) *TransactionService {
	mock := &TransactionService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}