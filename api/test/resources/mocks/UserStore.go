// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	resources "github.com/wspowell/snailmail/resources"
)

// UserStore is an autogenerated mock type for the UserStore type
type UserStore struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: username, attributes
func (_m *UserStore) CreateUser(username string, attributes resources.UserAttributes) (uint32, error) {
	ret := _m.Called(username, attributes)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(string, resources.UserAttributes) uint32); ok {
		r0 = rf(username, attributes)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, resources.UserAttributes) error); ok {
		r1 = rf(username, attributes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUser provides a mock function with given fields: userId
func (_m *UserStore) DeleteUser(userId uint32) error {
	ret := _m.Called(userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint32) error); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUser provides a mock function with given fields: userId
func (_m *UserStore) GetUser(userId uint32) (*resources.User, error) {
	ret := _m.Called(userId)

	var r0 *resources.User
	if rf, ok := ret.Get(0).(func(uint32) *resources.User); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resources.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint32) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: userId, newAttributes
func (_m *UserStore) UpdateUser(userId uint32, newAttributes resources.UserAttributes) error {
	ret := _m.Called(userId, newAttributes)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint32, resources.UserAttributes) error); ok {
		r0 = rf(userId, newAttributes)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}