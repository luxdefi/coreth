// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package contract is a generated GoMock package.
package contract

import (
	big "math/big"
	reflect "reflect"

	snow "github.com/luxdefi/node/snow"
	precompileconfig "github.com/luxdefi/coreth/precompile/precompileconfig"
	common "github.com/ethereum/go-ethereum/common"
	gomock "go.uber.org/mock/gomock"
)

// MockStatefulPrecompiledContract is a mock of StatefulPrecompiledContract interface.
type MockStatefulPrecompiledContract struct {
	ctrl     *gomock.Controller
	recorder *MockStatefulPrecompiledContractMockRecorder
}

// MockStatefulPrecompiledContractMockRecorder is the mock recorder for MockStatefulPrecompiledContract.
type MockStatefulPrecompiledContractMockRecorder struct {
	mock *MockStatefulPrecompiledContract
}

// NewMockStatefulPrecompiledContract creates a new mock instance.
func NewMockStatefulPrecompiledContract(ctrl *gomock.Controller) *MockStatefulPrecompiledContract {
	mock := &MockStatefulPrecompiledContract{ctrl: ctrl}
	mock.recorder = &MockStatefulPrecompiledContractMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStatefulPrecompiledContract) EXPECT() *MockStatefulPrecompiledContractMockRecorder {
	return m.recorder
}

// Run mocks base method.
func (m *MockStatefulPrecompiledContract) Run(accessibleState AccessibleState, caller, addr common.Address, input []byte, suppliedGas uint64, readOnly bool) ([]byte, uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", accessibleState, caller, addr, input, suppliedGas, readOnly)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Run indicates an expected call of Run.
func (mr *MockStatefulPrecompiledContractMockRecorder) Run(accessibleState, caller, addr, input, suppliedGas, readOnly interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockStatefulPrecompiledContract)(nil).Run), accessibleState, caller, addr, input, suppliedGas, readOnly)
}

// MockStateDB is a mock of StateDB interface.
type MockStateDB struct {
	ctrl     *gomock.Controller
	recorder *MockStateDBMockRecorder
}

// MockStateDBMockRecorder is the mock recorder for MockStateDB.
type MockStateDBMockRecorder struct {
	mock *MockStateDB
}

// NewMockStateDB creates a new mock instance.
func NewMockStateDB(ctrl *gomock.Controller) *MockStateDB {
	mock := &MockStateDB{ctrl: ctrl}
	mock.recorder = &MockStateDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStateDB) EXPECT() *MockStateDBMockRecorder {
	return m.recorder
}

// AddBalance mocks base method.
func (m *MockStateDB) AddBalance(arg0 common.Address, arg1 *big.Int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddBalance", arg0, arg1)
}

// AddBalance indicates an expected call of AddBalance.
func (mr *MockStateDBMockRecorder) AddBalance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBalance", reflect.TypeOf((*MockStateDB)(nil).AddBalance), arg0, arg1)
}

// AddLog mocks base method.
func (m *MockStateDB) AddLog(addr common.Address, topics []common.Hash, data []byte, blockNumber uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddLog", addr, topics, data, blockNumber)
}

// AddLog indicates an expected call of AddLog.
func (mr *MockStateDBMockRecorder) AddLog(addr, topics, data, blockNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddLog", reflect.TypeOf((*MockStateDB)(nil).AddLog), addr, topics, data, blockNumber)
}

// CreateAccount mocks base method.
func (m *MockStateDB) CreateAccount(arg0 common.Address) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateAccount", arg0)
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockStateDBMockRecorder) CreateAccount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockStateDB)(nil).CreateAccount), arg0)
}

// Exist mocks base method.
func (m *MockStateDB) Exist(arg0 common.Address) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exist", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Exist indicates an expected call of Exist.
func (mr *MockStateDBMockRecorder) Exist(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exist", reflect.TypeOf((*MockStateDB)(nil).Exist), arg0)
}

// Finalise mocks base method.
func (m *MockStateDB) Finalise(deleteEmptyObjects bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Finalise", deleteEmptyObjects)
}

// Finalise indicates an expected call of Finalise.
func (mr *MockStateDBMockRecorder) Finalise(deleteEmptyObjects interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Finalise", reflect.TypeOf((*MockStateDB)(nil).Finalise), deleteEmptyObjects)
}

