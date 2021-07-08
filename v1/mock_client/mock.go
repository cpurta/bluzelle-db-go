// Code generated by MockGen. DO NOT EDIT.
// Source: ./client.go

// Package mock_client is a generated GoMock package.
package mock_client

import (
	context "context"
	reflect "reflect"

	bluzelledbgo "github.com/cpurta/bluzelle-db-go/v1"
	gomock "github.com/golang/mock/gomock"
)

// MockBluezelleClient is a mock of BluezelleClient interface.
type MockBluezelleClient struct {
	ctrl     *gomock.Controller
	recorder *MockBluezelleClientMockRecorder
}

// MockBluezelleClientMockRecorder is the mock recorder for MockBluezelleClient.
type MockBluezelleClientMockRecorder struct {
	mock *MockBluezelleClient
}

// NewMockBluezelleClient creates a new mock instance.
func NewMockBluezelleClient(ctrl *gomock.Controller) *MockBluezelleClient {
	mock := &MockBluezelleClient{ctrl: ctrl}
	mock.recorder = &MockBluezelleClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBluezelleClient) EXPECT() *MockBluezelleClientMockRecorder {
	return m.recorder
}

// Query mocks base method.
func (m *MockBluezelleClient) Query() bluzelledbgo.QueryClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query")
	ret0, _ := ret[0].(bluzelledbgo.QueryClient)
	return ret0
}

// Query indicates an expected call of Query.
func (mr *MockBluezelleClientMockRecorder) Query() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockBluezelleClient)(nil).Query))
}

// Transaction mocks base method.
func (m *MockBluezelleClient) Transaction() bluzelledbgo.TransactionClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transaction")
	ret0, _ := ret[0].(bluzelledbgo.TransactionClient)
	return ret0
}

// Transaction indicates an expected call of Transaction.
func (mr *MockBluezelleClientMockRecorder) Transaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transaction", reflect.TypeOf((*MockBluezelleClient)(nil).Transaction))
}

// MockTransactionOperation is a mock of TransactionOperation interface.
type MockTransactionOperation struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionOperationMockRecorder
}

// MockTransactionOperationMockRecorder is the mock recorder for MockTransactionOperation.
type MockTransactionOperationMockRecorder struct {
	mock *MockTransactionOperation
}

// NewMockTransactionOperation creates a new mock instance.
func NewMockTransactionOperation(ctrl *gomock.Controller) *MockTransactionOperation {
	mock := &MockTransactionOperation{ctrl: ctrl}
	mock.recorder = &MockTransactionOperationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionOperation) EXPECT() *MockTransactionOperationMockRecorder {
	return m.recorder
}

// PerformOperation mocks base method.
func (m *MockTransactionOperation) PerformOperation(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PerformOperation", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// PerformOperation indicates an expected call of PerformOperation.
func (mr *MockTransactionOperationMockRecorder) PerformOperation(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PerformOperation", reflect.TypeOf((*MockTransactionOperation)(nil).PerformOperation), ctx)
}
