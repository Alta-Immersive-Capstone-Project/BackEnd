// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// NewNewIHouseHandlerTT is an autogenerated mock type for the NewNewIHouseHandlerTT type
type NewNewIHouseHandlerTT struct {
	mock.Mock
}

// Cleanup provides a mock function with given fields: _a0
func (_m *NewNewIHouseHandlerTT) Cleanup(_a0 func()) {
	_m.Called(_a0)
}

// Errorf provides a mock function with given fields: format, args
func (_m *NewNewIHouseHandlerTT) Errorf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// FailNow provides a mock function with given fields:
func (_m *NewNewIHouseHandlerTT) FailNow() {
	_m.Called()
}

// Logf provides a mock function with given fields: format, args
func (_m *NewNewIHouseHandlerTT) Logf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

type NewNewNewIHouseHandlerTTT interface {
	mock.TestingT
	Cleanup(func())
}

// NewNewNewIHouseHandlerTT creates a new instance of NewNewIHouseHandlerTT. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNewNewIHouseHandlerTT(t NewNewNewIHouseHandlerTTT) *NewNewIHouseHandlerTT {
	mock := &NewNewIHouseHandlerTT{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}