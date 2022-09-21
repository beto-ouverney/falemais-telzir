// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	customerror "github.com/beto-ouverney/falemais-telzir/customerror"

	mock "github.com/stretchr/testify/mock"
)

// IDDDCostController is an autogenerated mock type for the IDDDCostController type
type IDDDCostController struct {
	mock.Mock
}

// GetAllDDDCodes provides a mock function with given fields: ctx
func (_m *IDDDCostController) GetAllDDDCodes(ctx context.Context) (*[]byte, *customerror.CustomError) {
	ret := _m.Called(ctx)

	var r0 *[]byte
	if rf, ok := ret.Get(0).(func(context.Context) *[]byte); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]byte)
		}
	}

	var r1 *customerror.CustomError
	if rf, ok := ret.Get(1).(func(context.Context) *customerror.CustomError); ok {
		r1 = rf(ctx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*customerror.CustomError)
		}
	}

	return r0, r1
}

// GetCostByOriginDestination provides a mock function with given fields: ctx, origin, destination, min
func (_m *IDDDCostController) GetCostByOriginDestination(ctx context.Context, origin *int, destination *int, min *int) (*[]byte, *customerror.CustomError) {
	ret := _m.Called(ctx, origin, destination, min)

	var r0 *[]byte
	if rf, ok := ret.Get(0).(func(context.Context, *int, *int, *int) *[]byte); ok {
		r0 = rf(ctx, origin, destination, min)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]byte)
		}
	}

	var r1 *customerror.CustomError
	if rf, ok := ret.Get(1).(func(context.Context, *int, *int, *int) *customerror.CustomError); ok {
		r1 = rf(ctx, origin, destination, min)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*customerror.CustomError)
		}
	}

	return r0, r1
}

type mockConstructorTestingTNewIDDDCostController interface {
	mock.TestingT
	Cleanup(func())
}

// NewIDDDCostController creates a new instance of IDDDCostController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIDDDCostController(t mockConstructorTestingTNewIDDDCostController) *IDDDCostController {
	mock := &IDDDCostController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
