// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// NewRepoFacilityT is an autogenerated mock type for the NewRepoFacilityT type
type NewRepoFacilityT struct {
	mock.Mock
}

// Cleanup provides a mock function with given fields: _a0
func (_m *NewRepoFacilityT) Cleanup(_a0 func()) {
	_m.Called(_a0)
}

// Errorf provides a mock function with given fields: format, args
func (_m *NewRepoFacilityT) Errorf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// FailNow provides a mock function with given fields:
func (_m *NewRepoFacilityT) FailNow() {
	_m.Called()
}

// Logf provides a mock function with given fields: format, args
func (_m *NewRepoFacilityT) Logf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

type NewNewRepoFacilityTT interface {
	mock.TestingT
	Cleanup(func())
}

// NewNewRepoFacilityT creates a new instance of NewRepoFacilityT. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNewRepoFacilityT(t NewNewRepoFacilityTT) *NewRepoFacilityT {
	mock := &NewRepoFacilityT{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
