package main

import (
	"chainmaker.org/chainmaker-go/consensus"
	"chainmaker.org/chainmaker-go/consensus/chainedbft"
	"chainmaker.org/chainmaker-go/consensus/dpos"
	"chainmaker.org/chainmaker-go/consensus/implconfig"
	"chainmaker.org/chainmaker-go/consensus/raft"
	"chainmaker.org/chainmaker-go/consensus/solo"
	"chainmaker.org/chainmaker-go/consensus/tbft"
	"chainmaker.org/chainmaker-go/txpool"
	"chainmaker.org/chainmaker-go/vm"
	consensusPb "chainmaker.org/chainmaker/pb-go/v2/consensus"
	"chainmaker.org/chainmaker/protocol/v2"
	batch "chainmaker.org/chainmaker/txpool-batch/v2"
	single "chainmaker.org/chainmaker/txpool-single/v2"
	evm "chainmaker.org/chainmaker/vm-evm"
	gasm "chainmaker.org/chainmaker/vm-gasm"
	wasmer "chainmaker.org/chainmaker/vm-wasmer"
	wxvm "chainmaker.org/chainmaker/vm-wxvm"
)

func init() {
	// txPool
	txpool.RegisterTxPoolProvider(single.TxPoolType, single.NewTxPoolImpl)
	txpool.RegisterTxPoolProvider(batch.TxPoolType, batch.NewBatchTxPool)

	// vm
	vm.RegisterVmProvider(
		"GASM",
		func(chainId string) (protocol.VmInstancesManager, error) {
			return &gasm.InstancesManager{}, nil
		},
	)

	vm.RegisterVmProvider(
		"WASMER",
		func(chainId string) (protocol.VmInstancesManager, error) {
			return wasmer.NewInstancesManager(chainId), nil
		},
	)

	vm.RegisterVmProvider(
		"WXVM",
		func(chainId string) (protocol.VmInstancesManager, error) {
			return &wxvm.InstancesManager{}, nil
		},
	)

	vm.RegisterVmProvider(
		"EVM",
		func(chainId string) (protocol.VmInstancesManager, error) {
			return &evm.InstancesManager{}, nil
		},
	)

	// consensus
	consensus.RegisterConsensusProvider(
		consensusPb.ConsensusType_SOLO,
		func(config *implconfig.ConsensusImplConfig) (protocol.ConsensusEngine, error) {
			return solo.New(config)
		},
	)

	consensus.RegisterConsensusProvider(
		consensusPb.ConsensusType_DPOS,
		func(config *implconfig.ConsensusImplConfig) (protocol.ConsensusEngine, error) {
			dposEngine:= dpos.NewDPoSImpl(config)
			tbftEngine, err := tbft.New(config, dposEngine)
			if err != nil {
				return nil, err
			}
			dposEngine.SetConsensusEngine(tbftEngine)
			return tbftEngine, nil
		},
	)

	consensus.RegisterConsensusProvider(
		consensusPb.ConsensusType_RAFT,
		func(config *implconfig.ConsensusImplConfig) (protocol.ConsensusEngine, error) {
			return raft.New(config)
		},
	)

	consensus.RegisterConsensusProvider(
		consensusPb.ConsensusType_TBFT,
		func(config *implconfig.ConsensusImplConfig) (protocol.ConsensusEngine, error) {
			dp := dpos.NewNilDPoSImpl()
			return tbft.New(config, dp)
		},
	)

	consensus.RegisterConsensusProvider(
		consensusPb.ConsensusType_HOTSTUFF,
		func(config *implconfig.ConsensusImplConfig) (protocol.ConsensusEngine, error) {
			return chainedbft.New(config)
		},
	)
}
