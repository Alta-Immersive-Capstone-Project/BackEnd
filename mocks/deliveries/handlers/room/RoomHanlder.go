// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// RoomHanlder is an autogenerated mock type for the RoomHanlder type
type RoomHanlder struct {
	mock.Mock
}

// CreateRoom provides a mock function with given fields:
func (_m *RoomHanlder) CreateRoom() echo.HandlerFunc {
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

<<<<<<< HEAD
=======
// DeleteImageUpdate provides a mock function with given fields:
func (_m *RoomHanlder) DeleteImageUpdate() echo.HandlerFunc {
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

>>>>>>> 3d2f172cae4224571053c1b5658836fe1402c6a9
// DeleteRoom provides a mock function with given fields:
func (_m *RoomHanlder) DeleteRoom() echo.HandlerFunc {
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

// GetAllRoom provides a mock function with given fields:
func (_m *RoomHanlder) GetAllRoom() echo.HandlerFunc {
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

// GetIDRoom provides a mock function with given fields:
func (_m *RoomHanlder) GetIDRoom() echo.HandlerFunc {
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

// UpdateRoom provides a mock function with given fields:
func (_m *RoomHanlder) UpdateRoom() echo.HandlerFunc {
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

type NewRoomHanlderT interface {
	mock.TestingT
	Cleanup(func())
}

// NewRoomHanlder creates a new instance of RoomHanlder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRoomHanlder(t NewRoomHanlderT) *RoomHanlder {
	mock := &RoomHanlder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
