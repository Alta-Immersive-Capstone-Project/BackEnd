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

// CreateTransaction provides a mock function with given fields: customer_id, req
func (_m *TransactionService) CreateTransaction(customer_id uint, req entities.TransactionRequest) (entities.TransactionResponse, error) {
	ret := _m.Called(customer_id, req)

	var r0 entities.TransactionResponse
	if rf, ok := ret.Get(0).(func(uint, entities.TransactionRequest) entities.TransactionResponse); ok {
		r0 = rf(customer_id, req)
	} else {
		r0 = ret.Get(0).(entities.TransactionResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, entities.TransactionRequest) error); ok {
		r1 = rf(customer_id, req)
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

// GetAllTransactionbyKost provides a mock function with given fields: duration, status, name
func (_m *TransactionService) GetAllTransactionbyKost(duration int, status string, name string) []entities.TransactionKost {
	ret := _m.Called(duration, status, name)

	var r0 []entities.TransactionKost
	if rf, ok := ret.Get(0).(func(int, string, string) []entities.TransactionKost); ok {
		r0 = rf(duration, status, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.TransactionKost)
		}
	}

	return r0
}

// GetAllTransactionbyUser provides a mock function with given fields: role, user, status, city, district
func (_m *TransactionService) GetAllTransactionbyUser(role string, user uint, status string, city uint, district uint) []entities.TransactionJoin {
	ret := _m.Called(role, user, status, city, district)

	var r0 []entities.TransactionJoin
	if rf, ok := ret.Get(0).(func(string, uint, string, uint, uint) []entities.TransactionJoin); ok {
		r0 = rf(role, user, status, city, district)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.TransactionJoin)
		}
	}

	return r0
}

// GetReport provides a mock function with given fields: _a0
func (_m *TransactionService) GetReport(_a0 []entities.TransactionKost) string {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func([]entities.TransactionKost) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// UpdateCallback provides a mock function with given fields: req
func (_m *TransactionService) UpdateCallback(req entities.Callback) (entities.Callback, error) {
	ret := _m.Called(req)

	var r0 entities.Callback
	if rf, ok := ret.Get(0).(func(entities.Callback) entities.Callback); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Get(0).(entities.Callback)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entities.Callback) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTransaction provides a mock function with given fields: customer_id, booking_id, request
func (_m *TransactionService) UpdateTransaction(customer_id uint, booking_id string, request entities.TransactionUpdateRequest) (entities.TransactionUpdateResponse, error) {
	ret := _m.Called(customer_id, booking_id, request)

	var r0 entities.TransactionUpdateResponse
	if rf, ok := ret.Get(0).(func(uint, string, entities.TransactionUpdateRequest) entities.TransactionUpdateResponse); ok {
		r0 = rf(customer_id, booking_id, request)
	} else {
		r0 = ret.Get(0).(entities.TransactionUpdateResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, string, entities.TransactionUpdateRequest) error); ok {
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
