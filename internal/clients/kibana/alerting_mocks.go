// Code generated by MockGen. DO NOT EDIT.
// Source: ./alerting.go
//
// Generated by this command:
//
//	mockgen -destination=./alerting_mocks.go -package=kibana -source ./alerting.go ApiClient
//

// Package kibana is a generated GoMock package.
package kibana

import (
	context "context"
	reflect "reflect"

	alerting "github.com/elastic/terraform-provider-elasticstack/generated/alerting"
	gomock "go.uber.org/mock/gomock"
)

// MockApiClient is a mock of ApiClient interface.
type MockApiClient struct {
	ctrl     *gomock.Controller
	recorder *MockApiClientMockRecorder
	isgomock struct{}
}

// MockApiClientMockRecorder is the mock recorder for MockApiClient.
type MockApiClientMockRecorder struct {
	mock *MockApiClient
}

// NewMockApiClient creates a new mock instance.
func NewMockApiClient(ctrl *gomock.Controller) *MockApiClient {
	mock := &MockApiClient{ctrl: ctrl}
	mock.recorder = &MockApiClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApiClient) EXPECT() *MockApiClientMockRecorder {
	return m.recorder
}

// GetAlertingClient mocks base method.
func (m *MockApiClient) GetAlertingClient() (alerting.AlertingAPI, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAlertingClient")
	ret0, _ := ret[0].(alerting.AlertingAPI)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAlertingClient indicates an expected call of GetAlertingClient.
func (mr *MockApiClientMockRecorder) GetAlertingClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAlertingClient", reflect.TypeOf((*MockApiClient)(nil).GetAlertingClient))
}

// SetAlertingAuthContext mocks base method.
func (m *MockApiClient) SetAlertingAuthContext(arg0 context.Context) context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetAlertingAuthContext", arg0)
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// SetAlertingAuthContext indicates an expected call of SetAlertingAuthContext.
func (mr *MockApiClientMockRecorder) SetAlertingAuthContext(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAlertingAuthContext", reflect.TypeOf((*MockApiClient)(nil).SetAlertingAuthContext), arg0)
}
