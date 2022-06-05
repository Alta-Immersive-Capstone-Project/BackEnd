// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	entities "kost/entities"

	mock "github.com/stretchr/testify/mock"
)

// ForgotInterface is an autogenerated mock type for the ForgotInterface type
type ForgotInterface struct {
	mock.Mock
}

// FindUserByEmail provides a mock function with given fields: email
func (_m *ForgotInterface) FindUserByEmail(email string) (entities.User, error) {
	ret := _m.Called(email)

	var r0 entities.User
	if rf, ok := ret.Get(0).(func(string) entities.User); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(entities.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResetPassword provides a mock function with given fields: id, password
func (_m *ForgotInterface) ResetPassword(id int, password string) (entities.User, error) {
	ret := _m.Called(id, password)

	var r0 entities.User
	if rf, ok := ret.Get(0).(func(int, string) entities.User); ok {
		r0 = rf(id, password)
	} else {
		r0 = ret.Get(0).(entities.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, string) error); ok {
		r1 = rf(id, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewForgotInterfaceT interface {
	mock.TestingT
	Cleanup(func())
}

// NewForgotInterface creates a new instance of ForgotInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewForgotInterface(t NewForgotInterfaceT) *ForgotInterface {
	mock := &ForgotInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
