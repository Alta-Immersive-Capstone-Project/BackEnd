// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// NewNewNewNewRoomHanlderTTTT is an autogenerated mock type for the NewNewNewNewRoomHanlderTTTT type
type NewNewNewNewRoomHanlderTTTT struct {
	mock.Mock
}

// Cleanup provides a mock function with given fields: _a0
func (_m *NewNewNewNewRoomHanlderTTTT) Cleanup(_a0 func()) {
	_m.Called(_a0)
}

// Errorf provides a mock function with given fields: format, args
func (_m *NewNewNewNewRoomHanlderTTTT) Errorf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// FailNow provides a mock function with given fields:
func (_m *NewNewNewNewRoomHanlderTTTT) FailNow() {
	_m.Called()
}

// Logf provides a mock function with given fields: format, args
func (_m *NewNewNewNewRoomHanlderTTTT) Logf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

type NewNewNewNewNewRoomHanlderTTTTT interface {
	mock.TestingT
	Cleanup(func())
}

// NewNewNewNewNewRoomHanlderTTTT creates a new instance of NewNewNewNewRoomHanlderTTTT. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNewNewNewNewRoomHanlderTTTT(t NewNewNewNewNewRoomHanlderTTTTT) *NewNewNewNewRoomHanlderTTTT {
	mock := &NewNewNewNewRoomHanlderTTTT{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}