// GetBalance mocks base method.
func (m *MockStateDB) GetBalance(arg0 common.Address) *big.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", arg0)
	ret0, _ := ret[0].(*big.Int)
	return ret0
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockStateDBMockRecorder) GetBalance(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockStateDB)(nil).GetBalance), arg0)
}

// GetBalanceMultiCoin mocks base method.
func (m *MockStateDB) GetBalanceMultiCoin(arg0 common.Address, arg1 common.Hash) *big.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalanceMultiCoin", arg0, arg1)
	ret0, _ := ret[0].(*big.Int)
	return ret0
}

// GetBalanceMultiCoin indicates an expected call of GetBalanceMultiCoin.
func (mr *MockStateDBMockRecorder) GetBalanceMultiCoin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalanceMultiCoin", reflect.TypeOf((*MockStateDB)(nil).GetBalanceMultiCoin), arg0, arg1)
}

// GetLogData mocks base method.
func (m *MockStateDB) GetLogData() [][]byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogData")
	ret0, _ := ret[0].([][]byte)
	return ret0
}

// GetLogData indicates an expected call of GetLogData.
func (mr *MockStateDBMockRecorder) GetLogData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogData", reflect.TypeOf((*MockStateDB)(nil).GetLogData))
}

// GetNonce mocks base method.
func (m *MockStateDB) GetNonce(arg0 common.Address) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNonce", arg0)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetNonce indicates an expected call of GetNonce.
func (mr *MockStateDBMockRecorder) GetNonce(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNonce", reflect.TypeOf((*MockStateDB)(nil).GetNonce), arg0)
}

// GetPredicateStorageSlots mocks base method.
func (m *MockStateDB) GetPredicateStorageSlots(address common.Address, index int) ([]byte, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPredicateStorageSlots", address, index)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetPredicateStorageSlots indicates an expected call of GetPredicateStorageSlots.
func (mr *MockStateDBMockRecorder) GetPredicateStorageSlots(address, index interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPredicateStorageSlots", reflect.TypeOf((*MockStateDB)(nil).GetPredicateStorageSlots), address, index)
}

// GetState mocks base method.
func (m *MockStateDB) GetState(arg0 common.Address, arg1 common.Hash) common.Hash {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetState", arg0, arg1)
	ret0, _ := ret[0].(common.Hash)
	return ret0
}

// GetState indicates an expected call of GetState.
func (mr *MockStateDBMockRecorder) GetState(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetState", reflect.TypeOf((*MockStateDB)(nil).GetState), arg0, arg1)
}

// GetTxHash mocks base method.
func (m *MockStateDB) GetTxHash() common.Hash {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTxHash")
	ret0, _ := ret[0].(common.Hash)
	return ret0
}

// GetTxHash indicates an expected call of GetTxHash.
func (mr *MockStateDBMockRecorder) GetTxHash() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTxHash", reflect.TypeOf((*MockStateDB)(nil).GetTxHash))
}

// RevertToSnapshot mocks base method.
func (m *MockStateDB) RevertToSnapshot(arg0 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RevertToSnapshot", arg0)
}

// RevertToSnapshot indicates an expected call of RevertToSnapshot.
func (mr *MockStateDBMockRecorder) RevertToSnapshot(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevertToSnapshot", reflect.TypeOf((*MockStateDB)(nil).RevertToSnapshot), arg0)
}

// SetNonce mocks base method.
func (m *MockStateDB) SetNonce(arg0 common.Address, arg1 uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetNonce", arg0, arg1)
}

// SetNonce indicates an expected call of SetNonce.
func (mr *MockStateDBMockRecorder) SetNonce(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNonce", reflect.TypeOf((*MockStateDB)(nil).SetNonce), arg0, arg1)
}

// SetPredicateStorageSlots mocks base method.
func (m *MockStateDB) SetPredicateStorageSlots(address common.Address, predicates [][]byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetPredicateStorageSlots", address, predicates)
}

// SetPredicateStorageSlots indicates an expected call of SetPredicateStorageSlots.
func (mr *MockStateDBMockRecorder) SetPredicateStorageSlots(address, predicates interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPredicateStorageSlots", reflect.TypeOf((*MockStateDB)(nil).SetPredicateStorageSlots), address, predicates)
}

// SetState mocks base method.
func (m *MockStateDB) SetState(arg0 common.Address, arg1, arg2 common.Hash) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetState", arg0, arg1, arg2)
}

// SetState indicates an expected call of SetState.
func (mr *MockStateDBMockRecorder) SetState(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetState", reflect.TypeOf((*MockStateDB)(nil).SetState), arg0, arg1, arg2)
}

