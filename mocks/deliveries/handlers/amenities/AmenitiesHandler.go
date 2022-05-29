// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// AmenitiesHandler is an autogenerated mock type for the AmenitiesHandler type
type AmenitiesHandler struct {
	mock.Mock
}

// CreateAmenities provides a mock function with given fields:
func (_m *AmenitiesHandler) CreateAmenities() echo.HandlerFunc {
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

// DeleteAmenities provides a mock function with given fields:
func (_m *AmenitiesHandler) DeleteAmenities() echo.HandlerFunc {
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

// GetAllAmenities provides a mock function with given fields:
func (_m *AmenitiesHandler) GetAllAmenities() echo.HandlerFunc {
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

// GetAmenitiesID provides a mock function with given fields:
func (_m *AmenitiesHandler) GetAmenitiesID() echo.HandlerFunc {
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

// UpdateAmenities provides a mock function with given fields:
func (_m *AmenitiesHandler) UpdateAmenities() echo.HandlerFunc {
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

type NewAmenitiesHandlerT interface {
	mock.TestingT
	Cleanup(func())
}

// NewAmenitiesHandler creates a new instance of AmenitiesHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAmenitiesHandler(t NewAmenitiesHandlerT) *AmenitiesHandler {
	mock := &AmenitiesHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
