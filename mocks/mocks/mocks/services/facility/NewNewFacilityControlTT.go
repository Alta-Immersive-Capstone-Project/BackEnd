// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// NewNewFacilityControlTT is an autogenerated mock type for the NewNewFacilityControlTT type
type NewNewFacilityControlTT struct {
	mock.Mock
}

// Cleanup provides a mock function with given fields: _a0
func (_m *NewNewFacilityControlTT) Cleanup(_a0 func()) {
	_m.Called(_a0)
}

// Errorf provides a mock function with given fields: format, args
func (_m *NewNewFacilityControlTT) Errorf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// FailNow provides a mock function with given fields:
func (_m *NewNewFacilityControlTT) FailNow() {
	_m.Called()
}

// Logf provides a mock function with given fields: format, args
func (_m *NewNewFacilityControlTT) Logf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

type NewNewNewFacilityControlTTT interface {
	mock.TestingT
	Cleanup(func())
}

// NewNewNewFacilityControlTT creates a new instance of NewNewFacilityControlTT. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNewNewFacilityControlTT(t NewNewNewFacilityControlTTT) *NewNewFacilityControlTT {
	mock := &NewNewFacilityControlTT{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