// Snapshot mocks base method.
func (m *MockStateDB) Snapshot() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Snapshot")
	ret0, _ := ret[0].(int)
	return ret0
}

// Snapshot indicates an expected call of Snapshot.
func (mr *MockStateDBMockRecorder) Snapshot() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Snapshot", reflect.TypeOf((*MockStateDB)(nil).Snapshot))
}

// Suicide mocks base method.
func (m *MockStateDB) Suicide(arg0 common.Address) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Suicide", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Suicide indicates an expected call of Suicide.
func (mr *MockStateDBMockRecorder) Suicide(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Suicide", reflect.TypeOf((*MockStateDB)(nil).Suicide), arg0)
}

// MockAccessibleState is a mock of AccessibleState interface.
type MockAccessibleState struct {
	ctrl     *gomock.Controller
	recorder *MockAccessibleStateMockRecorder
}

// MockAccessibleStateMockRecorder is the mock recorder for MockAccessibleState.
type MockAccessibleStateMockRecorder struct {
	mock *MockAccessibleState
}

// NewMockAccessibleState creates a new mock instance.
func NewMockAccessibleState(ctrl *gomock.Controller) *MockAccessibleState {
	mock := &MockAccessibleState{ctrl: ctrl}
	mock.recorder = &MockAccessibleStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessibleState) EXPECT() *MockAccessibleStateMockRecorder {
	return m.recorder
}

// GetBlockContext mocks base method.
func (m *MockAccessibleState) GetBlockContext() BlockContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockContext")
	ret0, _ := ret[0].(BlockContext)
	return ret0
}

// GetBlockContext indicates an expected call of GetBlockContext.
func (mr *MockAccessibleStateMockRecorder) GetBlockContext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockContext", reflect.TypeOf((*MockAccessibleState)(nil).GetBlockContext))
}

// GetChainConfig mocks base method.
func (m *MockAccessibleState) GetChainConfig() precompileconfig.ChainConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChainConfig")
	ret0, _ := ret[0].(precompileconfig.ChainConfig)
	return ret0
}

// GetChainConfig indicates an expected call of GetChainConfig.
func (mr *MockAccessibleStateMockRecorder) GetChainConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChainConfig", reflect.TypeOf((*MockAccessibleState)(nil).GetChainConfig))
}

// GetSnowContext mocks base method.
func (m *MockAccessibleState) GetSnowContext() *snow.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSnowContext")
	ret0, _ := ret[0].(*snow.Context)
	return ret0
}

// GetSnowContext indicates an expected call of GetSnowContext.
func (mr *MockAccessibleStateMockRecorder) GetSnowContext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSnowContext", reflect.TypeOf((*MockAccessibleState)(nil).GetSnowContext))
}

// GetStateDB mocks base method.
func (m *MockAccessibleState) GetStateDB() StateDB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStateDB")
	ret0, _ := ret[0].(StateDB)
	return ret0
}

// GetStateDB indicates an expected call of GetStateDB.
func (mr *MockAccessibleStateMockRecorder) GetStateDB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStateDB", reflect.TypeOf((*MockAccessibleState)(nil).GetStateDB))
}

// NativeAssetCall mocks base method.
func (m *MockAccessibleState) NativeAssetCall(caller common.Address, input []byte, suppliedGas, gasCost uint64, readOnly bool) ([]byte, uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NativeAssetCall", caller, input, suppliedGas, gasCost, readOnly)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// NativeAssetCall indicates an expected call of NativeAssetCall.
func (mr *MockAccessibleStateMockRecorder) NativeAssetCall(caller, input, suppliedGas, gasCost, readOnly interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NativeAssetCall", reflect.TypeOf((*MockAccessibleState)(nil).NativeAssetCall), caller, input, suppliedGas, gasCost, readOnly)
}

// MockConfigurationBlockContext is a mock of ConfigurationBlockContext interface.
type MockConfigurationBlockContext struct {
	ctrl     *gomock.Controller
	recorder *MockConfigurationBlockContextMockRecorder
}

// MockConfigurationBlockContextMockRecorder is the mock recorder for MockConfigurationBlockContext.
type MockConfigurationBlockContextMockRecorder struct {
	mock *MockConfigurationBlockContext
}

// NewMockConfigurationBlockContext creates a new mock instance.
func NewMockConfigurationBlockContext(ctrl *gomock.Controller) *MockConfigurationBlockContext {
	mock := &MockConfigurationBlockContext{ctrl: ctrl}
	mock.recorder = &MockConfigurationBlockContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfigurationBlockContext) EXPECT() *MockConfigurationBlockContextMockRecorder {
	return m.recorder
}

// Number mocks base method.
func (m *MockConfigurationBlockContext) Number() *big.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Number")
	ret0, _ := ret[0].(*big.Int)
	return ret0
}

