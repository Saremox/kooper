// Code generated by mockery v1.0.0. DO NOT EDIT.

package controller

import mock "github.com/stretchr/testify/mock"

// Controller is an autogenerated mock type for the Controller type
type Controller struct {
	mock.Mock
}

// Run provides a mock function with given fields: stopper
func (_m *Controller) Run(stopper <-chan struct{}) error {
	ret := _m.Called(stopper)

	var r0 error
	if rf, ok := ret.Get(0).(func(<-chan struct{}) error); ok {
		r0 = rf(stopper)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}