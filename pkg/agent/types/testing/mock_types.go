// Copyright 2022 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: antrea.io/antrea/pkg/agent/types (interfaces: McastNetworkPolicyController)

// Package testing is a generated GoMock package.
package testing

import (
	types "antrea.io/antrea/pkg/agent/types"
	gomock "github.com/golang/mock/gomock"
	net "net"
	reflect "reflect"
)

// MockMcastNetworkPolicyController is a mock of McastNetworkPolicyController interface
type MockMcastNetworkPolicyController struct {
	ctrl     *gomock.Controller
	recorder *MockMcastNetworkPolicyControllerMockRecorder
}

// MockMcastNetworkPolicyControllerMockRecorder is the mock recorder for MockMcastNetworkPolicyController
type MockMcastNetworkPolicyControllerMockRecorder struct {
	mock *MockMcastNetworkPolicyController
}

// NewMockMcastNetworkPolicyController creates a new mock instance
func NewMockMcastNetworkPolicyController(ctrl *gomock.Controller) *MockMcastNetworkPolicyController {
	mock := &MockMcastNetworkPolicyController{ctrl: ctrl}
	mock.recorder = &MockMcastNetworkPolicyControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMcastNetworkPolicyController) EXPECT() *MockMcastNetworkPolicyControllerMockRecorder {
	return m.recorder
}

// GetIGMPNPRuleInfo mocks base method
func (m *MockMcastNetworkPolicyController) GetIGMPNPRuleInfo(arg0, arg1 string, arg2 net.IP, arg3 byte) (*types.IGMPNPRuleInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIGMPNPRuleInfo", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*types.IGMPNPRuleInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIGMPNPRuleInfo indicates an expected call of GetIGMPNPRuleInfo
func (mr *MockMcastNetworkPolicyControllerMockRecorder) GetIGMPNPRuleInfo(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIGMPNPRuleInfo", reflect.TypeOf((*MockMcastNetworkPolicyController)(nil).GetIGMPNPRuleInfo), arg0, arg1, arg2, arg3)
}
