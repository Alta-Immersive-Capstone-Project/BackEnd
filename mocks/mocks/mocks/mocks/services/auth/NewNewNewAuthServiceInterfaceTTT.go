// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// NewNewNewAuthServiceInterfaceTTT is an autogenerated mock type for the NewNewNewAuthServiceInterfaceTTT type
type NewNewNewAuthServiceInterfaceTTT struct {
	mock.Mock
}

// Cleanup provides a mock function with given fields: _a0
func (_m *NewNewNewAuthServiceInterfaceTTT) Cleanup(_a0 func()) {
	_m.Called(_a0)
}

// Errorf provides a mock function with given fields: format, args
func (_m *NewNewNewAuthServiceInterfaceTTT) Errorf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// FailNow provides a mock function with given fields:
func (_m *NewNewNewAuthServiceInterfaceTTT) FailNow() {
	_m.Called()
}

// Logf provides a mock function with given fields: format, args
func (_m *NewNewNewAuthServiceInterfaceTTT) Logf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

type NewNewNewNewAuthServiceInterfaceTTTT interface {
	mock.TestingT
	Cleanup(func())
}

// NewNewNewNewAuthServiceInterfaceTTT creates a new instance of NewNewNewAuthServiceInterfaceTTT. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNewNewNewAuthServiceInterfaceTTT(t NewNewNewNewAuthServiceInterfaceTTTT) *NewNewNewAuthServiceInterfaceTTT {
	mock := &NewNewNewAuthServiceInterfaceTTT{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}