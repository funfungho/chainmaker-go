// Code generated by MockGen. DO NOT EDIT.
// Source: vm_interface.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	common "chainmaker.org/chainmaker-go/pb/protogo/common"
	protocol "chainmaker.org/chainmaker-go/protocol"
	gomock "github.com/golang/mock/gomock"
)

// MockVmManager is a mock of VmManager interface.
type MockVmManager struct {
	ctrl     *gomock.Controller
	recorder *MockVmManagerMockRecorder
}

// MockVmManagerMockRecorder is the mock recorder for MockVmManager.
type MockVmManagerMockRecorder struct {
	mock *MockVmManager
}

// NewMockVmManager creates a new mock instance.
func NewMockVmManager(ctrl *gomock.Controller) *MockVmManager {
	mock := &MockVmManager{ctrl: ctrl}
	mock.recorder = &MockVmManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVmManager) EXPECT() *MockVmManagerMockRecorder {
	return m.recorder
}

// GetAccessControl mocks base method.
func (m *MockVmManager) GetAccessControl() protocol.AccessControlProvider {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccessControl")
	ret0, _ := ret[0].(protocol.AccessControlProvider)
	return ret0
}

// GetAccessControl indicates an expected call of GetAccessControl.
func (mr *MockVmManagerMockRecorder) GetAccessControl() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessControl", reflect.TypeOf((*MockVmManager)(nil).GetAccessControl))
}

// GetChainNodesInfoProvider mocks base method.
func (m *MockVmManager) GetChainNodesInfoProvider() protocol.ChainNodesInfoProvider {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChainNodesInfoProvider")
	ret0, _ := ret[0].(protocol.ChainNodesInfoProvider)
	return ret0
}

// GetChainNodesInfoProvider indicates an expected call of GetChainNodesInfoProvider.
func (mr *MockVmManagerMockRecorder) GetChainNodesInfoProvider() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChainNodesInfoProvider", reflect.TypeOf((*MockVmManager)(nil).GetChainNodesInfoProvider))
}

// RunContract mocks base method.
func (m *MockVmManager) RunContract(contractId *common.ContractId, method string, byteCode []byte, parameters map[string]string, txContext protocol.TxSimContext, gasUsed uint64, refTxType common.TxType) (*common.ContractResult, common.TxStatusCode) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunContract", contractId, method, byteCode, parameters, txContext, gasUsed, refTxType)
	ret0, _ := ret[0].(*common.ContractResult)
	ret1, _ := ret[1].(common.TxStatusCode)
	return ret0, ret1
}

// RunContract indicates an expected call of RunContract.
func (mr *MockVmManagerMockRecorder) RunContract(contractId, method, byteCode, parameters, txContext, gasUsed, refTxType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunContract", reflect.TypeOf((*MockVmManager)(nil).RunContract), contractId, method, byteCode, parameters, txContext, gasUsed, refTxType)
}

// MockContractWacsiCommon is a mock of ContractWacsiCommon interface.
type MockContractWacsiCommon struct {
	ctrl     *gomock.Controller
	recorder *MockContractWacsiCommonMockRecorder
}

// MockContractWacsiCommonMockRecorder is the mock recorder for MockContractWacsiCommon.
type MockContractWacsiCommonMockRecorder struct {
	mock *MockContractWacsiCommon
}

// NewMockContractWacsiCommon creates a new mock instance.
func NewMockContractWacsiCommon(ctrl *gomock.Controller) *MockContractWacsiCommon {
	mock := &MockContractWacsiCommon{ctrl: ctrl}
	mock.recorder = &MockContractWacsiCommonMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContractWacsiCommon) EXPECT() *MockContractWacsiCommonMockRecorder {
	return m.recorder
}

// CallContract mocks base method.
func (m *MockContractWacsiCommon) CallContract() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallContract")
	ret0, _ := ret[0].(int32)
	return ret0
}

// CallContract indicates an expected call of CallContract.
func (mr *MockContractWacsiCommonMockRecorder) CallContract() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallContract", reflect.TypeOf((*MockContractWacsiCommon)(nil).CallContract))
}

