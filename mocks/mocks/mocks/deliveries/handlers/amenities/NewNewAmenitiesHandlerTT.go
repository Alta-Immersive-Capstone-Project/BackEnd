// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// NewNewAmenitiesHandlerTT is an autogenerated mock type for the NewNewAmenitiesHandlerTT type
type NewNewAmenitiesHandlerTT struct {
	mock.Mock
}

// Cleanup provides a mock function with given fields: _a0
func (_m *NewNewAmenitiesHandlerTT) Cleanup(_a0 func()) {
	_m.Called(_a0)
}

// Errorf provides a mock function with given fields: format, args
func (_m *NewNewAmenitiesHandlerTT) Errorf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// FailNow provides a mock function with given fields:
func (_m *NewNewAmenitiesHandlerTT) FailNow() {
	_m.Called()
}

// Logf provides a mock function with given fields: format, args
func (_m *NewNewAmenitiesHandlerTT) Logf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

type NewNewNewAmenitiesHandlerTTT interface {
	mock.TestingT
	Cleanup(func())
}

// NewNewNewAmenitiesHandlerTT creates a new instance of NewNewAmenitiesHandlerTT. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNewNewAmenitiesHandlerTT(t NewNewNewAmenitiesHandlerTTT) *NewNewAmenitiesHandlerTT {
	mock := &NewNewAmenitiesHandlerTT{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
