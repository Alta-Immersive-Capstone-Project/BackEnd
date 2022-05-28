// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	entities "kost/entities"

	mock "github.com/stretchr/testify/mock"
)

// RepoFacility is an autogenerated mock type for the RepoFacility type
type RepoFacility struct {
	mock.Mock
}

// CreateFacility provides a mock function with given fields: New
func (_m *RepoFacility) CreateFacility(New entities.Facility) (entities.Facility, error) {
	ret := _m.Called(New)

	var r0 entities.Facility
	if rf, ok := ret.Get(0).(func(entities.Facility) entities.Facility); ok {
		r0 = rf(New)
	} else {
		r0 = ret.Get(0).(entities.Facility)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entities.Facility) error); ok {
		r1 = rf(New)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteFacility provides a mock function with given fields: id
func (_m *RepoFacility) DeleteFacility(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllFacility provides a mock function with given fields: HouseID
func (_m *RepoFacility) GetAllFacility(HouseID uint) ([]entities.Facility, error) {
	ret := _m.Called(HouseID)

	var r0 []entities.Facility
	if rf, ok := ret.Get(0).(func(uint) []entities.Facility); ok {
		r0 = rf(HouseID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.Facility)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(HouseID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFacilityID provides a mock function with given fields: id
func (_m *RepoFacility) GetFacilityID(id uint) (entities.Facility, error) {
	ret := _m.Called(id)

	var r0 entities.Facility
	if rf, ok := ret.Get(0).(func(uint) entities.Facility); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(entities.Facility)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateFacility provides a mock function with given fields: id, UpdateFacility
func (_m *RepoFacility) UpdateFacility(id uint, UpdateFacility entities.Facility) (entities.Facility, error) {
	ret := _m.Called(id, UpdateFacility)

	var r0 entities.Facility
	if rf, ok := ret.Get(0).(func(uint, entities.Facility) entities.Facility); ok {
		r0 = rf(id, UpdateFacility)
	} else {
		r0 = ret.Get(0).(entities.Facility)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, entities.Facility) error); ok {
		r1 = rf(id, UpdateFacility)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewRepoFacilityT interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepoFacility creates a new instance of RepoFacility. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepoFacility(t NewRepoFacilityT) *RepoFacility {
	mock := &RepoFacility{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}