// ErrorResult mocks base method.
func (m *MockContractWacsiCommon) ErrorResult() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ErrorResult")
	ret0, _ := ret[0].(int32)
	return ret0
}

// ErrorResult indicates an expected call of ErrorResult.
func (mr *MockContractWacsiCommonMockRecorder) ErrorResult() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ErrorResult", reflect.TypeOf((*MockContractWacsiCommon)(nil).ErrorResult))
}

// LogMessage mocks base method.
func (m *MockContractWacsiCommon) LogMessage() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LogMessage")
	ret0, _ := ret[0].(int32)
	return ret0
}

// LogMessage indicates an expected call of LogMessage.
func (mr *MockContractWacsiCommonMockRecorder) LogMessage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogMessage", reflect.TypeOf((*MockContractWacsiCommon)(nil).LogMessage))
}

// SuccessResult mocks base method.
func (m *MockContractWacsiCommon) SuccessResult() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SuccessResult")
	ret0, _ := ret[0].(int32)
	return ret0
}

// SuccessResult indicates an expected call of SuccessResult.
func (mr *MockContractWacsiCommonMockRecorder) SuccessResult() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SuccessResult", reflect.TypeOf((*MockContractWacsiCommon)(nil).SuccessResult))
}

// MockContractWacsiKV is a mock of ContractWacsiKV interface.
type MockContractWacsiKV struct {
	ctrl     *gomock.Controller
	recorder *MockContractWacsiKVMockRecorder
}

// MockContractWacsiKVMockRecorder is the mock recorder for MockContractWacsiKV.
type MockContractWacsiKVMockRecorder struct {
	mock *MockContractWacsiKV
}

// NewMockContractWacsiKV creates a new mock instance.
func NewMockContractWacsiKV(ctrl *gomock.Controller) *MockContractWacsiKV {
	mock := &MockContractWacsiKV{ctrl: ctrl}
	mock.recorder = &MockContractWacsiKVMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContractWacsiKV) EXPECT() *MockContractWacsiKVMockRecorder {
	return m.recorder
}

// CallContract mocks base method.
func (m *MockContractWacsiKV) CallContract() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallContract")
	ret0, _ := ret[0].(int32)
	return ret0
}

// CallContract indicates an expected call of CallContract.
func (mr *MockContractWacsiKVMockRecorder) CallContract() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallContract", reflect.TypeOf((*MockContractWacsiKV)(nil).CallContract))
}

// DeleteState mocks base method.
func (m *MockContractWacsiKV) DeleteState() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteState")
	ret0, _ := ret[0].(int32)
	return ret0
}

// DeleteState indicates an expected call of DeleteState.
func (mr *MockContractWacsiKVMockRecorder) DeleteState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteState", reflect.TypeOf((*MockContractWacsiKV)(nil).DeleteState))
}

// ErrorResult mocks base method.
func (m *MockContractWacsiKV) ErrorResult() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ErrorResult")
	ret0, _ := ret[0].(int32)
	return ret0
}

// ErrorResult indicates an expected call of ErrorResult.
func (mr *MockContractWacsiKVMockRecorder) ErrorResult() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ErrorResult", reflect.TypeOf((*MockContractWacsiKV)(nil).ErrorResult))
}

// GetState mocks base method.
func (m *MockContractWacsiKV) GetState() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetState")
	ret0, _ := ret[0].(int32)
	return ret0
}

// GetState indicates an expected call of GetState.
func (mr *MockContractWacsiKVMockRecorder) GetState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetState", reflect.TypeOf((*MockContractWacsiKV)(nil).GetState))
}

// KvIterator mocks base method.
func (m *MockContractWacsiKV) KvIterator() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KvIterator")
	ret0, _ := ret[0].(int32)
	return ret0
}

// KvIterator indicates an expected call of KvIterator.
func (mr *MockContractWacsiKVMockRecorder) KvIterator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KvIterator", reflect.TypeOf((*MockContractWacsiKV)(nil).KvIterator))
}

