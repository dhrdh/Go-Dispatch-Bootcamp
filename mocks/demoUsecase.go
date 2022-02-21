// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	types "Go-Dispatch-Bootcamp/types"

	mock "github.com/stretchr/testify/mock"
)

// DemoUsecase is an autogenerated mock type for the DemoUsecase type
type DemoUsecase struct {
	mock.Mock
}

// Feed provides a mock function with given fields:
func (_m *DemoUsecase) Feed() ([][]string, error) {
	ret := _m.Called()

	var r0 [][]string
	if rf, ok := ret.Get(0).(func() [][]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]string)
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

// Fetch provides a mock function with given fields:
func (_m *DemoUsecase) Fetch() (*[]types.User, error) {
	ret := _m.Called()

	var r0 *[]types.User
	if rf, ok := ret.Get(0).(func() *[]types.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]types.User)
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

// FetchById provides a mock function with given fields: _a0
func (_m *DemoUsecase) FetchById(_a0 int) (*types.User, error) {
	ret := _m.Called(_a0)

	var r0 *types.User
	if rf, ok := ret.Get(0).(func(int) *types.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchConcurrently provides a mock function with given fields: _a0, _a1, _a2
func (_m *DemoUsecase) FetchConcurrently(_a0 string, _a1 int, _a2 int) (*[]types.User, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *[]types.User
	if rf, ok := ret.Get(0).(func(string, int, int) *[]types.User); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]types.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int, int) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUsersFromFeed provides a mock function with given fields:
func (_m *DemoUsecase) UpdateUsersFromFeed() (bool, error) {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
