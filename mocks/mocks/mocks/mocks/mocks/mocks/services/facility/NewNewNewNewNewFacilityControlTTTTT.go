// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// NewNewNewNewNewFacilityControlTTTTT is an autogenerated mock type for the NewNewNewNewNewFacilityControlTTTTT type
type NewNewNewNewNewFacilityControlTTTTT struct {
	mock.Mock
}

// Cleanup provides a mock function with given fields: _a0
func (_m *NewNewNewNewNewFacilityControlTTTTT) Cleanup(_a0 func()) {
	_m.Called(_a0)
}

// Errorf provides a mock function with given fields: format, args
func (_m *NewNewNewNewNewFacilityControlTTTTT) Errorf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// FailNow provides a mock function with given fields:
func (_m *NewNewNewNewNewFacilityControlTTTTT) FailNow() {
	_m.Called()
}

// Logf provides a mock function with given fields: format, args
func (_m *NewNewNewNewNewFacilityControlTTTTT) Logf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

type NewNewNewNewNewNewFacilityControlTTTTTT interface {
	mock.TestingT
	Cleanup(func())
}

// NewNewNewNewNewNewFacilityControlTTTTT creates a new instance of NewNewNewNewNewFacilityControlTTTTT. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNewNewNewNewNewFacilityControlTTTTT(t NewNewNewNewNewNewFacilityControlTTTTTT) *NewNewNewNewNewFacilityControlTTTTT {
	mock := &NewNewNewNewNewFacilityControlTTTTT{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
