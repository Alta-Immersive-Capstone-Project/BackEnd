// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// NewNewNewNewNewNewFacilityHandlerTTTTTT is an autogenerated mock type for the NewNewNewNewNewNewFacilityHandlerTTTTTT type
type NewNewNewNewNewNewFacilityHandlerTTTTTT struct {
	mock.Mock
}

// Cleanup provides a mock function with given fields: _a0
func (_m *NewNewNewNewNewNewFacilityHandlerTTTTTT) Cleanup(_a0 func()) {
	_m.Called(_a0)
}

// Errorf provides a mock function with given fields: format, args
func (_m *NewNewNewNewNewNewFacilityHandlerTTTTTT) Errorf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// FailNow provides a mock function with given fields:
func (_m *NewNewNewNewNewNewFacilityHandlerTTTTTT) FailNow() {
	_m.Called()
}

// Logf provides a mock function with given fields: format, args
func (_m *NewNewNewNewNewNewFacilityHandlerTTTTTT) Logf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

type NewNewNewNewNewNewNewFacilityHandlerTTTTTTT interface {
	mock.TestingT
	Cleanup(func())
}

// NewNewNewNewNewNewNewFacilityHandlerTTTTTT creates a new instance of NewNewNewNewNewNewFacilityHandlerTTTTTT. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNewNewNewNewNewNewFacilityHandlerTTTTTT(t NewNewNewNewNewNewNewFacilityHandlerTTTTTTT) *NewNewNewNewNewNewFacilityHandlerTTTTTT {
	mock := &NewNewNewNewNewNewFacilityHandlerTTTTTT{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
