// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// AsyncEdgeWriter is an autogenerated mock type for the AsyncEdgeWriter type
type AsyncEdgeWriter struct {
	mock.Mock
}

type AsyncEdgeWriter_Expecter struct {
	mock *mock.Mock
}

func (_m *AsyncEdgeWriter) EXPECT() *AsyncEdgeWriter_Expecter {
	return &AsyncEdgeWriter_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields: ctx
func (_m *AsyncEdgeWriter) Close(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AsyncEdgeWriter_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type AsyncEdgeWriter_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
//   - ctx context.Context
func (_e *AsyncEdgeWriter_Expecter) Close(ctx interface{}) *AsyncEdgeWriter_Close_Call {
	return &AsyncEdgeWriter_Close_Call{Call: _e.mock.On("Close", ctx)}
}

func (_c *AsyncEdgeWriter_Close_Call) Run(run func(ctx context.Context)) *AsyncEdgeWriter_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *AsyncEdgeWriter_Close_Call) Return(_a0 error) *AsyncEdgeWriter_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AsyncEdgeWriter_Close_Call) RunAndReturn(run func(context.Context) error) *AsyncEdgeWriter_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Flush provides a mock function with given fields: ctx
func (_m *AsyncEdgeWriter) Flush(ctx context.Context) (chan struct{}, error) {
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

// AsyncEdgeWriter_Flush_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Flush'
type AsyncEdgeWriter_Flush_Call struct {
	*mock.Call
}

// Flush is a helper method to define mock.On call
//   - ctx context.Context
func (_e *AsyncEdgeWriter_Expecter) Flush(ctx interface{}) *AsyncEdgeWriter_Flush_Call {
	return &AsyncEdgeWriter_Flush_Call{Call: _e.mock.On("Flush", ctx)}
}

func (_c *AsyncEdgeWriter_Flush_Call) Run(run func(ctx context.Context)) *AsyncEdgeWriter_Flush_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *AsyncEdgeWriter_Flush_Call) Return(_a0 chan struct{}, _a1 error) *AsyncEdgeWriter_Flush_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AsyncEdgeWriter_Flush_Call) RunAndReturn(run func(context.Context) (chan struct{}, error)) *AsyncEdgeWriter_Flush_Call {
	_c.Call.Return(run)
	return _c
}

// Queue provides a mock function with given fields: ctx, e
func (_m *AsyncEdgeWriter) Queue(ctx context.Context, e interface{}) error {
	ret := _m.Called(ctx, e)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) error); ok {
		r0 = rf(ctx, e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AsyncEdgeWriter_Queue_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Queue'
type AsyncEdgeWriter_Queue_Call struct {
	*mock.Call
}

// Queue is a helper method to define mock.On call
//   - ctx context.Context
//   - e interface{}
func (_e *AsyncEdgeWriter_Expecter) Queue(ctx interface{}, e interface{}) *AsyncEdgeWriter_Queue_Call {
	return &AsyncEdgeWriter_Queue_Call{Call: _e.mock.On("Queue", ctx, e)}
}

func (_c *AsyncEdgeWriter_Queue_Call) Run(run func(ctx context.Context, e interface{})) *AsyncEdgeWriter_Queue_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(interface{}))
	})
	return _c
}

func (_c *AsyncEdgeWriter_Queue_Call) Return(_a0 error) *AsyncEdgeWriter_Queue_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AsyncEdgeWriter_Queue_Call) RunAndReturn(run func(context.Context, interface{}) error) *AsyncEdgeWriter_Queue_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewAsyncEdgeWriter interface {
	mock.TestingT
	Cleanup(func())
}

// NewAsyncEdgeWriter creates a new instance of AsyncEdgeWriter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAsyncEdgeWriter(t mockConstructorTestingTNewAsyncEdgeWriter) *AsyncEdgeWriter {
	mock := &AsyncEdgeWriter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}