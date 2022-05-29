// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	entities "kost/entities"

	mock "github.com/stretchr/testify/mock"
)

// AmenitiesControl is an autogenerated mock type for the AmenitiesControl type
type AmenitiesControl struct {
	mock.Mock
}

// CreateAmenities provides a mock function with given fields: Insert
func (_m *AmenitiesControl) CreateAmenities(Insert entities.AddAmenities) (entities.RespondAmenities, error) {
	ret := _m.Called(Insert)

	var r0 entities.RespondAmenities
	if rf, ok := ret.Get(0).(func(entities.AddAmenities) entities.RespondAmenities); ok {
		r0 = rf(Insert)
	} else {
		r0 = ret.Get(0).(entities.RespondAmenities)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entities.AddAmenities) error); ok {
		r1 = rf(Insert)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAmenities provides a mock function with given fields: id
func (_m *AmenitiesControl) DeleteAmenities(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllAmenities provides a mock function with given fields: RoomID
func (_m *AmenitiesControl) GetAllAmenities(RoomID uint) ([]entities.RespondAmenities, error) {
	ret := _m.Called(RoomID)

	var r0 []entities.RespondAmenities
	if rf, ok := ret.Get(0).(func(uint) []entities.RespondAmenities); ok {
		r0 = rf(RoomID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.RespondAmenities)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(RoomID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAmenitiesID provides a mock function with given fields: id
func (_m *AmenitiesControl) GetAmenitiesID(id uint) (entities.RespondAmenities, error) {
	ret := _m.Called(id)

	var r0 entities.RespondAmenities
	if rf, ok := ret.Get(0).(func(uint) entities.RespondAmenities); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(entities.RespondAmenities)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAmenities provides a mock function with given fields: id, update
func (_m *AmenitiesControl) UpdateAmenities(id uint, update entities.UpdateAmenities) (entities.RespondAmenities, error) {
	ret := _m.Called(id, update)

	var r0 entities.RespondAmenities
	if rf, ok := ret.Get(0).(func(uint, entities.UpdateAmenities) entities.RespondAmenities); ok {
		r0 = rf(id, update)
	} else {
		r0 = ret.Get(0).(entities.RespondAmenities)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, entities.UpdateAmenities) error); ok {
		r1 = rf(id, update)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewAmenitiesControlT interface {
	mock.TestingT
	Cleanup(func())
}

// NewAmenitiesControl creates a new instance of AmenitiesControl. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAmenitiesControl(t NewAmenitiesControlT) *AmenitiesControl {
	mock := &AmenitiesControl{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
