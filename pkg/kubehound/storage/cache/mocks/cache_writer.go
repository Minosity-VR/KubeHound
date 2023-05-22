// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	cache "github.com/DataDog/KubeHound/pkg/kubehound/storage/cache"

	mock "github.com/stretchr/testify/mock"
)

// AsyncWriter is an autogenerated mock type for the AsyncWriter type
type AsyncWriter struct {
	mock.Mock
}

type AsyncWriter_Expecter struct {
	mock *mock.Mock
}

func (_m *AsyncWriter) EXPECT() *AsyncWriter_Expecter {
	return &AsyncWriter_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields: ctx
func (_m *AsyncWriter) Close(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AsyncWriter_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type AsyncWriter_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
//   - ctx context.Context
func (_e *AsyncWriter_Expecter) Close(ctx interface{}) *AsyncWriter_Close_Call {
	return &AsyncWriter_Close_Call{Call: _e.mock.On("Close", ctx)}
}

func (_c *AsyncWriter_Close_Call) Run(run func(ctx context.Context)) *AsyncWriter_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *AsyncWriter_Close_Call) Return(_a0 error) *AsyncWriter_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AsyncWriter_Close_Call) RunAndReturn(run func(context.Context) error) *AsyncWriter_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Flush provides a mock function with given fields: ctx
func (_m *AsyncWriter) Flush(ctx context.Context) (chan struct{}, error) {
	ret := _m.Called(ctx)

	var r0 chan struct{}
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (chan struct{}, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) chan struct{}); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan struct{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AsyncWriter_Flush_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Flush'
type AsyncWriter_Flush_Call struct {
	*mock.Call
}

// Flush is a helper method to define mock.On call
//   - ctx context.Context
func (_e *AsyncWriter_Expecter) Flush(ctx interface{}) *AsyncWriter_Flush_Call {
	return &AsyncWriter_Flush_Call{Call: _e.mock.On("Flush", ctx)}
}

func (_c *AsyncWriter_Flush_Call) Run(run func(ctx context.Context)) *AsyncWriter_Flush_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *AsyncWriter_Flush_Call) Return(_a0 chan struct{}, _a1 error) *AsyncWriter_Flush_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AsyncWriter_Flush_Call) RunAndReturn(run func(context.Context) (chan struct{}, error)) *AsyncWriter_Flush_Call {
	_c.Call.Return(run)
	return _c
}

// Queue provides a mock function with given fields: ctx, key, value
func (_m *AsyncWriter) Queue(ctx context.Context, key cache.CacheKey, value interface{}) error {
	ret := _m.Called(ctx, key, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, cache.CacheKey, interface{}) error); ok {
		r0 = rf(ctx, key, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AsyncWriter_Queue_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Queue'
type AsyncWriter_Queue_Call struct {
	*mock.Call
}

// Queue is a helper method to define mock.On call
//   - ctx context.Context
//   - key cache.CacheKey
//   - value interface{}
func (_e *AsyncWriter_Expecter) Queue(ctx interface{}, key interface{}, value interface{}) *AsyncWriter_Queue_Call {
	return &AsyncWriter_Queue_Call{Call: _e.mock.On("Queue", ctx, key, value)}
}

func (_c *AsyncWriter_Queue_Call) Run(run func(ctx context.Context, key cache.CacheKey, value interface{})) *AsyncWriter_Queue_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(cache.CacheKey), args[2].(interface{}))
	})
	return _c
}

func (_c *AsyncWriter_Queue_Call) Return(_a0 error) *AsyncWriter_Queue_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AsyncWriter_Queue_Call) RunAndReturn(run func(context.Context, cache.CacheKey, interface{}) error) *AsyncWriter_Queue_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewAsyncWriter interface {
	mock.TestingT
	Cleanup(func())
}

// NewAsyncWriter creates a new instance of AsyncWriter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAsyncWriter(t mockConstructorTestingTNewAsyncWriter) *AsyncWriter {
	mock := &AsyncWriter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
