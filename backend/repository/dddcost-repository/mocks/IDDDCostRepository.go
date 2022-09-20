// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	customerror "github.com/beto-ouverney/falemais-telzir/customerror"

	entity "github.com/beto-ouverney/falemais-telzir/entity"

	mock "github.com/stretchr/testify/mock"
)

// IDDDCostRepository is an autogenerated mock type for the IDDDCostRepository type
type IDDDCostRepository struct {
	mock.Mock
}

// GetAllDDDCodes provides a mock function with given fields: ctx
func (_m *IDDDCostRepository) GetAllDDDCodes(ctx context.Context) (*[]entity.DDDCost, *customerror.CustomError) {
	ret := _m.Called(ctx)

	var r0 *[]entity.DDDCost
	if rf, ok := ret.Get(0).(func(context.Context) *[]entity.DDDCost); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]entity.DDDCost)
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

// GetByOriginDestination provides a mock function with given fields: ctx, origin, destination
func (_m *IDDDCostRepository) GetByOriginDestination(ctx context.Context, origin string, destination string) (*entity.DDDCost, *customerror.CustomError) {
	ret := _m.Called(ctx, origin, destination)

	var r0 *entity.DDDCost
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *entity.DDDCost); ok {
		r0 = rf(ctx, origin, destination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.DDDCost)
		}
	}

	var r1 *customerror.CustomError
	if rf, ok := ret.Get(1).(func(context.Context, string, string) *customerror.CustomError); ok {
		r1 = rf(ctx, origin, destination)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*customerror.CustomError)
		}
	}

	return r0, r1
}

type mockConstructorTestingTNewIDDDCostRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewIDDDCostRepository creates a new instance of IDDDCostRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIDDDCostRepository(t mockConstructorTestingTNewIDDDCostRepository) *IDDDCostRepository {
	mock := &IDDDCostRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
