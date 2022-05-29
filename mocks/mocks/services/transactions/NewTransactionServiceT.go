// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// NewTransactionServiceT is an autogenerated mock type for the NewTransactionServiceT type
type NewTransactionServiceT struct {
	mock.Mock
}

// Cleanup provides a mock function with given fields: _a0
func (_m *NewTransactionServiceT) Cleanup(_a0 func()) {
	_m.Called(_a0)
}

// Errorf provides a mock function with given fields: format, args
func (_m *NewTransactionServiceT) Errorf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// FailNow provides a mock function with given fields:
func (_m *NewTransactionServiceT) FailNow() {
	_m.Called()
}

// Logf provides a mock function with given fields: format, args
func (_m *NewTransactionServiceT) Logf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

type NewNewTransactionServiceTT interface {
	mock.TestingT
	Cleanup(func())
}

// NewNewTransactionServiceT creates a new instance of NewTransactionServiceT. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNewTransactionServiceT(t NewNewTransactionServiceTT) *NewTransactionServiceT {
	mock := &NewTransactionServiceT{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}