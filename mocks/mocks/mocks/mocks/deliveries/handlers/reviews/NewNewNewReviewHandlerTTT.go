// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// NewNewNewReviewHandlerTTT is an autogenerated mock type for the NewNewNewReviewHandlerTTT type
type NewNewNewReviewHandlerTTT struct {
	mock.Mock
}

// Cleanup provides a mock function with given fields: _a0
func (_m *NewNewNewReviewHandlerTTT) Cleanup(_a0 func()) {
	_m.Called(_a0)
}

// Errorf provides a mock function with given fields: format, args
func (_m *NewNewNewReviewHandlerTTT) Errorf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// FailNow provides a mock function with given fields:
func (_m *NewNewNewReviewHandlerTTT) FailNow() {
	_m.Called()
}

// Logf provides a mock function with given fields: format, args
func (_m *NewNewNewReviewHandlerTTT) Logf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

type NewNewNewNewReviewHandlerTTTT interface {
	mock.TestingT
	Cleanup(func())
}

// NewNewNewNewReviewHandlerTTT creates a new instance of NewNewNewReviewHandlerTTT. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNewNewNewReviewHandlerTTT(t NewNewNewNewReviewHandlerTTTT) *NewNewNewReviewHandlerTTT {
	mock := &NewNewNewReviewHandlerTTT{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}