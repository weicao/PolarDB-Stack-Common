/* 
*Copyright (c) 2019-2021, Alibaba Group Holding Limited;
*Licensed under the Apache License, Version 2.0 (the "License");
*you may not use this file except in compliance with the License.
*You may obtain a copy of the License at

*   http://www.apache.org/licenses/LICENSE-2.0

*Unless required by applicable law or agreed to in writing, software
*distributed under the License is distributed on an "AS IS" BASIS,
*WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*See the License for the specific language governing permissions and
*limitations under the License.
 */


// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/qinchao/go/pkg/mod/sigs.k8s.io/controller-runtime@v0.5.0/pkg/manager/manager.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	meta "k8s.io/apimachinery/pkg/api/meta"
	runtime "k8s.io/apimachinery/pkg/runtime"
	rest "k8s.io/client-go/rest"
	record "k8s.io/client-go/tools/record"
	cache "sigs.k8s.io/controller-runtime/pkg/cache"
	client "sigs.k8s.io/controller-runtime/pkg/client"
	healthz "sigs.k8s.io/controller-runtime/pkg/healthz"
	manager "sigs.k8s.io/controller-runtime/pkg/manager"
	webhook "sigs.k8s.io/controller-runtime/pkg/webhook"
)

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockManager) Add(arg0 manager.Runnable) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockManagerMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockManager)(nil).Add), arg0)
}

// AddHealthzCheck mocks base method.
func (m *MockManager) AddHealthzCheck(name string, check healthz.Checker) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddHealthzCheck", name, check)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddHealthzCheck indicates an expected call of AddHealthzCheck.
func (mr *MockManagerMockRecorder) AddHealthzCheck(name, check interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddHealthzCheck", reflect.TypeOf((*MockManager)(nil).AddHealthzCheck), name, check)
}

// AddReadyzCheck mocks base method.
func (m *MockManager) AddReadyzCheck(name string, check healthz.Checker) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddReadyzCheck", name, check)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddReadyzCheck indicates an expected call of AddReadyzCheck.
func (mr *MockManagerMockRecorder) AddReadyzCheck(name, check interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddReadyzCheck", reflect.TypeOf((*MockManager)(nil).AddReadyzCheck), name, check)
}

// GetAPIReader mocks base method.
func (m *MockManager) GetAPIReader() client.Reader {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAPIReader")
	ret0, _ := ret[0].(client.Reader)
	return ret0
}

// GetAPIReader indicates an expected call of GetAPIReader.
func (mr *MockManagerMockRecorder) GetAPIReader() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAPIReader", reflect.TypeOf((*MockManager)(nil).GetAPIReader))
}

// GetCache mocks base method.
func (m *MockManager) GetCache() cache.Cache {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCache")
	ret0, _ := ret[0].(cache.Cache)
	return ret0
}

// GetCache indicates an expected call of GetCache.
func (mr *MockManagerMockRecorder) GetCache() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCache", reflect.TypeOf((*MockManager)(nil).GetCache))
}

// GetClient mocks base method.
func (m *MockManager) GetClient() client.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClient")
	ret0, _ := ret[0].(client.Client)
	return ret0
}

// GetClient indicates an expected call of GetClient.
func (mr *MockManagerMockRecorder) GetClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClient", reflect.TypeOf((*MockManager)(nil).GetClient))
}

// GetConfig mocks base method.
func (m *MockManager) GetConfig() *rest.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig")
	ret0, _ := ret[0].(*rest.Config)
	return ret0
}

// GetConfig indicates an expected call of GetConfig.
func (mr *MockManagerMockRecorder) GetConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockManager)(nil).GetConfig))
}

// GetEventRecorderFor mocks base method.
func (m *MockManager) GetEventRecorderFor(name string) record.EventRecorder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventRecorderFor", name)
	ret0, _ := ret[0].(record.EventRecorder)
	return ret0
}

// GetEventRecorderFor indicates an expected call of GetEventRecorderFor.
func (mr *MockManagerMockRecorder) GetEventRecorderFor(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventRecorderFor", reflect.TypeOf((*MockManager)(nil).GetEventRecorderFor), name)
}

// GetFieldIndexer mocks base method.
func (m *MockManager) GetFieldIndexer() client.FieldIndexer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFieldIndexer")
	ret0, _ := ret[0].(client.FieldIndexer)
	return ret0
}