// Number indicates an expected call of Number.
func (mr *MockConfigurationBlockContextMockRecorder) Number() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Number", reflect.TypeOf((*MockConfigurationBlockContext)(nil).Number))
}

// Timestamp mocks base method.
func (m *MockConfigurationBlockContext) Timestamp() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Timestamp")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// Timestamp indicates an expected call of Timestamp.
func (mr *MockConfigurationBlockContextMockRecorder) Timestamp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Timestamp", reflect.TypeOf((*MockConfigurationBlockContext)(nil).Timestamp))
}

// MockBlockContext is a mock of BlockContext interface.
type MockBlockContext struct {
	ctrl     *gomock.Controller
	recorder *MockBlockContextMockRecorder
}

// MockBlockContextMockRecorder is the mock recorder for MockBlockContext.
type MockBlockContextMockRecorder struct {
	mock *MockBlockContext
}

// NewMockBlockContext creates a new mock instance.
func NewMockBlockContext(ctrl *gomock.Controller) *MockBlockContext {
	mock := &MockBlockContext{ctrl: ctrl}
	mock.recorder = &MockBlockContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlockContext) EXPECT() *MockBlockContextMockRecorder {
	return m.recorder
}

// GetPredicateResults mocks base method.
func (m *MockBlockContext) GetPredicateResults(txHash common.Hash, precompileAddress common.Address) []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPredicateResults", txHash, precompileAddress)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// GetPredicateResults indicates an expected call of GetPredicateResults.
func (mr *MockBlockContextMockRecorder) GetPredicateResults(txHash, precompileAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPredicateResults", reflect.TypeOf((*MockBlockContext)(nil).GetPredicateResults), txHash, precompileAddress)
}

// Number mocks base method.
func (m *MockBlockContext) Number() *big.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Number")
	ret0, _ := ret[0].(*big.Int)
	return ret0
}

// Number indicates an expected call of Number.
func (mr *MockBlockContextMockRecorder) Number() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Number", reflect.TypeOf((*MockBlockContext)(nil).Number))
}

// Timestamp mocks base method.
func (m *MockBlockContext) Timestamp() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Timestamp")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// Timestamp indicates an expected call of Timestamp.
func (mr *MockBlockContextMockRecorder) Timestamp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Timestamp", reflect.TypeOf((*MockBlockContext)(nil).Timestamp))
}

// MockConfigurator is a mock of Configurator interface.
type MockConfigurator struct {
	ctrl     *gomock.Controller
	recorder *MockConfiguratorMockRecorder
}

// MockConfiguratorMockRecorder is the mock recorder for MockConfigurator.
type MockConfiguratorMockRecorder struct {
	mock *MockConfigurator
}

// NewMockConfigurator creates a new mock instance.
func NewMockConfigurator(ctrl *gomock.Controller) *MockConfigurator {
	mock := &MockConfigurator{ctrl: ctrl}
	mock.recorder = &MockConfiguratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfigurator) EXPECT() *MockConfiguratorMockRecorder {
	return m.recorder
}

// Configure mocks base method.
func (m *MockConfigurator) Configure(chainConfig precompileconfig.ChainConfig, precompileconfig precompileconfig.Config, state StateDB, blockContext ConfigurationBlockContext) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Configure", chainConfig, precompileconfig, state, blockContext)
	ret0, _ := ret[0].(error)
	return ret0
}

// Configure indicates an expected call of Configure.
func (mr *MockConfiguratorMockRecorder) Configure(chainConfig, precompileconfig, state, blockContext interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Configure", reflect.TypeOf((*MockConfigurator)(nil).Configure), chainConfig, precompileconfig, state, blockContext)
}

// MakeConfig mocks base method.
func (m *MockConfigurator) MakeConfig() precompileconfig.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeConfig")
	ret0, _ := ret[0].(precompileconfig.Config)
	return ret0
}

// MakeConfig indicates an expected call of MakeConfig.
func (mr *MockConfiguratorMockRecorder) MakeConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeConfig", reflect.TypeOf((*MockConfigurator)(nil).MakeConfig))
}