// KvIteratorClose mocks base method.
func (m *MockContractWacsiKV) KvIteratorClose() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KvIteratorClose")
	ret0, _ := ret[0].(int32)
	return ret0
}

// KvIteratorClose indicates an expected call of KvIteratorClose.
func (mr *MockContractWacsiKVMockRecorder) KvIteratorClose() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KvIteratorClose", reflect.TypeOf((*MockContractWacsiKV)(nil).KvIteratorClose))
}

// KvIteratorHasNext mocks base method.
func (m *MockContractWacsiKV) KvIteratorHasNext() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KvIteratorHasNext")
	ret0, _ := ret[0].(int32)
	return ret0
}

// KvIteratorHasNext indicates an expected call of KvIteratorHasNext.
func (mr *MockContractWacsiKVMockRecorder) KvIteratorHasNext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KvIteratorHasNext", reflect.TypeOf((*MockContractWacsiKV)(nil).KvIteratorHasNext))
}

// KvIteratorNext mocks base method.
func (m *MockContractWacsiKV) KvIteratorNext() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KvIteratorNext")
	ret0, _ := ret[0].(int32)
	return ret0
}

// KvIteratorNext indicates an expected call of KvIteratorNext.
func (mr *MockContractWacsiKVMockRecorder) KvIteratorNext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KvIteratorNext", reflect.TypeOf((*MockContractWacsiKV)(nil).KvIteratorNext))
}

// LogMessage mocks base method.
func (m *MockContractWacsiKV) LogMessage() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LogMessage")
	ret0, _ := ret[0].(int32)
	return ret0
}

// LogMessage indicates an expected call of LogMessage.
func (mr *MockContractWacsiKVMockRecorder) LogMessage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogMessage", reflect.TypeOf((*MockContractWacsiKV)(nil).LogMessage))
}

// PutState mocks base method.
func (m *MockContractWacsiKV) PutState() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutState")
	ret0, _ := ret[0].(int32)
	return ret0
}

// PutState indicates an expected call of PutState.
func (mr *MockContractWacsiKVMockRecorder) PutState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutState", reflect.TypeOf((*MockContractWacsiKV)(nil).PutState))
}

// SuccessResult mocks base method.
func (m *MockContractWacsiKV) SuccessResult() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SuccessResult")
	ret0, _ := ret[0].(int32)
	return ret0
}

// SuccessResult indicates an expected call of SuccessResult.
func (mr *MockContractWacsiKVMockRecorder) SuccessResult() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SuccessResult", reflect.TypeOf((*MockContractWacsiKV)(nil).SuccessResult))
}

// MockContractWacsiSQL is a mock of ContractWacsiSQL interface.
type MockContractWacsiSQL struct {
	ctrl     *gomock.Controller
	recorder *MockContractWacsiSQLMockRecorder
}

// MockContractWacsiSQLMockRecorder is the mock recorder for MockContractWacsiSQL.
type MockContractWacsiSQLMockRecorder struct {
	mock *MockContractWacsiSQL
}

// NewMockContractWacsiSQL creates a new mock instance.
func NewMockContractWacsiSQL(ctrl *gomock.Controller) *MockContractWacsiSQL {
	mock := &MockContractWacsiSQL{ctrl: ctrl}
	mock.recorder = &MockContractWacsiSQLMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContractWacsiSQL) EXPECT() *MockContractWacsiSQLMockRecorder {
	return m.recorder
}

// CallContract mocks base method.
func (m *MockContractWacsiSQL) CallContract() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallContract")
	ret0, _ := ret[0].(int32)
	return ret0
}

// CallContract indicates an expected call of CallContract.
func (mr *MockContractWacsiSQLMockRecorder) CallContract() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallContract", reflect.TypeOf((*MockContractWacsiSQL)(nil).CallContract))
}

// ErrorResult mocks base method.
func (m *MockContractWacsiSQL) ErrorResult() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ErrorResult")
	ret0, _ := ret[0].(int32)
	return ret0
}

