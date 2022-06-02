// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	entities "kost/entities"

	snap "github.com/midtrans/midtrans-go/snap"
	mock "github.com/stretchr/testify/mock"
)

// TransactionModel is an autogenerated mock type for the TransactionModel type
type TransactionModel struct {
	mock.Mock
}

// Create provides a mock function with given fields: transaction
func (_m *TransactionModel) Create(transaction entities.Transaction) (entities.Transaction, error) {
	ret := _m.Called(transaction)

	var r0 entities.Transaction
	if rf, ok := ret.Get(0).(func(entities.Transaction) entities.Transaction); ok {
		r0 = rf(transaction)
	} else {
		r0 = ret.Get(0).(entities.Transaction)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entities.Transaction) error); ok {
		r1 = rf(transaction)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSnap provides a mock function with given fields: req
func (_m *TransactionModel) CreateSnap(req *snap.Request) (*snap.Response, error) {
	ret := _m.Called(req)

	var r0 *snap.Response
	if rf, ok := ret.Get(0).(func(*snap.Request) *snap.Response); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snap.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*snap.Request) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllbyConsultant provides a mock function with given fields:
func (_m *TransactionModel) GetAllbyConsultant() []entities.Transaction {
	ret := _m.Called()

	var r0 []entities.Transaction
	if rf, ok := ret.Get(0).(func() []entities.Transaction); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.Transaction)
		}
	}

	return r0
}

// GetAllbyCustomer provides a mock function with given fields: role, user, status, city, district
func (_m *TransactionModel) GetAllbyCustomer(role string, user uint, status string, city uint, district uint) []entities.TransactionJoin {
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

// GetAllbyKost provides a mock function with given fields: duration, status, name
func (_m *TransactionModel) GetAllbyKost(duration int, status string, name string) []entities.TransactionKost {
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

// Request provides a mock function with given fields: booking_id
func (_m *TransactionModel) Request(booking_id string) (entities.TransactionResponse, error) {
	ret := _m.Called(booking_id)

	var r0 entities.TransactionResponse
	if rf, ok := ret.Get(0).(func(string) entities.TransactionResponse); ok {
		r0 = rf(booking_id)
	} else {
		r0 = ret.Get(0).(entities.TransactionResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(booking_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: booking_id, transaction
func (_m *TransactionModel) Update(booking_id string, transaction entities.Transaction) (entities.Transaction, error) {
	ret := _m.Called(booking_id, transaction)

	var r0 entities.Transaction
	if rf, ok := ret.Get(0).(func(string, entities.Transaction) entities.Transaction); ok {
		r0 = rf(booking_id, transaction)
	} else {
		r0 = ret.Get(0).(entities.Transaction)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, entities.Transaction) error); ok {
		r1 = rf(booking_id, transaction)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateSnap provides a mock function with given fields: booking_id, status
func (_m *TransactionModel) UpdateSnap(booking_id string, status entities.Callback) (entities.Callback, error) {
	ret := _m.Called(booking_id, status)

	var r0 entities.Callback
	if rf, ok := ret.Get(0).(func(string, entities.Callback) entities.Callback); ok {
		r0 = rf(booking_id, status)
	} else {
		r0 = ret.Get(0).(entities.Callback)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, entities.Callback) error); ok {
		r1 = rf(booking_id, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewTransactionModelT interface {
	mock.TestingT
	Cleanup(func())
}

// NewTransactionModel creates a new instance of TransactionModel. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTransactionModel(t NewTransactionModelT) *TransactionModel {
	mock := &TransactionModel{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
