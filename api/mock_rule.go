// Automatically generated by MockGen. DO NOT EDIT!
// Source: rule.go

package api

import (
	gomock "github.com/CrowdStrike/gomock/gomock"
)

// Mock of Rule interface
type MockRule struct {
	ctrl     *gomock.Controller
	recorder *_MockRuleRecorder
}

// Recorder for MockRule (not exported)
type _MockRuleRecorder struct {
	mock *MockRule
}

func NewMockRule(ctrl *gomock.Controller) *MockRule {
	mock := &MockRule{ctrl: ctrl}
	mock.recorder = &_MockRuleRecorder{mock}
	return mock
}

func (_m *MockRule) EXPECT() *_MockRuleRecorder {
	return _m.recorder
}

func (_m *MockRule) Name() string {
	ret := _m.ctrl.Call(_m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockRuleRecorder) Name() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Name")
}

func (_m *MockRule) Description() string {
	ret := _m.ctrl.Call(_m, "Description")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockRuleRecorder) Description() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Description")
}

func (_m *MockRule) Priority() int {
	ret := _m.ctrl.Call(_m, "Priority")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockRuleRecorder) Priority() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Priority")
}

func (_m *MockRule) Evaluate() bool {
	ret := _m.ctrl.Call(_m, "Evaluate")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockRuleRecorder) Evaluate() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Evaluate")
}

func (_m *MockRule) Execute() error {
	ret := _m.ctrl.Call(_m, "Execute")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockRuleRecorder) Execute() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Execute")
}
