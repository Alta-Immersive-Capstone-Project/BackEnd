// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// NewNewNewNewNewNewIDistrictHandlerTTTTTT is an autogenerated mock type for the NewNewNewNewNewNewIDistrictHandlerTTTTTT type
type NewNewNewNewNewNewIDistrictHandlerTTTTTT struct {
	mock.Mock
}

// Cleanup provides a mock function with given fields: _a0
func (_m *NewNewNewNewNewNewIDistrictHandlerTTTTTT) Cleanup(_a0 func()) {
	_m.Called(_a0)
}

// Errorf provides a mock function with given fields: format, args
func (_m *NewNewNewNewNewNewIDistrictHandlerTTTTTT) Errorf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// FailNow provides a mock function with given fields:
func (_m *NewNewNewNewNewNewIDistrictHandlerTTTTTT) FailNow() {
	_m.Called()
}

// Logf provides a mock function with given fields: format, args
func (_m *NewNewNewNewNewNewIDistrictHandlerTTTTTT) Logf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

type NewNewNewNewNewNewNewIDistrictHandlerTTTTTTT interface {
	mock.TestingT
	Cleanup(func())
}

// NewNewNewNewNewNewNewIDistrictHandlerTTTTTT creates a new instance of NewNewNewNewNewNewIDistrictHandlerTTTTTT. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNewNewNewNewNewNewIDistrictHandlerTTTTTT(t NewNewNewNewNewNewNewIDistrictHandlerTTTTTTT) *NewNewNewNewNewNewIDistrictHandlerTTTTTT {
	mock := &NewNewNewNewNewNewIDistrictHandlerTTTTTT{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
