// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	account "github.com/ha-dev/banking/internal/entities/account"
	accounts "github.com/ha-dev/banking/services/accounts"
)

// MockIAccount is a mock of IAccount interface.
type MockIAccount struct {
	ctrl     *gomock.Controller
	recorder *MockIAccountMockRecorder
}

// MockIAccountMockRecorder is the mock recorder for MockIAccount.
type MockIAccountMockRecorder struct {
	mock *MockIAccount
}

// NewMockIAccount creates a new mock instance.
func NewMockIAccount(ctrl *gomock.Controller) *MockIAccount {
	mock := &MockIAccount{ctrl: ctrl}
	mock.recorder = &MockIAccountMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAccount) EXPECT() *MockIAccountMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockIAccount) Get(ID string) (*account.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ID)
	ret0, _ := ret[0].(*account.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIAccountMockRecorder) Get(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIAccount)(nil).Get), ID)
}

// ListAll mocks base method.
func (m *MockIAccount) ListAll() ([]*account.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAll")
	ret0, _ := ret[0].([]*account.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAll indicates an expected call of ListAll.
func (mr *MockIAccountMockRecorder) ListAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAll", reflect.TypeOf((*MockIAccount)(nil).ListAll))
}

// Transfer mocks base method.
func (m *MockIAccount) Transfer(param *accounts.TransferParams) (*account.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transfer", param)
	ret0, _ := ret[0].(*account.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Transfer indicates an expected call of Transfer.
func (mr *MockIAccountMockRecorder) Transfer(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transfer", reflect.TypeOf((*MockIAccount)(nil).Transfer), param)
}
