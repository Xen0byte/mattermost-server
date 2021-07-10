// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost-server/v5/model"
	mock "github.com/stretchr/testify/mock"
)

// RemoteClusterStore is an autogenerated mock type for the RemoteClusterStore type
type RemoteClusterStore struct {
	mock.Mock
}

// Delete provides a mock function with given fields: remoteClusterID
func (_m *RemoteClusterStore) Delete(remoteClusterID string) (bool, error) {
	ret := _m.Called(remoteClusterID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(remoteClusterID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(remoteClusterID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: remoteClusterID
func (_m *RemoteClusterStore) Get(remoteClusterID string) (*model.RemoteCluster, error) {
	ret := _m.Called(remoteClusterID)

	var r0 *model.RemoteCluster
	if rf, ok := ret.Get(0).(func(string) *model.RemoteCluster); ok {
		r0 = rf(remoteClusterID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.RemoteCluster)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(remoteClusterID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: filter
func (_m *RemoteClusterStore) GetAll(filter model.RemoteClusterQueryFilter) ([]*model.RemoteCluster, error) {
	ret := _m.Called(filter)

	var r0 []*model.RemoteCluster
	if rf, ok := ret.Get(0).(func(model.RemoteClusterQueryFilter) []*model.RemoteCluster); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.RemoteCluster)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.RemoteClusterQueryFilter) error); ok {
		r1 = rf(filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: rc
func (_m *RemoteClusterStore) Save(rc *model.RemoteCluster) (*model.RemoteCluster, error) {
	ret := _m.Called(rc)

	var r0 *model.RemoteCluster
	if rf, ok := ret.Get(0).(func(*model.RemoteCluster) *model.RemoteCluster); ok {
		r0 = rf(rc)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.RemoteCluster)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.RemoteCluster) error); ok {
		r1 = rf(rc)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetLastPingAt provides a mock function with given fields: remoteClusterID
func (_m *RemoteClusterStore) SetLastPingAt(remoteClusterID string) error {
	ret := _m.Called(remoteClusterID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(remoteClusterID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: rc
func (_m *RemoteClusterStore) Update(rc *model.RemoteCluster) (*model.RemoteCluster, error) {
	ret := _m.Called(rc)

	var r0 *model.RemoteCluster
	if rf, ok := ret.Get(0).(func(*model.RemoteCluster) *model.RemoteCluster); ok {
		r0 = rf(rc)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.RemoteCluster)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.RemoteCluster) error); ok {
		r1 = rf(rc)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTopics provides a mock function with given fields: remoteClusterID, topics
func (_m *RemoteClusterStore) UpdateTopics(remoteClusterID string, topics string) (*model.RemoteCluster, error) {
	ret := _m.Called(remoteClusterID, topics)

	var r0 *model.RemoteCluster
	if rf, ok := ret.Get(0).(func(string, string) *model.RemoteCluster); ok {
		r0 = rf(remoteClusterID, topics)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.RemoteCluster)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(remoteClusterID, topics)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
