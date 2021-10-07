// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	runtime "github.com/go-openapi/runtime"
	robot "github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client/robot"
	mock "github.com/stretchr/testify/mock"
)

// MockRobotClientService is an autogenerated mock type for the ClientService type
type MockRobotClientService struct {
	mock.Mock
}

// CreateRobot provides a mock function with given fields: params, authInfo
func (_m *MockRobotClientService) CreateRobot(params *robot.CreateRobotParams, authInfo runtime.ClientAuthInfoWriter) (*robot.CreateRobotCreated, error) {
	ret := _m.Called(params, authInfo)

	var r0 *robot.CreateRobotCreated
	if rf, ok := ret.Get(0).(func(*robot.CreateRobotParams, runtime.ClientAuthInfoWriter) *robot.CreateRobotCreated); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*robot.CreateRobotCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*robot.CreateRobotParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteRobot provides a mock function with given fields: params, authInfo
func (_m *MockRobotClientService) DeleteRobot(params *robot.DeleteRobotParams, authInfo runtime.ClientAuthInfoWriter) (*robot.DeleteRobotOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *robot.DeleteRobotOK
	if rf, ok := ret.Get(0).(func(*robot.DeleteRobotParams, runtime.ClientAuthInfoWriter) *robot.DeleteRobotOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*robot.DeleteRobotOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*robot.DeleteRobotParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRobotByID provides a mock function with given fields: params, authInfo
func (_m *MockRobotClientService) GetRobotByID(params *robot.GetRobotByIDParams, authInfo runtime.ClientAuthInfoWriter) (*robot.GetRobotByIDOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *robot.GetRobotByIDOK
	if rf, ok := ret.Get(0).(func(*robot.GetRobotByIDParams, runtime.ClientAuthInfoWriter) *robot.GetRobotByIDOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*robot.GetRobotByIDOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*robot.GetRobotByIDParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListRobot provides a mock function with given fields: params, authInfo
func (_m *MockRobotClientService) ListRobot(params *robot.ListRobotParams, authInfo runtime.ClientAuthInfoWriter) (*robot.ListRobotOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *robot.ListRobotOK
	if rf, ok := ret.Get(0).(func(*robot.ListRobotParams, runtime.ClientAuthInfoWriter) *robot.ListRobotOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*robot.ListRobotOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*robot.ListRobotParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RefreshSec provides a mock function with given fields: params, authInfo
func (_m *MockRobotClientService) RefreshSec(params *robot.RefreshSecParams, authInfo runtime.ClientAuthInfoWriter) (*robot.RefreshSecOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *robot.RefreshSecOK
	if rf, ok := ret.Get(0).(func(*robot.RefreshSecParams, runtime.ClientAuthInfoWriter) *robot.RefreshSecOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*robot.RefreshSecOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*robot.RefreshSecParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *MockRobotClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}

// UpdateRobot provides a mock function with given fields: params, authInfo
func (_m *MockRobotClientService) UpdateRobot(params *robot.UpdateRobotParams, authInfo runtime.ClientAuthInfoWriter) (*robot.UpdateRobotOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *robot.UpdateRobotOK
	if rf, ok := ret.Get(0).(func(*robot.UpdateRobotParams, runtime.ClientAuthInfoWriter) *robot.UpdateRobotOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*robot.UpdateRobotOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*robot.UpdateRobotParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
