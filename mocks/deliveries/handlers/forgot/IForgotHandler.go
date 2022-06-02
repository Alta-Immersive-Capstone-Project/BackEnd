// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// IForgotHandler is an autogenerated mock type for the IForgotHandler type
type IForgotHandler struct {
	mock.Mock
}

// ResetPassword provides a mock function with given fields:
func (_m *IForgotHandler) ResetPassword() echo.HandlerFunc {
	ret := _m.Called()

	var r0 echo.HandlerFunc
	if rf, ok := ret.Get(0).(func() echo.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.HandlerFunc)
		}
	}

	return r0
}

// SendEmail provides a mock function with given fields:
func (_m *IForgotHandler) SendEmail() echo.HandlerFunc {
	ret := _m.Called()

	var r0 echo.HandlerFunc
	if rf, ok := ret.Get(0).(func() echo.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.HandlerFunc)
		}
	}

	return r0
}

type NewIForgotHandlerT interface {
	mock.TestingT
	Cleanup(func())
}

// NewIForgotHandler creates a new instance of IForgotHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIForgotHandler(t NewIForgotHandlerT) *IForgotHandler {
	mock := &IForgotHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
