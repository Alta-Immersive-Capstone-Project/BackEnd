// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// NewReviewHandlerT is an autogenerated mock type for the NewReviewHandlerT type
type NewReviewHandlerT struct {
	mock.Mock
}

// Cleanup provides a mock function with given fields: _a0
func (_m *NewReviewHandlerT) Cleanup(_a0 func()) {
	_m.Called(_a0)
}

// Errorf provides a mock function with given fields: format, args
func (_m *NewReviewHandlerT) Errorf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// FailNow provides a mock function with given fields:
func (_m *NewReviewHandlerT) FailNow() {
	_m.Called()
}

// Logf provides a mock function with given fields: format, args
func (_m *NewReviewHandlerT) Logf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

type NewNewReviewHandlerTT interface {
	mock.TestingT
	Cleanup(func())
}

// NewNewReviewHandlerT creates a new instance of NewReviewHandlerT. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNewReviewHandlerT(t NewNewReviewHandlerTT) *NewReviewHandlerT {
	mock := &NewReviewHandlerT{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
