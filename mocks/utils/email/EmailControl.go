// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// EmailControl is an autogenerated mock type for the EmailControl type
type EmailControl struct {
	mock.Mock
}

// SendEmail provides a mock function with given fields: sender, subject, body, recipient
func (_m *EmailControl) SendEmail(sender string, subject string, body string, recipient string) (string, error) {
	ret := _m.Called(sender, subject, body, recipient)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string, string, string) string); ok {
		r0 = rf(sender, subject, body, recipient)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, string) error); ok {
		r1 = rf(sender, subject, body, recipient)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewEmailControlT interface {
	mock.TestingT
	Cleanup(func())
}

// NewEmailControl creates a new instance of EmailControl. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEmailControl(t NewEmailControlT) *EmailControl {
	mock := &EmailControl{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
