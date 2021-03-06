// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import "github.com/kyma-project/kyma/components/application-gateway/pkg/authorization"
import "github.com/stretchr/testify/mock"

// StrategyFactory is an autogenerated mock type for the StrategyFactory type
type StrategyFactory struct {
	mock.Mock
}

// Create provides a mock function with given fields: credentials
func (_m *StrategyFactory) Create(credentials *authorization.Credentials) authorization.Strategy {
	ret := _m.Called(credentials)

	var r0 authorization.Strategy
	if rf, ok := ret.Get(0).(func(*authorization.Credentials) authorization.Strategy); ok {
		r0 = rf(credentials)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(authorization.Strategy)
		}
	}

	return r0
}
