// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// NewCityHandlerT is an autogenerated mock type for the NewCityHandlerT type
type NewCityHandlerT struct {
	mock.Mock
}

// Cleanup provides a mock function with given fields: _a0
func (_m *NewCityHandlerT) Cleanup(_a0 func()) {
	_m.Called(_a0)
}

// Errorf provides a mock function with given fields: format, args
func (_m *NewCityHandlerT) Errorf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// FailNow provides a mock function with given fields:
func (_m *NewCityHandlerT) FailNow() {
	_m.Called()
}

// Logf provides a mock function with given fields: format, args
func (_m *NewCityHandlerT) Logf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

type NewNewCityHandlerTT interface {
	mock.TestingT
	Cleanup(func())
}

// NewNewCityHandlerT creates a new instance of NewCityHandlerT. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNewCityHandlerT(t NewNewCityHandlerTT) *NewCityHandlerT {
	mock := &NewCityHandlerT{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}