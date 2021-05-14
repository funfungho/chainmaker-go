/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package sync

import (
	commonPb "chainmaker.org/chainmaker-go/pb/protogo/common"
	netPb "chainmaker.org/chainmaker-go/pb/protogo/net"
	storePb "chainmaker.org/chainmaker-go/pb/protogo/store"
	syncPb "chainmaker.org/chainmaker-go/pb/protogo/sync"
	"fmt"

	"chainmaker.org/chainmaker-go/protocol"
)

var errStr = "implement me"

type netMsg struct {
	msgType netPb.NetMsg_MsgType
	bz      []byte
}

type MockNet struct {
	broadcastMsgs []netMsg
	sendMsgs      []string
}

func NewMockNet() *MockNet {
	return &MockNet{broadcastMsgs: make([]netMsg, 0, 100)}
}

func (m MockNet) ChainId() string {
	panic(errStr)
}

func (m *MockNet) BroadcastMsg(msg []byte, msgType netPb.NetMsg_MsgType) error {
	m.broadcastMsgs = append(m.broadcastMsgs, netMsg{msgType: msgType, bz: msg})
	return nil
}

func (m *MockNet) Subscribe(msgType netPb.NetMsg_MsgType, handler protocol.MsgHandler) error {
	return nil
}

func (m MockNet) CancelSubscribe(msgType netPb.NetMsg_MsgType) error {
	panic(errStr)
}

func (m MockNet) ConsensusBroadcastMsg(msg []byte, msgType netPb.NetMsg_MsgType) error {
	panic(errStr)
}

func (m MockNet) ConsensusSubscribe(msgType netPb.NetMsg_MsgType, handler protocol.MsgHandler) error {
	panic(errStr)
}

func (m MockNet) CancelConsensusSubscribe(msgType netPb.NetMsg_MsgType) error {
	panic(errStr)
}

func (m *MockNet) SendMsg(msg []byte, msgType netPb.NetMsg_MsgType, to ...string) error {
	m.sendMsgs = append(m.sendMsgs, fmt.Sprintf("msgType: %d, to: %v", msgType, to))
	return nil
}

func (m MockNet) ReceiveMsg(msgType netPb.NetMsg_MsgType, handler protocol.MsgHandler) error {
	return nil
}

func (m MockNet) Start() error {
	panic(errStr)
}

func (m MockNet) Stop() error {
	panic(errStr)
}

func (m MockNet) GetNodeUidByCertId(certId string) (string, error) {
	panic(errStr)
}

func (m MockNet) GetChainNodesInfoProvider() protocol.ChainNodesInfoProvider {
	panic(errStr)
}

type MockStore struct {
	blocks map[int64]*commonPb.Block
}

func (m MockStore) GetArchivedPivot() uint64 {
	panic("implement me")
}

func (m MockStore) ArchiveBlock(archiveHeight uint64) error {
	panic("implement me")
}

func NewMockStore() *MockStore {
	return &MockStore{blocks: make(map[int64]*commonPb.Block)}
}

func (m MockStore) GetBlockByHash(blockHash []byte) (*commonPb.Block, error) {
	panic(errStr)
}
func (m MockStore) GetTopicTableColumn(tableName string) ([]string, error) {
	panic(errStr)
}
func (m MockStore) BlockExists(blockHash []byte) (bool, error) {
	panic(errStr)

}

func (m MockStore) GetBlock(height int64) (*commonPb.Block, error) {
	if blk, exist := m.blocks[height]; exist {
		return blk, nil
	}
	return nil, fmt.Errorf("block not find")
}

func (m MockStore) GetBlockWithRWSets(height int64) (*storePb.BlockWithRWSet, error) {
	panic(errStr)
}

func (m MockStore) TxExists(txId string) (bool, error) {
	panic(errStr)
}

func (m MockStore) GetTxConfirmedTime(txId string) (int64, error) {
	panic(errStr)
}

