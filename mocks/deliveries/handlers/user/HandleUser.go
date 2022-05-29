// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// HandleUser is an autogenerated mock type for the HandleUser type
type HandleUser struct {
	mock.Mock
}

// CreateCustomer provides a mock function with given fields: c
func (_m *HandleUser) CreateCustomer(c echo.Context) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateInternal provides a mock function with given fields: c
func (_m *HandleUser) CreateInternal(c echo.Context) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCustomer provides a mock function with given fields: c
func (_m *HandleUser) DeleteCustomer(c echo.Context) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteInternal provides a mock function with given fields: c
func (_m *HandleUser) DeleteInternal(c echo.Context) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateCustomer provides a mock function with given fields: c
func (_m *HandleUser) UpdateCustomer(c echo.Context) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateInternal provides a mock function with given fields: c
func (_m *HandleUser) UpdateInternal(c echo.Context) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type NewHandleUserT interface {
	mock.TestingT
	Cleanup(func())
}

// NewHandleUser creates a new instance of HandleUser. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHandleUser(t NewHandleUserT) *HandleUser {
	mock := &HandleUser{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
