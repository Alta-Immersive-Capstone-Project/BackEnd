// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// NewNewNewNewReviewHandlerTTTT is an autogenerated mock type for the NewNewNewNewReviewHandlerTTTT type
type NewNewNewNewReviewHandlerTTTT struct {
	mock.Mock
}

// Cleanup provides a mock function with given fields: _a0
func (_m *NewNewNewNewReviewHandlerTTTT) Cleanup(_a0 func()) {
	_m.Called(_a0)
}

// Errorf provides a mock function with given fields: format, args
func (_m *NewNewNewNewReviewHandlerTTTT) Errorf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// FailNow provides a mock function with given fields:
func (_m *NewNewNewNewReviewHandlerTTTT) FailNow() {
	_m.Called()
}

// Logf provides a mock function with given fields: format, args
func (_m *NewNewNewNewReviewHandlerTTTT) Logf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

type NewNewNewNewNewReviewHandlerTTTTT interface {
	mock.TestingT
	Cleanup(func())
}

// NewNewNewNewNewReviewHandlerTTTT creates a new instance of NewNewNewNewReviewHandlerTTTT. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNewNewNewNewReviewHandlerTTTT(t NewNewNewNewNewReviewHandlerTTTTT) *NewNewNewNewReviewHandlerTTTT {
	mock := &NewNewNewNewReviewHandlerTTTT{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
