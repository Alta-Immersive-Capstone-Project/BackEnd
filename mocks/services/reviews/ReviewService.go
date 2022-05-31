// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	entities "kost/entities"

	mock "github.com/stretchr/testify/mock"
)

// ReviewService is an autogenerated mock type for the ReviewService type
type ReviewService struct {
	mock.Mock
}

// AddComment provides a mock function with given fields: customer_id, request
func (_m *ReviewService) AddComment(customer_id uint, request entities.ReviewRequest) (entities.ReviewResponse, error) {
	ret := _m.Called(customer_id, request)

	var r0 entities.ReviewResponse
	if rf, ok := ret.Get(0).(func(uint, entities.ReviewRequest) entities.ReviewResponse); ok {
		r0 = rf(customer_id, request)
	} else {
		r0 = ret.Get(0).(entities.ReviewResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, entities.ReviewRequest) error); ok {
		r1 = rf(customer_id, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByRoomID provides a mock function with given fields: room_id
func (_m *ReviewService) GetByRoomID(room_id uint) ([]entities.ReviewGetResponse, error) {
	ret := _m.Called(room_id)

	var r0 []entities.ReviewGetResponse
	if rf, ok := ret.Get(0).(func(uint) []entities.ReviewGetResponse); ok {
		r0 = rf(room_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.ReviewGetResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(room_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRating provides a mock function with given fields: room_id
func (_m *ReviewService) GetRating(room_id uint) float32 {
	ret := _m.Called(room_id)

	var r0 float32
	if rf, ok := ret.Get(0).(func(uint) float32); ok {
		r0 = rf(room_id)
	} else {
		r0 = ret.Get(0).(float32)
	}

	return r0
}

type NewReviewServiceT interface {
	mock.TestingT
	Cleanup(func())
}

// NewReviewService creates a new instance of ReviewService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewReviewService(t NewReviewServiceT) *ReviewService {
	mock := &ReviewService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
