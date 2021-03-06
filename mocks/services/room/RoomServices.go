// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	entities "kost/entities"

	mock "github.com/stretchr/testify/mock"
)

// RoomServices is an autogenerated mock type for the RoomServices type
type RoomServices struct {
	mock.Mock
}

// CreateRoom provides a mock function with given fields: id, Insert
func (_m *RoomServices) CreateRoom(id uint, Insert entities.AddRoom) (entities.RespondRoomcreat, error) {
	ret := _m.Called(id, Insert)

	var r0 entities.RespondRoomcreat
	if rf, ok := ret.Get(0).(func(uint, entities.AddRoom) entities.RespondRoomcreat); ok {
		r0 = rf(id, Insert)
	} else {
		r0 = ret.Get(0).(entities.RespondRoomcreat)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, entities.AddRoom) error); ok {
		r1 = rf(id, Insert)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteRoom provides a mock function with given fields: id
func (_m *RoomServices) DeleteRoom(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllRoom provides a mock function with given fields: id
func (_m *RoomServices) GetAllRoom(id uint) ([]entities.RespondRoom, error) {
	ret := _m.Called(id)

	var r0 []entities.RespondRoom
	if rf, ok := ret.Get(0).(func(uint) []entities.RespondRoom); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.RespondRoom)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetIDRoom provides a mock function with given fields: id
func (_m *RoomServices) GetIDRoom(id uint) (entities.Room, error) {
	ret := _m.Called(id)

	var r0 entities.Room
	if rf, ok := ret.Get(0).(func(uint) entities.Room); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(entities.Room)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateRoom provides a mock function with given fields: id, update
func (_m *RoomServices) UpdateRoom(id uint, update entities.UpdateRoom) (entities.RespondRoom, error) {
	ret := _m.Called(id, update)

	var r0 entities.RespondRoom
	if rf, ok := ret.Get(0).(func(uint, entities.UpdateRoom) entities.RespondRoom); ok {
		r0 = rf(id, update)
	} else {
		r0 = ret.Get(0).(entities.RespondRoom)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, entities.UpdateRoom) error); ok {
		r1 = rf(id, update)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewRoomServicesT interface {
	mock.TestingT
	Cleanup(func())
}

// NewRoomServices creates a new instance of RoomServices. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRoomServices(t NewRoomServicesT) *RoomServices {
	mock := &RoomServices{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
