// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import charlie "github.com/Bhinneka/alpha/api/service/domain/v1/charlie"
import context "context"
import mock "github.com/stretchr/testify/mock"
import response "github.com/Bhinneka/alpha/api/lib/response"

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

// AddCharlie provides a mock function with given fields: ctx, param
func (_m *UseCase) AddCharlie(ctx context.Context, param charlie.ParamAdd) response.Response {
	ret := _m.Called(ctx, param)

	var r0 response.Response
	if rf, ok := ret.Get(0).(func(context.Context, charlie.ParamAdd) response.Response); ok {
		r0 = rf(ctx, param)
	} else {
		r0 = ret.Get(0).(response.Response)
	}

	return r0
}

// DeleteCharlie provides a mock function with given fields: ctx, param
func (_m *UseCase) DeleteCharlie(ctx context.Context, param charlie.ParamDelete) response.Response {
	ret := _m.Called(ctx, param)

	var r0 response.Response
	if rf, ok := ret.Get(0).(func(context.Context, charlie.ParamDelete) response.Response); ok {
		r0 = rf(ctx, param)
	} else {
		r0 = ret.Get(0).(response.Response)
	}

	return r0
}

// GetCharlie provides a mock function with given fields: ctx, param
func (_m *UseCase) GetCharlie(ctx context.Context, param charlie.ParamGet) response.Response {
	ret := _m.Called(ctx, param)

	var r0 response.Response
	if rf, ok := ret.Get(0).(func(context.Context, charlie.ParamGet) response.Response); ok {
		r0 = rf(ctx, param)
	} else {
		r0 = ret.Get(0).(response.Response)
	}

	return r0
}

// UpdateCharlie provides a mock function with given fields: ctx, param
func (_m *UseCase) UpdateCharlie(ctx context.Context, param charlie.ParamUpdate) response.Response {
	ret := _m.Called(ctx, param)

	var r0 response.Response
	if rf, ok := ret.Get(0).(func(context.Context, charlie.ParamUpdate) response.Response); ok {
		r0 = rf(ctx, param)
	} else {
		r0 = ret.Get(0).(response.Response)
	}

	return r0
}
