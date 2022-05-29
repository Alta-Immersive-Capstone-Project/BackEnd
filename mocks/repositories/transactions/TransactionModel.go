// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	entities "kost/entities"

	mock "github.com/stretchr/testify/mock"
)

// TransactionModel is an autogenerated mock type for the TransactionModel type
type TransactionModel struct {
	mock.Mock
}

// Create provides a mock function with given fields: transaction
func (_m *TransactionModel) Create(transaction *entities.Transaction) (*entities.Transaction, error) {
	ret := _m.Called(transaction)

	var r0 *entities.Transaction
	if rf, ok := ret.Get(0).(func(*entities.Transaction) *entities.Transaction); ok {
		r0 = rf(transaction)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entities.Transaction) error); ok {
		r1 = rf(transaction)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: booking_id
func (_m *TransactionModel) Get(booking_id string) (entities.Transaction, error) {
	ret := _m.Called(booking_id)

	var r0 entities.Transaction
	if rf, ok := ret.Get(0).(func(string) entities.Transaction); ok {
		r0 = rf(booking_id)
	} else {
		r0 = ret.Get(0).(entities.Transaction)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(booking_id)
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

// GetAllbyCustomer provides a mock function with given fields: customer_id, status
func (_m *TransactionModel) GetAllbyCustomer(customer_id uint, status string) []entities.Transaction {
	ret := _m.Called(customer_id, status)

	var r0 []entities.Transaction
	if rf, ok := ret.Get(0).(func(uint, string) []entities.Transaction); ok {
		r0 = rf(customer_id, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.Transaction)
		}
	}

	return r0
}

// Update provides a mock function with given fields: booking_id, transaction
func (_m *TransactionModel) Update(booking_id string, transaction *entities.Transaction) (*entities.Transaction, error) {
	ret := _m.Called(booking_id, transaction)

	var r0 *entities.Transaction
	if rf, ok := ret.Get(0).(func(string, *entities.Transaction) *entities.Transaction); ok {
		r0 = rf(booking_id, transaction)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *entities.Transaction) error); ok {
		r1 = rf(booking_id, transaction)
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