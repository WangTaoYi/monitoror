// Code generated by mockery v1.0.0. DO NOT EDIT.

// If you want to rebuild this file, make mock-monitorable

package mocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/monitoror/monitoror/monitorable/config/models"
import tiles "github.com/monitoror/monitoror/models/tiles"
import utils "github.com/monitoror/monitoror/pkg/monitoror/utils"

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// Config provides a mock function with given fields: params
func (_m *Usecase) Config(params *models.ConfigParams) (*models.Config, error) {
	ret := _m.Called(params)

	var r0 *models.Config
	if rf, ok := ret.Get(0).(func(*models.ConfigParams) *models.Config); ok {
		r0 = rf(params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Config)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.ConfigParams) error); ok {
		r1 = rf(params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Hydrate provides a mock function with given fields: _a0
func (_m *Usecase) Hydrate(_a0 *models.Config) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Config) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Register provides a mock function with given fields: tileType, path, configValidator
func (_m *Usecase) Register(tileType tiles.TileType, path string, configValidator utils.Validator) {
	_m.Called(tileType, path, configValidator)
}

// Verify provides a mock function with given fields: _a0
func (_m *Usecase) Verify(_a0 *models.Config) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Config) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
