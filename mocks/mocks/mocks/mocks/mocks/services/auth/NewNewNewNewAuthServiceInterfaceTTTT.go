// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// NewNewNewNewAuthServiceInterfaceTTTT is an autogenerated mock type for the NewNewNewNewAuthServiceInterfaceTTTT type
type NewNewNewNewAuthServiceInterfaceTTTT struct {
	mock.Mock
}

// Cleanup provides a mock function with given fields: _a0
func (_m *NewNewNewNewAuthServiceInterfaceTTTT) Cleanup(_a0 func()) {
	_m.Called(_a0)
}

// Errorf provides a mock function with given fields: format, args
func (_m *NewNewNewNewAuthServiceInterfaceTTTT) Errorf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// FailNow provides a mock function with given fields:
func (_m *NewNewNewNewAuthServiceInterfaceTTTT) FailNow() {
	_m.Called()
}

// Logf provides a mock function with given fields: format, args
func (_m *NewNewNewNewAuthServiceInterfaceTTTT) Logf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

type NewNewNewNewNewAuthServiceInterfaceTTTTT interface {
	mock.TestingT
	Cleanup(func())
}

// NewNewNewNewNewAuthServiceInterfaceTTTT creates a new instance of NewNewNewNewAuthServiceInterfaceTTTT. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNewNewNewNewAuthServiceInterfaceTTTT(t NewNewNewNewNewAuthServiceInterfaceTTTTT) *NewNewNewNewAuthServiceInterfaceTTTT {
	mock := &NewNewNewNewAuthServiceInterfaceTTTT{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
