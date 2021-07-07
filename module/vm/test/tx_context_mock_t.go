/*
Copyright (C) BABEC. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package test

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"

	"chainmaker.org/chainmaker-go/chainconf"
	"chainmaker.org/chainmaker-go/utils"
	"chainmaker.org/chainmaker/pb-go/config"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"

	"chainmaker.org/chainmaker-go/accesscontrol"
	"chainmaker.org/chainmaker-go/logger"
	"chainmaker.org/chainmaker-go/vm"
	"chainmaker.org/chainmaker-go/wasmer"
	wasm "chainmaker.org/chainmaker-go/wasmer/wasmer-go"
	"chainmaker.org/chainmaker-go/wxvm/xvm"
	acPb "chainmaker.org/chainmaker/pb-go/accesscontrol"
	commonPb "chainmaker.org/chainmaker/pb-go/common"
	storePb "chainmaker.org/chainmaker/pb-go/store"
	"chainmaker.org/chainmaker/protocol"
)

var testOrgId = "wx-org1.chainmaker.org"
var CertFilePath = "../../../../config/crypto-config/wx-org1.chainmaker.org/user/admin1/admin1.sign.crt"
var WasmFile = "../../../../test/wasm/rust-func-verify-1.2.1.wasm"
var isSql = false

var txType = commonPb.TxType_INVOKE_CONTRACT
var pool *wasmer.VmPoolManager

const (
	ContractNameTest    = "contract01"
	ContractVersionTest = "v1.0.0"
	ChainIdTest         = "chain01"
)

func GetVmPoolManager() *wasmer.VmPoolManager {
	if pool == nil {
		pool = wasmer.NewVmPoolManager(ChainIdTest)
	}
	return pool
}

var bytes []byte
var file []byte

// 初始化上下文和wasm字节码
func InitContextTest(runtimeType commonPb.RuntimeType) (*commonPb.ContractId, protocol.TxSimContext, []byte) {
	if bytes == nil {
		bytes, _ = wasm.ReadBytes(WasmFile)
		fmt.Printf("Wasm file size=%d\n", len(bytes))
	}

	contractId := commonPb.ContractId{
		ContractName:    ContractNameTest,
		ContractVersion: ContractVersionTest,
		RuntimeType:     runtimeType,
	}

	wxvmCodeManager := xvm.NewCodeManager(ChainIdTest, "tmp/wxvm-data")
	wxvmContextService := xvm.NewContextService(ChainIdTest)
	log := logger.GetLoggerByChain(logger.MODULE_VM, ChainIdTest)

	if file == nil {
		var err error
		file, err = ioutil.ReadFile(CertFilePath)
		if err != nil {
			panic("file is nil" + err.Error())
		}
	}
	sender := &acPb.SerializedMember{
		OrgId:      testOrgId,
		MemberInfo: file,
		//IsFullCert: true,
	}

	db, _ := leveldb.OpenFile("tmp/leveldb"+utils.GetRandTxId(), nil)

	chainConf := &chainconf.ChainConf{
		ChainConf: &config.ChainConfig{
			Contract: &config.ContractConfig{EnableSqlSupport: isSql},
		},
	}

	txContext := TxContextMockTest{
		lock:  &sync.Mutex{},
		lock2: &sync.Mutex{},
		vmManager: &vm.ManagerImpl{
			WasmerVmPoolManager:    GetVmPoolManager(),
			WxvmCodeManager:        wxvmCodeManager,
			WxvmContextService:     wxvmContextService,
			SnapshotManager:        nil,
			AccessControl:          accesscontrol.MockAccessControl(),
			ChainNodesInfoProvider: nil,
			ChainId:                ChainIdTest,
			Log:                    log,
			ChainConf:              chainConf,
		},
		hisResult:  make([]*callContractResult, 0),
		creator:    sender,
		sender:     sender,
		cacheMap:   make(map[string][]byte),
		db:         db,
		kvRowCache: make(map[int32]protocol.StateIterator),
	}

	versionKey := []byte(protocol.ContractVersion + ContractNameTest)
	runtimeTypeKey := []byte(protocol.ContractRuntimeType + ContractNameTest)
	versionedByteCodeKey := append([]byte(protocol.ContractByteCode+ContractNameTest), []byte(contractId.ContractVersion)...)

	txContext.Put(commonPb.ContractName_SYSTEM_CONTRACT_USER_CONTRACT_MANAGE.String(), versionedByteCodeKey, bytes)
	txContext.Put(commonPb.ContractName_SYSTEM_CONTRACT_USER_CONTRACT_MANAGE.String(), versionKey, []byte(contractId.ContractVersion))
	txContext.Put(commonPb.ContractName_SYSTEM_CONTRACT_USER_CONTRACT_MANAGE.String(), runtimeTypeKey, []byte(strconv.Itoa(int(runtimeType))))

	return &contractId, &txContext, bytes
}

// test
// test
// test
// test

type TxContextMockTest struct {
	lock          *sync.Mutex
	lock2         *sync.Mutex
	vmManager     protocol.VmManager
	gasUsed       uint64 // only for callContract
	currentDepth  int
	currentResult []byte
	hisResult     []*callContractResult

	sender     *acPb.SerializedMember
	creator    *acPb.SerializedMember
	cacheMap   map[string][]byte
	db         *leveldb.DB
	kvRowCache map[int32]protocol.StateIterator
}

func (s *TxContextMockTest) PutRecord(contractName string, value []byte, sqlType protocol.SqlType) {
	panic("implement me")
}
func (s *TxContextMockTest) SetStateKvHandle(index int32, rows protocol.StateIterator) {
	s.lock2.Lock()
	defer s.lock2.Unlock()
	s.kvRowCache[index] = rows
}

func (s *TxContextMockTest) GetStateKvHandle(index int32) (protocol.StateIterator, bool) {
	s.lock2.Lock()
	defer s.lock2.Unlock()
	data, ok := s.kvRowCache[index]
	return data, ok
}

func (s *TxContextMockTest) GetBlockProposer() []byte {
	panic("implement me")
}

func (s *TxContextMockTest) SetStateSqlHandle(i int32, rows protocol.SqlRows) {
	panic("implement me")
}

func (s *TxContextMockTest) GetStateSqlHandle(i int32) (protocol.SqlRows, bool) {
	panic("implement me")
}

type callContractResult struct {
	contractName string
	method       string
	param        map[string]string
	deep         int
	gasUsed      uint64
	result       []byte
}

func constructStateKey(contractName string, key []byte) []byte {
	return append(append([]byte(contractName), contractStoreSeparator), key...)
}
func (s *TxContextMockTest) Get(name string, key []byte) ([]byte, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	key = constructStateKey(name, key)
	k := string(key)
	val, err := s.db.Get([]byte(k), nil)
	if err != nil {
		fmt.Println("get", err)
	}
	return val, nil
}

const contractStoreSeparator = '#'

func (s *TxContextMockTest) Put(name string, key []byte, value []byte) error {
	key = constructStateKey(name, key)
	k := string(key)
	if !strings.Contains(k, "SYSTEM_CONTRACT_STATE#:B:") {
		//fmt.Println("【put】 key:", k, "val:", string(value))
	}
	wo := &opt.WriteOptions{Sync: true}
	s.db.Put([]byte(k), value, wo)
	//s.cacheMap[k] = value
	return nil
}

func (s *TxContextMockTest) Select(name string, startKey []byte, limit []byte) (protocol.StateIterator, error) {
	startKey = constructStateKey(name, startKey)
	limit = constructStateKey(name, limit)
	fmt.Println("select ", string(startKey), string(limit))
	keyRange := &util.Range{Start: startKey, Limit: limit}
	iter := s.db.NewIterator(keyRange, nil)
	return &kvi{
		iter:         iter,
		contractName: name,
	}, nil
}

func (s *TxContextMockTest) Del(name string, key []byte) error {
	key = constructStateKey(name, key)
	k := string(key)
	if !strings.Contains(k, "SYSTEM_CONTRACT_STATE#:B:") {
		//fmt.Println("【put】 key:", k, "val:", string(value))
	}
	wo := &opt.WriteOptions{Sync: true}
	s.db.Put([]byte(k), nil, wo)
	//s.cacheMap[k] = value
	return nil
}

func (s *TxContextMockTest) CallContract(contract *commonPb.Contract, method string, byteCode []byte,
	parameter map[string]string, gasUsed uint64, refTxType commonPb.TxType) (*commonPb.ContractResult, commonPb.TxStatusCode) {
	s.gasUsed = gasUsed
	s.currentDepth = s.currentDepth + 1
	if s.currentDepth > protocol.CallContractDepth {
		contractResult := &commonPb.ContractResult{
			Code:    commonPb.ContractResultCode_FAIL,
			Result:  nil,
			Message: fmt.Sprintf("CallContract too deep %d", s.currentDepth),
		}
		return contractResult, commonPb.TxStatusCode_CONTRACT_TOO_DEEP_FAILED
	}
	if s.gasUsed > protocol.GasLimit {
		contractResult := &commonPb.ContractResult{
			Code:    commonPb.ContractResultCode_FAIL,
			Result:  nil,
			Message: fmt.Sprintf("There is not enough gas, gasUsed %d GasLimit %d ", gasUsed, int64(protocol.GasLimit)),
		}
		return contractResult, commonPb.TxStatusCode_CONTRACT_FAIL
	}
	r, code := s.vmManager.RunContract(contract, method, byteCode, parameter, s, s.gasUsed, refTxType)

	result := callContractResult{
		deep:         s.currentDepth,
		gasUsed:      s.gasUsed,
		result:       r.Result,
		contractName: contract.Name,
		method:       method,
		param:        parameter,
	}
	s.hisResult = append(s.hisResult, &result)
	s.currentResult = r.Result
	s.currentDepth = s.currentDepth - 1
	return r, code
}

func (s *TxContextMockTest) GetCurrentResult() []byte {
	return s.currentResult
}

func (s *TxContextMockTest) GetTx() *commonPb.Transaction {
	return &commonPb.Transaction{
		Header: &commonPb.TxHeader{
			ChainId:        ChainIdTest,
			Sender:         s.GetSender(),
			TxType:         txType,
			TxId:           "abcdef12345678",
			Timestamp:      0,
			ExpirationTime: 0,
		},
		RequestPayload:   nil,
		RequestSignature: nil,
		Result:           nil,
	}
}

func (*TxContextMockTest) GetBlockHeight() int64 {
	return 0
}
func (s *TxContextMockTest) GetTxResult() *commonPb.Result {
	panic("implement me")
}

func (s *TxContextMockTest) SetTxResult(txResult *commonPb.Result) {
	panic("implement me")
}

func (TxContextMockTest) GetTxRWSet(runVmSuccess bool) *commonPb.TxRWSet {
	return &commonPb.TxRWSet{
		TxId:     "txId",
		TxReads:  nil,
		TxWrites: nil,
	}
}

func (s *TxContextMockTest) GetCreator(namespace string) *acPb.SerializedMember {
	return s.creator
}

func (s *TxContextMockTest) GetSender() *acPb.SerializedMember {
	return s.sender
}

func (*TxContextMockTest) GetBlockchainStore() protocol.BlockchainStore {
	return &mockBlockchainStore{}
}

func (*TxContextMockTest) GetAccessControl() (protocol.AccessControlProvider, error) {
	panic("implement me")
}

func (s *TxContextMockTest) GetChainNodesInfoProvider() (protocol.ChainNodesInfoProvider, error) {
	panic("implement me")
}

func (*TxContextMockTest) GetTxExecSeq() int {
	panic("implement me")
}

func (*TxContextMockTest) SetTxExecSeq(i int) {
	panic("implement me")
}

func (s *TxContextMockTest) GetDepth() int {
	return s.currentDepth
}

func BaseParam(parameters map[string]string) {
	parameters[protocol.ContractTxIdParam] = "TX_ID"
	parameters[protocol.ContractCreatorOrgIdParam] = "CREATOR_ORG_ID"
	parameters[protocol.ContractCreatorRoleParam] = "CREATOR_ROLE"
	parameters[protocol.ContractCreatorPkParam] = "CREATOR_PK"
	parameters[protocol.ContractSenderOrgIdParam] = "SENDER_ORG_ID"
	parameters[protocol.ContractSenderRoleParam] = "SENDER_ROLE"
	parameters[protocol.ContractSenderPkParam] = "SENDER_PK"
	parameters[protocol.ContractBlockHeightParam] = "111"
}

type mockBlockchainStore struct {
}

func (m mockBlockchainStore) GetContractByName(name string) (*commonPb.Contract, error) {
	panic("implement me")
}

func (m mockBlockchainStore) GetContractBytecode(name string) ([]byte, error) {
	panic("implement me")
}

func (m mockBlockchainStore) GetHeightByHash(blockHash []byte) (uint64, error) {
	panic("implement me")
}

func (m mockBlockchainStore) GetBlockHeaderByHeight(height int64) (*commonPb.BlockHeader, error) {
	panic("implement me")
}

func (m mockBlockchainStore) GetTxHeight(txId string) (uint64, error) {
	panic("implement me")
}

func (m mockBlockchainStore) GetArchivedPivot() uint64 {
	panic("implement me")
}

func (m mockBlockchainStore) ArchiveBlock(archiveHeight uint64) error {
	panic("implement me")
}

func (m mockBlockchainStore) RestoreBlocks(serializedBlocks [][]byte) error {
	panic("implement me")
}

func (m mockBlockchainStore) GetLastChainConfig() (*config.ChainConfig, error) {
	panic("implement me")
}

func (m mockBlockchainStore) SelectObject(contractName string, startKey []byte, limit []byte) (protocol.StateIterator, error) {
	panic("implement me")
}

func (m mockBlockchainStore) QuerySingle(contractName, sql string, values ...interface{}) (protocol.SqlRow, error) {
	panic("implement me")
}

func (m mockBlockchainStore) QueryMulti(contractName, sql string, values ...interface{}) (protocol.SqlRows, error) {
	panic("implement me")
}

func (m mockBlockchainStore) ExecDdlSql(contractName, sql string) error {
	panic("implement me")
}

func (m mockBlockchainStore) BeginDbTransaction(txName string) (protocol.SqlDBTransaction, error) {
	panic("implement me")
}

func (m mockBlockchainStore) GetDbTransaction(txName string) (protocol.SqlDBTransaction, error) {
	panic("implement me")
}

func (m mockBlockchainStore) CommitDbTransaction(txName string) error {
	panic("implement me")
}

func (m mockBlockchainStore) RollbackDbTransaction(txName string) error {
	panic("implement me")
}

func (m mockBlockchainStore) InitGenesis(genesisBlock *storePb.BlockWithRWSet) error {
	panic("implement me")
}

func (m mockBlockchainStore) PutBlock(block *commonPb.Block, txRWSets []*commonPb.TxRWSet) error {
	panic("implement me")
}

func (m mockBlockchainStore) GetHistoryForKey(contractName string, key []byte) (protocol.KeyHistoryIterator, error) {
	panic("implement me")
}

func (m mockBlockchainStore) GetAccountTxHistory(accountId []byte) (protocol.TxHistoryIterator, error) {
	panic("implement me")
}

func (m mockBlockchainStore) GetContractTxHistory(contractName string) (protocol.TxHistoryIterator, error) {
	panic("implement me")
}

func (m mockBlockchainStore) GetBlockByHash(blockHash []byte) (*commonPb.Block, error) {
	panic("implement me")
}

func (m mockBlockchainStore) BlockExists(blockHash []byte) (bool, error) {
	panic("implement me")
}

func (m mockBlockchainStore) GetBlock(height int64) (*commonPb.Block, error) {
	panic("implement me")
}

func (m mockBlockchainStore) GetLastConfigBlock() (*commonPb.Block, error) {
	panic("implement me")
}

func (m mockBlockchainStore) GetBlockByTx(txId string) (*commonPb.Block, error) {
	panic("implement me")
}

func (m mockBlockchainStore) GetBlockWithRWSets(height int64) (*storePb.BlockWithRWSet, error) {
	panic("implement me")
}

func (m mockBlockchainStore) GetTx(txId string) (*commonPb.Transaction, error) {
	panic("implement me")
}

func (m mockBlockchainStore) TxExists(txId string) (bool, error) {
	panic("implement me")
}

func (m mockBlockchainStore) GetTxConfirmedTime(txId string) (int64, error) {
	panic("implement me")
}

func (m mockBlockchainStore) GetLastBlock() (*commonPb.Block, error) {
	return &commonPb.Block{
		Header: &commonPb.BlockHeader{
			ChainId:        "",
			BlockHeight:    0,
			PreBlockHash:   nil,
			BlockHash:      nil,
			PreConfHeight:  0,
			BlockVersion:   nil,
			DagHash:        nil,
			RwSetRoot:      nil,
			TxRoot:         nil,
			BlockTimestamp: 0,
			Proposer:       nil,
			ConsensusArgs:  nil,
			TxCount:        0,
			Signature:      nil,
		},
		Dag:            nil,
		Txs:            nil,
		AdditionalData: nil,
	}, nil
}

func (m mockBlockchainStore) ReadObject(contractName string, key []byte) ([]byte, error) {
	panic("implement me")
}

func (m mockBlockchainStore) GetTxRWSet(txId string) (*commonPb.TxRWSet, error) {
	panic("implement me")
}

func (m mockBlockchainStore) GetTxRWSetsByHeight(height int64) ([]*commonPb.TxRWSet, error) {
	panic("implement me")
}

func (m mockBlockchainStore) GetDBHandle(dbName string) protocol.DBHandle {
	panic("implement me")
}

func (m mockBlockchainStore) Close() error {
	panic("implement me")
}

type kvi struct {
	iter         protocol.Iterator
	contractName string
}

func (i *kvi) Next() bool {
	return i.iter.Next()
}
func (i *kvi) Value() (*storePb.KV, error) {
	err := i.iter.Error()
	if err != nil {
		return nil, err
	}
	return &storePb.KV{
		ContractName: i.contractName,
		Key:          parseStateKey(i.iter.Key(), i.contractName),
		Value:        i.iter.Value(),
	}, nil
}

func (i *kvi) Release() {
	i.iter.Release()
}

// parseStateKey corresponding to the constructStateKey(),  delete contract name from leveldb key
func parseStateKey(key []byte, contractName string) []byte {
	return key[len(contractName)+1:]
}
