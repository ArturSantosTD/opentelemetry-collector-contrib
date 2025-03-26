// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	time "time"

	aerospike "github.com/aerospike/aerospike-client-go/v8"
	mock "github.com/stretchr/testify/mock"
)

// Asconn is an autogenerated mock type for the asconn type
type Asconn struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Asconn) Close() {
	_m.Called()
}

// Login provides a mock function with given fields: _a0
func (_m *Asconn) Login(_a0 *aerospike.ClientPolicy) aerospike.Error {
	ret := _m.Called(_a0)

	var r0 aerospike.Error
	if rf, ok := ret.Get(0).(func(*aerospike.ClientPolicy) aerospike.Error); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(aerospike.Error)
		}
	}

	return r0
}

// RequestInfo provides a mock function with given fields: _a0
func (_m *Asconn) RequestInfo(_a0 ...string) (map[string]string, aerospike.Error) {
	_va := make([]any, len(_a0))
	for _i := range _a0 {
		_va[_i] = _a0[_i]
	}
	var _ca []any
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func(...string) map[string]string); ok {
		r0 = rf(_a0...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	var r1 aerospike.Error
	if rf, ok := ret.Get(1).(func(...string) aerospike.Error); ok {
		r1 = rf(_a0...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(aerospike.Error)
		}
	}

	return r0, r1
}

// SetTimeout provides a mock function with given fields: _a0, _a1
func (_m *Asconn) SetTimeout(_a0 time.Time, _a1 time.Duration) aerospike.Error {
	ret := _m.Called(_a0, _a1)

	var r0 aerospike.Error
	if rf, ok := ret.Get(0).(func(time.Time, time.Duration) aerospike.Error); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(aerospike.Error)
		}
	}

	return r0
}

type mockConstructorTestingTNewAsconn interface {
	mock.TestingT
	Cleanup(func())
}

// NewAsconn creates a new instance of Asconn. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAsconn(t mockConstructorTestingTNewAsconn) *Asconn {
	mock := &Asconn{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