// GetFieldIndexer indicates an expected call of GetFieldIndexer.
func (mr *MockManagerMockRecorder) GetFieldIndexer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFieldIndexer", reflect.TypeOf((*MockManager)(nil).GetFieldIndexer))
}

// GetRESTMapper mocks base method.
func (m *MockManager) GetRESTMapper() meta.RESTMapper {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRESTMapper")
	ret0, _ := ret[0].(meta.RESTMapper)
	return ret0
}

// GetRESTMapper indicates an expected call of GetRESTMapper.
func (mr *MockManagerMockRecorder) GetRESTMapper() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRESTMapper", reflect.TypeOf((*MockManager)(nil).GetRESTMapper))
}

// GetScheme mocks base method.
func (m *MockManager) GetScheme() *runtime.Scheme {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScheme")
	ret0, _ := ret[0].(*runtime.Scheme)
	return ret0
}

// GetScheme indicates an expected call of GetScheme.
func (mr *MockManagerMockRecorder) GetScheme() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScheme", reflect.TypeOf((*MockManager)(nil).GetScheme))
}

// GetWebhookServer mocks base method.
func (m *MockManager) GetWebhookServer() *webhook.Server {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWebhookServer")
	ret0, _ := ret[0].(*webhook.Server)
	return ret0
}

// GetWebhookServer indicates an expected call of GetWebhookServer.
func (mr *MockManagerMockRecorder) GetWebhookServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWebhookServer", reflect.TypeOf((*MockManager)(nil).GetWebhookServer))
}

// SetFields mocks base method.
func (m *MockManager) SetFields(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetFields", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetFields indicates an expected call of SetFields.
func (mr *MockManagerMockRecorder) SetFields(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFields", reflect.TypeOf((*MockManager)(nil).SetFields), arg0)
}

// Start mocks base method.
func (m *MockManager) Start(arg0 <-chan struct{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockManagerMockRecorder) Start(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockManager)(nil).Start), arg0)
}

// MockRunnable is a mock of Runnable interface.
type MockRunnable struct {
	ctrl     *gomock.Controller
	recorder *MockRunnableMockRecorder
}

// MockRunnableMockRecorder is the mock recorder for MockRunnable.
type MockRunnableMockRecorder struct {
	mock *MockRunnable
}

// NewMockRunnable creates a new mock instance.
func NewMockRunnable(ctrl *gomock.Controller) *MockRunnable {
	mock := &MockRunnable{ctrl: ctrl}
	mock.recorder = &MockRunnableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRunnable) EXPECT() *MockRunnableMockRecorder {
	return m.recorder
}

// Start mocks base method.
func (m *MockRunnable) Start(arg0 <-chan struct{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockRunnableMockRecorder) Start(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockRunnable)(nil).Start), arg0)
}

// MockLeaderElectionRunnable is a mock of LeaderElectionRunnable interface.
type MockLeaderElectionRunnable struct {
	ctrl     *gomock.Controller
	recorder *MockLeaderElectionRunnableMockRecorder
}

// MockLeaderElectionRunnableMockRecorder is the mock recorder for MockLeaderElectionRunnable.
type MockLeaderElectionRunnableMockRecorder struct {
	mock *MockLeaderElectionRunnable
}

// NewMockLeaderElectionRunnable creates a new mock instance.
func NewMockLeaderElectionRunnable(ctrl *gomock.Controller) *MockLeaderElectionRunnable {
	mock := &MockLeaderElectionRunnable{ctrl: ctrl}
	mock.recorder = &MockLeaderElectionRunnableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLeaderElectionRunnable) EXPECT() *MockLeaderElectionRunnableMockRecorder {
	return m.recorder
}

// NeedLeaderElection mocks base method.
func (m *MockLeaderElectionRunnable) NeedLeaderElection() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NeedLeaderElection")
	ret0, _ := ret[0].(bool)
	return ret0
}

// NeedLeaderElection indicates an expected call of NeedLeaderElection.
func (mr *MockLeaderElectionRunnableMockRecorder) NeedLeaderElection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NeedLeaderElection", reflect.TypeOf((*MockLeaderElectionRunnable)(nil).NeedLeaderElection))
}
