// Code generated by mockery v2.7.4. DO NOT EDIT.

package mocks

import (
	api "bidder/code/database/api"

	mock "github.com/stretchr/testify/mock"
)

// BudgetRepository is an autogenerated mock type for the BudgetRepository type
type BudgetRepository struct {
	mock.Mock
}

// FetchAll provides a mock function with given fields: consume
func (_m *BudgetRepository) FetchAll(consume func(api.Budget) error) error {
	ret := _m.Called(consume)

	var r0 error
	if rf, ok := ret.Get(0).(func(func(api.Budget) error) error); ok {
		r0 = rf(consume)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}