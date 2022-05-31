// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// NewTransactionHandlerT is an autogenerated mock type for the NewTransactionHandlerT type
type NewTransactionHandlerT struct {
	mock.Mock
}

// Cleanup provides a mock function with given fields: _a0
func (_m *NewTransactionHandlerT) Cleanup(_a0 func()) {
	_m.Called(_a0)
}

// Errorf provides a mock function with given fields: format, args
func (_m *NewTransactionHandlerT) Errorf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// FailNow provides a mock function with given fields:
func (_m *NewTransactionHandlerT) FailNow() {
	_m.Called()
}

// Logf provides a mock function with given fields: format, args
func (_m *NewTransactionHandlerT) Logf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

type NewNewTransactionHandlerTT interface {
	mock.TestingT
	Cleanup(func())
}

// NewNewTransactionHandlerT creates a new instance of NewTransactionHandlerT. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNewTransactionHandlerT(t NewNewTransactionHandlerTT) *NewTransactionHandlerT {
	mock := &NewTransactionHandlerT{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}