// ErrorResult indicates an expected call of ErrorResult.
func (mr *MockContractWacsiSQLMockRecorder) ErrorResult() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ErrorResult", reflect.TypeOf((*MockContractWacsiSQL)(nil).ErrorResult))
}

// ExecuteDDL mocks base method.
func (m *MockContractWacsiSQL) ExecuteDDL() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecuteDDL")
	ret0, _ := ret[0].(int32)
	return ret0
}

// ExecuteDDL indicates an expected call of ExecuteDDL.
func (mr *MockContractWacsiSQLMockRecorder) ExecuteDDL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteDDL", reflect.TypeOf((*MockContractWacsiSQL)(nil).ExecuteDDL))
}

// ExecuteQuery mocks base method.
func (m *MockContractWacsiSQL) ExecuteQuery() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecuteQuery")
	ret0, _ := ret[0].(int32)
	return ret0
}

// ExecuteQuery indicates an expected call of ExecuteQuery.
func (mr *MockContractWacsiSQLMockRecorder) ExecuteQuery() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteQuery", reflect.TypeOf((*MockContractWacsiSQL)(nil).ExecuteQuery))
}

// ExecuteQueryOne mocks base method.
func (m *MockContractWacsiSQL) ExecuteQueryOne() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecuteQueryOne")
	ret0, _ := ret[0].(int32)
	return ret0
}

// ExecuteQueryOne indicates an expected call of ExecuteQueryOne.
func (mr *MockContractWacsiSQLMockRecorder) ExecuteQueryOne() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteQueryOne", reflect.TypeOf((*MockContractWacsiSQL)(nil).ExecuteQueryOne))
}

// ExecuteUpdate mocks base method.
func (m *MockContractWacsiSQL) ExecuteUpdate() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecuteUpdate")
	ret0, _ := ret[0].(int32)
	return ret0
}

// ExecuteUpdate indicates an expected call of ExecuteUpdate.
func (mr *MockContractWacsiSQLMockRecorder) ExecuteUpdate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteUpdate", reflect.TypeOf((*MockContractWacsiSQL)(nil).ExecuteUpdate))
}

// LogMessage mocks base method.
func (m *MockContractWacsiSQL) LogMessage() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LogMessage")
	ret0, _ := ret[0].(int32)
	return ret0
}

// LogMessage indicates an expected call of LogMessage.
func (mr *MockContractWacsiSQLMockRecorder) LogMessage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogMessage", reflect.TypeOf((*MockContractWacsiSQL)(nil).LogMessage))
}

// RSClose mocks base method.
func (m *MockContractWacsiSQL) RSClose() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RSClose")
	ret0, _ := ret[0].(int32)
	return ret0
}

// RSClose indicates an expected call of RSClose.
func (mr *MockContractWacsiSQLMockRecorder) RSClose() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RSClose", reflect.TypeOf((*MockContractWacsiSQL)(nil).RSClose))
}

// RSHasNext mocks base method.
func (m *MockContractWacsiSQL) RSHasNext() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RSHasNext")
	ret0, _ := ret[0].(int32)
	return ret0
}

// RSHasNext indicates an expected call of RSHasNext.
func (mr *MockContractWacsiSQLMockRecorder) RSHasNext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RSHasNext", reflect.TypeOf((*MockContractWacsiSQL)(nil).RSHasNext))
}

// RSNext mocks base method.
func (m *MockContractWacsiSQL) RSNext() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RSNext")
	ret0, _ := ret[0].(int32)
	return ret0
}

// RSNext indicates an expected call of RSNext.
func (mr *MockContractWacsiSQLMockRecorder) RSNext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RSNext", reflect.TypeOf((*MockContractWacsiSQL)(nil).RSNext))
}

// SuccessResult mocks base method.
func (m *MockContractWacsiSQL) SuccessResult() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SuccessResult")
	ret0, _ := ret[0].(int32)
	return ret0
}

// SuccessResult indicates an expected call of SuccessResult.
func (mr *MockContractWacsiSQLMockRecorder) SuccessResult() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SuccessResult", reflect.TypeOf((*MockContractWacsiSQL)(nil).SuccessResult))
}
