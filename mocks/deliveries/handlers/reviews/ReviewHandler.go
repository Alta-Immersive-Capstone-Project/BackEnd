// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// ReviewHandler is an autogenerated mock type for the ReviewHandler type
type ReviewHandler struct {
	mock.Mock
}

// GetByHouseID provides a mock function with given fields: c
func (_m *ReviewHandler) GetByHouseID(c echo.Context) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertComment provides a mock function with given fields: c
func (_m *ReviewHandler) InsertComment(c echo.Context) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type NewReviewHandlerT interface {
	mock.TestingT
	Cleanup(func())
}

// NewReviewHandler creates a new instance of ReviewHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewReviewHandler(t NewReviewHandlerT) *ReviewHandler {
	mock := &ReviewHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
