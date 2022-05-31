// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// NewNewIHouseServiceTT is an autogenerated mock type for the NewNewIHouseServiceTT type
type NewNewIHouseServiceTT struct {
	mock.Mock
}

// Cleanup provides a mock function with given fields: _a0
func (_m *NewNewIHouseServiceTT) Cleanup(_a0 func()) {
	_m.Called(_a0)
}

// Errorf provides a mock function with given fields: format, args
func (_m *NewNewIHouseServiceTT) Errorf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// FailNow provides a mock function with given fields:
func (_m *NewNewIHouseServiceTT) FailNow() {
	_m.Called()
}

// Logf provides a mock function with given fields: format, args
func (_m *NewNewIHouseServiceTT) Logf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

type NewNewNewIHouseServiceTTT interface {
	mock.TestingT
	Cleanup(func())
}

// NewNewNewIHouseServiceTT creates a new instance of NewNewIHouseServiceTT. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNewNewIHouseServiceTT(t NewNewNewIHouseServiceTTT) *NewNewIHouseServiceTT {
	mock := &NewNewIHouseServiceTT{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}