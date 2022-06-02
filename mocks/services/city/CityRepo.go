// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	entities "kost/entities"

	mock "github.com/stretchr/testify/mock"
)

// CityRepo is an autogenerated mock type for the CityRepo type
type CityRepo struct {
	mock.Mock
}

// CreateCity provides a mock function with given fields: Insert
func (_m *CityRepo) CreateCity(Insert entities.AddCity) (entities.CityResponse, error) {
	ret := _m.Called(Insert)

	var r0 entities.CityResponse
	if rf, ok := ret.Get(0).(func(entities.AddCity) entities.CityResponse); ok {
		r0 = rf(Insert)
	} else {
		r0 = ret.Get(0).(entities.CityResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entities.AddCity) error); ok {
		r1 = rf(Insert)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCity provides a mock function with given fields: id
func (_m *CityRepo) DeleteCity(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllCity provides a mock function with given fields:
func (_m *CityRepo) GetAllCity() ([]entities.CityResponse, error) {
	ret := _m.Called()

	var r0 []entities.CityResponse
	if rf, ok := ret.Get(0).(func() []entities.CityResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.CityResponse)
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

// GetIDCity provides a mock function with given fields: id
func (_m *CityRepo) GetIDCity(id uint) (entities.CityResponse, error) {
	ret := _m.Called(id)

	var r0 entities.CityResponse
	if rf, ok := ret.Get(0).(func(uint) entities.CityResponse); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(entities.CityResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCity provides a mock function with given fields: id, update
func (_m *CityRepo) UpdateCity(id uint, update entities.AddCity) (entities.CityResponse, error) {
	ret := _m.Called(id, update)

	var r0 entities.CityResponse
	if rf, ok := ret.Get(0).(func(uint, entities.AddCity) entities.CityResponse); ok {
		r0 = rf(id, update)
	} else {
		r0 = ret.Get(0).(entities.CityResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, entities.AddCity) error); ok {
		r1 = rf(id, update)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewCityRepoT interface {
	mock.TestingT
	Cleanup(func())
}

// NewCityRepo creates a new instance of CityRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCityRepo(t NewCityRepoT) *CityRepo {
	mock := &CityRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
