// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// NewNewNewNewNewAuthServiceInterfaceTTTTT is an autogenerated mock type for the NewNewNewNewNewAuthServiceInterfaceTTTTT type
type NewNewNewNewNewAuthServiceInterfaceTTTTT struct {
	mock.Mock
}

// Cleanup provides a mock function with given fields: _a0
func (_m *NewNewNewNewNewAuthServiceInterfaceTTTTT) Cleanup(_a0 func()) {
	_m.Called(_a0)
}

// Errorf provides a mock function with given fields: format, args
func (_m *NewNewNewNewNewAuthServiceInterfaceTTTTT) Errorf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// FailNow provides a mock function with given fields:
func (_m *NewNewNewNewNewAuthServiceInterfaceTTTTT) FailNow() {
	_m.Called()
}

// Logf provides a mock function with given fields: format, args
func (_m *NewNewNewNewNewAuthServiceInterfaceTTTTT) Logf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

type NewNewNewNewNewNewAuthServiceInterfaceTTTTTT interface {
	mock.TestingT
	Cleanup(func())
}

// NewNewNewNewNewNewAuthServiceInterfaceTTTTT creates a new instance of NewNewNewNewNewAuthServiceInterfaceTTTTT. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNewNewNewNewNewAuthServiceInterfaceTTTTT(t NewNewNewNewNewNewAuthServiceInterfaceTTTTTT) *NewNewNewNewNewAuthServiceInterfaceTTTTT {
	mock := &NewNewNewNewNewAuthServiceInterfaceTTTTT{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
