// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	account "github.com/ha-dev/banking/internal/entities/account"
)

// MockILocalDB is a mock of ILocalDB interface.
type MockILocalDB struct {
	ctrl     *gomock.Controller
	recorder *MockILocalDBMockRecorder
}

// MockILocalDBMockRecorder is the mock recorder for MockILocalDB.
type MockILocalDBMockRecorder struct {
	mock *MockILocalDB
}

// NewMockILocalDB creates a new mock instance.
func NewMockILocalDB(ctrl *gomock.Controller) *MockILocalDB {
	mock := &MockILocalDB{ctrl: ctrl}
	mock.recorder = &MockILocalDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockILocalDB) EXPECT() *MockILocalDBMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockILocalDB) Get(FieldName string) (*account.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", FieldName)
	ret0, _ := ret[0].(*account.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockILocalDBMockRecorder) Get(FieldName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockILocalDB)(nil).Get), FieldName)
}

// List mocks base method.
func (m *MockILocalDB) List() []*account.Account {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*account.Account)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockILocalDBMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockILocalDB)(nil).List))
}

// Update mocks base method.
func (m *MockILocalDB) Update(acc *account.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", acc)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockILocalDBMockRecorder) Update(acc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockILocalDB)(nil).Update), acc)
}
