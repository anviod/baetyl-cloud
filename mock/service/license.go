// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/baetyl/baetyl-cloud/v2/service (interfaces: LicenseService)

// Package service is a generated GoMock package.
package service

import (
	plugin "github.com/baetyl/baetyl-cloud/v2/plugin"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockLicenseService is a mock of LicenseService interface
type MockLicenseService struct {
	ctrl     *gomock.Controller
	recorder *MockLicenseServiceMockRecorder
}

// MockLicenseServiceMockRecorder is the mock recorder for MockLicenseService
type MockLicenseServiceMockRecorder struct {
	mock *MockLicenseService
}

// NewMockLicenseService creates a new mock instance
func NewMockLicenseService(ctrl *gomock.Controller) *MockLicenseService {
	mock := &MockLicenseService{ctrl: ctrl}
	mock.recorder = &MockLicenseServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLicenseService) EXPECT() *MockLicenseServiceMockRecorder {
	return m.recorder
}

// AcquireQuota mocks base method
func (m *MockLicenseService) AcquireQuota(arg0, arg1 string, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcquireQuota", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AcquireQuota indicates an expected call of AcquireQuota
func (mr *MockLicenseServiceMockRecorder) AcquireQuota(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcquireQuota", reflect.TypeOf((*MockLicenseService)(nil).AcquireQuota), arg0, arg1, arg2)
}

// CheckLicense mocks base method
func (m *MockLicenseService) CheckLicense() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckLicense")
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckLicense indicates an expected call of CheckLicense
func (mr *MockLicenseServiceMockRecorder) CheckLicense() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckLicense", reflect.TypeOf((*MockLicenseService)(nil).CheckLicense))
}

// CheckQuota mocks base method
func (m *MockLicenseService) CheckQuota(arg0 string, arg1 plugin.QuotaCollector) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckQuota", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckQuota indicates an expected call of CheckQuota
func (mr *MockLicenseServiceMockRecorder) CheckQuota(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckQuota", reflect.TypeOf((*MockLicenseService)(nil).CheckQuota), arg0, arg1)
}

// CreateQuota mocks base method
func (m *MockLicenseService) CreateQuota(arg0 string, arg1 map[string]int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateQuota", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateQuota indicates an expected call of CreateQuota
func (mr *MockLicenseServiceMockRecorder) CreateQuota(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateQuota", reflect.TypeOf((*MockLicenseService)(nil).CreateQuota), arg0, arg1)
}

// DeleteQuota mocks base method
func (m *MockLicenseService) DeleteQuota(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteQuota", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteQuota indicates an expected call of DeleteQuota
func (mr *MockLicenseServiceMockRecorder) DeleteQuota(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteQuota", reflect.TypeOf((*MockLicenseService)(nil).DeleteQuota), arg0, arg1)
}

// DeleteQuotaByNamespace mocks base method
func (m *MockLicenseService) DeleteQuotaByNamespace(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteQuotaByNamespace", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteQuotaByNamespace indicates an expected call of DeleteQuotaByNamespace
func (mr *MockLicenseServiceMockRecorder) DeleteQuotaByNamespace(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteQuotaByNamespace", reflect.TypeOf((*MockLicenseService)(nil).DeleteQuotaByNamespace), arg0)
}

// GetDefaultQuotas mocks base method
func (m *MockLicenseService) GetDefaultQuotas(arg0 string) (map[string]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDefaultQuotas", arg0)
	ret0, _ := ret[0].(map[string]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDefaultQuotas indicates an expected call of GetDefaultQuotas
func (mr *MockLicenseServiceMockRecorder) GetDefaultQuotas(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDefaultQuotas", reflect.TypeOf((*MockLicenseService)(nil).GetDefaultQuotas), arg0)
}

// GetQuota mocks base method
func (m *MockLicenseService) GetQuota(arg0 string) (map[string]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQuota", arg0)
	ret0, _ := ret[0].(map[string]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQuota indicates an expected call of GetQuota
func (mr *MockLicenseServiceMockRecorder) GetQuota(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQuota", reflect.TypeOf((*MockLicenseService)(nil).GetQuota), arg0)
}

// ProtectCode mocks base method
func (m *MockLicenseService) ProtectCode() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProtectCode")
	ret0, _ := ret[0].(error)
	return ret0
}

// ProtectCode indicates an expected call of ProtectCode
func (mr *MockLicenseServiceMockRecorder) ProtectCode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProtectCode", reflect.TypeOf((*MockLicenseService)(nil).ProtectCode))
}

// ReleaseQuota mocks base method
func (m *MockLicenseService) ReleaseQuota(arg0, arg1 string, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReleaseQuota", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReleaseQuota indicates an expected call of ReleaseQuota
func (mr *MockLicenseServiceMockRecorder) ReleaseQuota(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseQuota", reflect.TypeOf((*MockLicenseService)(nil).ReleaseQuota), arg0, arg1, arg2)
}

// UpdateQuota mocks base method
func (m *MockLicenseService) UpdateQuota(arg0, arg1 string, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateQuota", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateQuota indicates an expected call of UpdateQuota
func (mr *MockLicenseServiceMockRecorder) UpdateQuota(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateQuota", reflect.TypeOf((*MockLicenseService)(nil).UpdateQuota), arg0, arg1, arg2)
}
