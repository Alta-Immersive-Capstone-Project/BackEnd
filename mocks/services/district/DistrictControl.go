// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	entities "kost/entities"

	mock "github.com/stretchr/testify/mock"
)

// DistrictControl is an autogenerated mock type for the DistrictControl type
type DistrictControl struct {
	mock.Mock
}

// CreateDist provides a mock function with given fields: Insert
func (_m *DistrictControl) CreateDist(Insert entities.AddDistrict) (entities.RespondDistrict, error) {
	ret := _m.Called(Insert)

	var r0 entities.RespondDistrict
	if rf, ok := ret.Get(0).(func(entities.AddDistrict) entities.RespondDistrict); ok {
		r0 = rf(Insert)
	} else {
		r0 = ret.Get(0).(entities.RespondDistrict)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entities.AddDistrict) error); ok {
		r1 = rf(Insert)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteDist provides a mock function with given fields: id
func (_m *DistrictControl) DeleteDist(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllDist provides a mock function with given fields: cid
func (_m *DistrictControl) GetAllDist(cid uint) ([]entities.RespondDistrict, error) {
	ret := _m.Called(cid)

	var r0 []entities.RespondDistrict
	if rf, ok := ret.Get(0).(func(uint) []entities.RespondDistrict); ok {
		r0 = rf(cid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.RespondDistrict)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(cid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDistID provides a mock function with given fields: id
func (_m *DistrictControl) GetDistID(id uint) (entities.RespondDistrict, error) {
	ret := _m.Called(id)

	var r0 entities.RespondDistrict
	if rf, ok := ret.Get(0).(func(uint) entities.RespondDistrict); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(entities.RespondDistrict)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectAllDistrict provides a mock function with given fields:
func (_m *DistrictControl) SelectAllDistrict() ([]entities.RespondDistrict, error) {
	ret := _m.Called()

	var r0 []entities.RespondDistrict
	if rf, ok := ret.Get(0).(func() []entities.RespondDistrict); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.RespondDistrict)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDist provides a mock function with given fields: id, update
func (_m *DistrictControl) UpdateDist(id uint, update entities.UpdateDistrict) (entities.RespondDistrict, error) {
	ret := _m.Called(id, update)

	var r0 entities.RespondDistrict
	if rf, ok := ret.Get(0).(func(uint, entities.UpdateDistrict) entities.RespondDistrict); ok {
		r0 = rf(id, update)
	} else {
		r0 = ret.Get(0).(entities.RespondDistrict)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, entities.UpdateDistrict) error); ok {
		r1 = rf(id, update)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewDistrictControlT interface {
	mock.TestingT
	Cleanup(func())
}

// NewDistrictControl creates a new instance of DistrictControl. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDistrictControl(t NewDistrictControlT) *DistrictControl {
	mock := &DistrictControl{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
