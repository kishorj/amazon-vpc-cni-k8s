// Copyright 2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-vpc-cni-k8s/pkg/eniconfig (interfaces: ENIConfig)

// Package mock_eniconfig is a generated GoMock package.
package mock_eniconfig

import (
	reflect "reflect"

	v1alpha1 "github.com/aws/amazon-vpc-cni-k8s/pkg/apis/crd/v1alpha1"
	eniconfig "github.com/aws/amazon-vpc-cni-k8s/pkg/eniconfig"
	gomock "github.com/golang/mock/gomock"
)

// MockENIConfig is a mock of ENIConfig interface
type MockENIConfig struct {
	ctrl     *gomock.Controller
	recorder *MockENIConfigMockRecorder
}

// MockENIConfigMockRecorder is the mock recorder for MockENIConfig
type MockENIConfigMockRecorder struct {
	mock *MockENIConfig
}

// NewMockENIConfig creates a new mock instance
func NewMockENIConfig(ctrl *gomock.Controller) *MockENIConfig {
	mock := &MockENIConfig{ctrl: ctrl}
	mock.recorder = &MockENIConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockENIConfig) EXPECT() *MockENIConfigMockRecorder {
	return m.recorder
}

// Getter mocks base method
func (m *MockENIConfig) Getter() *eniconfig.ENIConfigInfo {
	ret := m.ctrl.Call(m, "Getter")
	ret0, _ := ret[0].(*eniconfig.ENIConfigInfo)
	return ret0
}

// Getter indicates an expected call of Getter
func (mr *MockENIConfigMockRecorder) Getter() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Getter", reflect.TypeOf((*MockENIConfig)(nil).Getter))
}

// MyENIConfig mocks base method
func (m *MockENIConfig) MyENIConfig() (*v1alpha1.ENIConfigSpec, error) {
	ret := m.ctrl.Call(m, "MyENIConfig")
	ret0, _ := ret[0].(*v1alpha1.ENIConfigSpec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MyENIConfig indicates an expected call of MyENIConfig
func (mr *MockENIConfigMockRecorder) MyENIConfig() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MyENIConfig", reflect.TypeOf((*MockENIConfig)(nil).MyENIConfig))
}