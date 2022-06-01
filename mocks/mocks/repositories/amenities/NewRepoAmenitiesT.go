// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// NewRepoAmenitiesT is an autogenerated mock type for the NewRepoAmenitiesT type
type NewRepoAmenitiesT struct {
	mock.Mock
}

// Cleanup provides a mock function with given fields: _a0
func (_m *NewRepoAmenitiesT) Cleanup(_a0 func()) {
	_m.Called(_a0)
}

// Errorf provides a mock function with given fields: format, args
func (_m *NewRepoAmenitiesT) Errorf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// FailNow provides a mock function with given fields:
func (_m *NewRepoAmenitiesT) FailNow() {
	_m.Called()
}

// Logf provides a mock function with given fields: format, args
func (_m *NewRepoAmenitiesT) Logf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

type NewNewRepoAmenitiesTT interface {
	mock.TestingT
	Cleanup(func())
}

// NewNewRepoAmenitiesT creates a new instance of NewRepoAmenitiesT. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNewRepoAmenitiesT(t NewNewRepoAmenitiesTT) *NewRepoAmenitiesT {
	mock := &NewRepoAmenitiesT{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
