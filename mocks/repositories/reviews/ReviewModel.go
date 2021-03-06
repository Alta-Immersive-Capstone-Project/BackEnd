// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	entities "kost/entities"

	mock "github.com/stretchr/testify/mock"
)

// ReviewModel is an autogenerated mock type for the ReviewModel type
type ReviewModel struct {
	mock.Mock
}

// Create provides a mock function with given fields: review
func (_m *ReviewModel) Create(review entities.Review) (entities.Review, error) {
	ret := _m.Called(review)

	var r0 entities.Review
	if rf, ok := ret.Get(0).(func(entities.Review) entities.Review); ok {
		r0 = rf(review)
	} else {
		r0 = ret.Get(0).(entities.Review)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entities.Review) error); ok {
		r1 = rf(review)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByHouseID provides a mock function with given fields: HouseID
func (_m *ReviewModel) GetByHouseID(HouseID uint) ([]entities.ReviewJoin, error) {
	ret := _m.Called(HouseID)

	var r0 []entities.ReviewJoin
	if rf, ok := ret.Get(0).(func(uint) []entities.ReviewJoin); ok {
		r0 = rf(HouseID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.ReviewJoin)
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

// GetRating provides a mock function with given fields: HouseID
func (_m *ReviewModel) GetRating(HouseID uint) ([]int, float32, error) {
	ret := _m.Called(HouseID)

	var r0 []int
	if rf, ok := ret.Get(0).(func(uint) []int); ok {
		r0 = rf(HouseID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int)
		}
	}

	var r1 float32
	if rf, ok := ret.Get(1).(func(uint) float32); ok {
		r1 = rf(HouseID)
	} else {
		r1 = ret.Get(1).(float32)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(uint) error); ok {
		r2 = rf(HouseID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type NewReviewModelT interface {
	mock.TestingT
	Cleanup(func())
}

// NewReviewModel creates a new instance of ReviewModel. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewReviewModel(t NewReviewModelT) *ReviewModel {
	mock := &ReviewModel{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
