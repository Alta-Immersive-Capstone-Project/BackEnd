// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	entities "kost/entities"

	mock "github.com/stretchr/testify/mock"
)

// RepoAmenities is an autogenerated mock type for the RepoAmenities type
type RepoAmenities struct {
	mock.Mock
}

// CreateAmenities provides a mock function with given fields: New
func (_m *RepoAmenities) CreateAmenities(New entities.Amenities) (entities.Amenities, error) {
	ret := _m.Called(New)

	var r0 entities.Amenities
	if rf, ok := ret.Get(0).(func(entities.Amenities) entities.Amenities); ok {
		r0 = rf(New)
	} else {
		r0 = ret.Get(0).(entities.Amenities)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entities.Amenities) error); ok {
		r1 = rf(New)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAmenities provides a mock function with given fields: id
func (_m *RepoAmenities) DeleteAmenities(id uint) error {
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
func (_m *RepoAmenities) GetAllAmenities(RoomID uint) ([]entities.Amenities, error) {
	ret := _m.Called(RoomID)

	var r0 []entities.Amenities
	if rf, ok := ret.Get(0).(func(uint) []entities.Amenities); ok {
		r0 = rf(RoomID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.Amenities)
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
func (_m *RepoAmenities) GetAmenitiesID(id uint) (entities.Amenities, error) {
	ret := _m.Called(id)

	var r0 entities.Amenities
	if rf, ok := ret.Get(0).(func(uint) entities.Amenities); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(entities.Amenities)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAmenities provides a mock function with given fields: id, UpdateAmenities
func (_m *RepoAmenities) UpdateAmenities(id uint, UpdateAmenities entities.Amenities) (entities.Amenities, error) {
	ret := _m.Called(id, UpdateAmenities)

	var r0 entities.Amenities
	if rf, ok := ret.Get(0).(func(uint, entities.Amenities) entities.Amenities); ok {
		r0 = rf(id, UpdateAmenities)
	} else {
		r0 = ret.Get(0).(entities.Amenities)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, entities.Amenities) error); ok {
		r1 = rf(id, UpdateAmenities)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewRepoAmenitiesT interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepoAmenities creates a new instance of RepoAmenities. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepoAmenities(t NewRepoAmenitiesT) *RepoAmenities {
	mock := &RepoAmenities{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}