func (m *MockStore) PutBlock(block *commonPb.Block, txRWSets []*commonPb.TxRWSet, contractEventInfo []*commonPb.ContractEvent) error {
	m.blocks[block.Header.BlockHeight] = block
	return nil
}

func (m MockStore) GetLastConfigBlock() (*commonPb.Block, error) {
	panic(errStr)
}

func (m MockStore) GetBlockByTx(txId string) (*commonPb.Block, error) {
	panic(errStr)
}

func (m MockStore) GetBlockWithTxRWSets(height int64) (*storePb.BlockWithRWSet, error) {
	panic(errStr)
}

func (m MockStore) GetTx(txId string) (*commonPb.Transaction, error) {
	panic(errStr)
}

func (m MockStore) HasTx(txId string) (bool, error) {
	panic(errStr)
}

func (m MockStore) GetLastBlock() (*commonPb.Block, error) {
	panic(errStr)
}

func (m MockStore) ReadObject(contractName string, key []byte) ([]byte, error) {
	panic(errStr)
}

func (m MockStore) SelectObject(contractName string, startKey []byte, limit []byte) protocol.Iterator {
	panic(errStr)
}

func (m MockStore) GetTxRWSet(txId string) (*commonPb.TxRWSet, error) {
	panic(errStr)
}

func (m MockStore) GetTxRWSetsByHeight(height int64) ([]*commonPb.TxRWSet, error) {
	panic(errStr)
}

func (m MockStore) GetDBHandle(dbName string) protocol.DBHandle {
	panic(errStr)
}

func (m MockStore) Close() error {
	panic(errStr)
}

type MockVerifier struct {
}

func NewMockVerifier() *MockVerifier {
	return &MockVerifier{}
}

func (m MockVerifier) VerifyBlock(block *commonPb.Block, mode protocol.VerifyMode) error {
	return nil
}

func (m MockVerifier) GetLastProposedBlock(b *commonPb.Block) (*commonPb.Block, map[string]*commonPb.TxRWSet) {
	panic(errStr)
}

type MockCommit struct {
	cache protocol.LedgerCache
}

func NewMockCommit(cache protocol.LedgerCache) *MockCommit {
	return &MockCommit{cache: cache}
}

func (m *MockCommit) AddBlock(blk *commonPb.Block) error {
	m.cache.SetLastCommittedBlock(blk)
	return nil
}

type MockSender struct {
	msgs []string
}

func NewMockSender() *MockSender {
	return &MockSender{}
}

func (m MockSender) broadcastMsg(msgType syncPb.SyncMsg_MsgType, msg []byte) error {
	panic(errStr)
}

func (m *MockSender) sendMsg(msgType syncPb.SyncMsg_MsgType, msg []byte, to string) error {
	m.msgs = append(m.msgs, fmt.Sprintf("msgType: %d, to: %s", msgType, to))
	return nil
}

type MockLedgerCache struct {
	blk *commonPb.Block
}

func NewMockLedgerCache(blk *commonPb.Block) *MockLedgerCache {
	return &MockLedgerCache{blk: blk}
}

func (m *MockLedgerCache) GetLastCommittedBlock() *commonPb.Block {
	return m.blk
}

func (m *MockLedgerCache) SetLastCommittedBlock(b *commonPb.Block) {
	m.blk = b
}

func (m *MockLedgerCache) CurrentHeight() (int64, error) {
	return m.blk.Header.BlockHeight, nil
}

type MockVerifyAndCommit struct {
	cache       protocol.LedgerCache
	receiveItem []*commonPb.Block
}

func NewMockVerifyAndCommit(cache protocol.LedgerCache) *MockVerifyAndCommit {
	return &MockVerifyAndCommit{cache: cache}
}

func (m *MockVerifyAndCommit) validateAndCommitBlock(block *commonPb.Block) processedBlockStatus {
	m.receiveItem = append(m.receiveItem, block)
	m.cache.SetLastCommittedBlock(block)
	return ok
}
