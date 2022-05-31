// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	entities "kost/entities"

	mock "github.com/stretchr/testify/mock"
)

// IHouseService is an autogenerated mock type for the IHouseService type
type IHouseService struct {
	mock.Mock
}

// CreateHouse provides a mock function with given fields: Insert
func (_m *IHouseService) CreateHouse(Insert entities.AddHouse) (entities.HouseResponse, error) {
	ret := _m.Called(Insert)

	var r0 entities.HouseResponse
	if rf, ok := ret.Get(0).(func(entities.AddHouse) entities.HouseResponse); ok {
		r0 = rf(Insert)
	} else {
		r0 = ret.Get(0).(entities.HouseResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entities.AddHouse) error); ok {
		r1 = rf(Insert)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteHouse provides a mock function with given fields: id
func (_m *IHouseService) DeleteHouse(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAllHouseByCities provides a mock function with given fields: cid
func (_m *IHouseService) FindAllHouseByCities(cid uint) ([]entities.HouseResponseJoin, error) {
	ret := _m.Called(cid)

	var r0 []entities.HouseResponseJoin
	if rf, ok := ret.Get(0).(func(uint) []entities.HouseResponseJoin); ok {
		r0 = rf(cid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.HouseResponseJoin)
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

// FindAllHouseByCtyAndDst provides a mock function with given fields: cid, dist_id
func (_m *IHouseService) FindAllHouseByCtyAndDst(cid uint, dist_id uint) ([]entities.HouseResponseJoin, error) {
	ret := _m.Called(cid, dist_id)

	var r0 []entities.HouseResponseJoin
	if rf, ok := ret.Get(0).(func(uint, uint) []entities.HouseResponseJoin); ok {
		r0 = rf(cid, dist_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.HouseResponseJoin)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, uint) error); ok {
		r1 = rf(cid, dist_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAllHouseByDistrict provides a mock function with given fields: dist_id
func (_m *IHouseService) FindAllHouseByDistrict(dist_id uint) ([]entities.HouseResponseJoin, error) {
	ret := _m.Called(dist_id)

	var r0 []entities.HouseResponseJoin
	if rf, ok := ret.Get(0).(func(uint) []entities.HouseResponseJoin); ok {
		r0 = rf(dist_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.HouseResponseJoin)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(dist_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

<<<<<<< HEAD
// FindHouseByTitle provides a mock function with given fields: title
func (_m *IHouseService) FindHouseByTitle(title string) ([]entities.HouseResponseJoin, error) {
	ret := _m.Called(title)

	var r0 []entities.HouseResponseJoin
	if rf, ok := ret.Get(0).(func(string) []entities.HouseResponseJoin); ok {
		r0 = rf(title)
=======
// FindHouseByLocation provides a mock function with given fields: lat, long
func (_m *IHouseService) FindHouseByLocation(lat float64, long float64) ([]entities.HouseResponseJoin, error) {
	ret := _m.Called(lat, long)

	var r0 []entities.HouseResponseJoin
	if rf, ok := ret.Get(0).(func(float64, float64) []entities.HouseResponseJoin); ok {
		r0 = rf(lat, long)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.HouseResponseJoin)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(float64, float64) error); ok {
		r1 = rf(lat, long)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindHouseByTitle provides a mock function with given fields: name
func (_m *IHouseService) FindHouseByTitle(name string) ([]entities.HouseResponseJoin, error) {
	ret := _m.Called(name)

	var r0 []entities.HouseResponseJoin
	if rf, ok := ret.Get(0).(func(string) []entities.HouseResponseJoin); ok {
		r0 = rf(name)
>>>>>>> f3427b123ba48ac06c4ad5b903738e404dfd285b
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.HouseResponseJoin)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
<<<<<<< HEAD
		r1 = rf(title)
=======
		r1 = rf(name)
>>>>>>> f3427b123ba48ac06c4ad5b903738e404dfd285b
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

<<<<<<< HEAD
// GetAllHouseByDistrict provides a mock function with given fields: dist_id
func (_m *IHouseService) GetAllHouseByDistrict(dist_id uint) ([]entities.HouseResponse, error) {
=======
// GetAllHouseByDist provides a mock function with given fields: dist_id
func (_m *IHouseService) GetAllHouseByDist(dist_id uint) ([]entities.HouseResponse, error) {
>>>>>>> f3427b123ba48ac06c4ad5b903738e404dfd285b
	ret := _m.Called(dist_id)

	var r0 []entities.HouseResponse
	if rf, ok := ret.Get(0).(func(uint) []entities.HouseResponse); ok {
		r0 = rf(dist_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.HouseResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(dist_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHouseID provides a mock function with given fields: id
func (_m *IHouseService) GetHouseID(id uint) (entities.HouseResponse, error) {
	ret := _m.Called(id)

	var r0 entities.HouseResponse
	if rf, ok := ret.Get(0).(func(uint) entities.HouseResponse); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(entities.HouseResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectAllHouse provides a mock function with given fields:
func (_m *IHouseService) SelectAllHouse() ([]entities.HouseResponseJoin, error) {
	ret := _m.Called()

	var r0 []entities.HouseResponseJoin
	if rf, ok := ret.Get(0).(func() []entities.HouseResponseJoin); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.HouseResponseJoin)
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

// UpdateHouse provides a mock function with given fields: id, update
func (_m *IHouseService) UpdateHouse(id uint, update entities.UpdateHouse) (entities.HouseResponse, error) {
	ret := _m.Called(id, update)

	var r0 entities.HouseResponse
	if rf, ok := ret.Get(0).(func(uint, entities.UpdateHouse) entities.HouseResponse); ok {
		r0 = rf(id, update)
	} else {
		r0 = ret.Get(0).(entities.HouseResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, entities.UpdateHouse) error); ok {
		r1 = rf(id, update)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewIHouseServiceT interface {
	mock.TestingT
	Cleanup(func())
}

// NewIHouseService creates a new instance of IHouseService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIHouseService(t NewIHouseServiceT) *IHouseService {
	mock := &IHouseService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
