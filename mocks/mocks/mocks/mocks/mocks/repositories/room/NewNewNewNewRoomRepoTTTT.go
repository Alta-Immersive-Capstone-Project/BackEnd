// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// NewNewNewNewRoomRepoTTTT is an autogenerated mock type for the NewNewNewNewRoomRepoTTTT type
type NewNewNewNewRoomRepoTTTT struct {
	mock.Mock
}

// Cleanup provides a mock function with given fields: _a0
func (_m *NewNewNewNewRoomRepoTTTT) Cleanup(_a0 func()) {
	_m.Called(_a0)
}

// Errorf provides a mock function with given fields: format, args
func (_m *NewNewNewNewRoomRepoTTTT) Errorf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// FailNow provides a mock function with given fields:
func (_m *NewNewNewNewRoomRepoTTTT) FailNow() {
	_m.Called()
}

// Logf provides a mock function with given fields: format, args
func (_m *NewNewNewNewRoomRepoTTTT) Logf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

type NewNewNewNewNewRoomRepoTTTTT interface {
	mock.TestingT
	Cleanup(func())
}

// NewNewNewNewNewRoomRepoTTTT creates a new instance of NewNewNewNewRoomRepoTTTT. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNewNewNewNewRoomRepoTTTT(t NewNewNewNewNewRoomRepoTTTTT) *NewNewNewNewRoomRepoTTTT {
	mock := &NewNewNewNewRoomRepoTTTT{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
