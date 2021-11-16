module chainmaker.org/chainmaker-go

go 1.15

require (
	chainmaker.org/chainmaker-go/blockchain v0.0.0
	chainmaker.org/chainmaker-go/rpcserver v0.0.0
	chainmaker.org/chainmaker-go/txpool v0.0.0
	chainmaker.org/chainmaker-go/vm v0.0.0
	chainmaker.org/chainmaker/localconf/v2 v2.1.0
	chainmaker.org/chainmaker/logger/v2 v2.1.0
	chainmaker.org/chainmaker/protocol/v2 v2.1.1-0.20211116063435-570b375a296a
	chainmaker.org/chainmaker/txpool-batch/v2 v2.1.0
	chainmaker.org/chainmaker/txpool-single/v2 v2.1.0
	chainmaker.org/chainmaker/vm-docker-go v0.0.0-20211116071451-11e5554ec51c
	chainmaker.org/chainmaker/vm-evm/v2 v2.1.1-0.20211116072535-cf0d24aed8a2
	chainmaker.org/chainmaker/vm-gasm/v2 v2.1.1-0.20211116071744-294b9085dea5
	chainmaker.org/chainmaker/vm-wasmer/v2 v2.1.1-0.20211116072941-58cde49507c2
	chainmaker.org/chainmaker/vm-wxvm/v2 v2.1.1-0.20211116072003-4ba32baad065
	code.cloudfoundry.org/bytefmt v0.0.0-20200131002437-cf55d5288a48
	github.com/common-nighthawk/go-figure v0.0.0-20200609044655-c4b36f998cf2
	github.com/ethereum/go-ethereum v1.10.4 // indirect
	github.com/prometheus/client_golang v1.11.0
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
)

replace (
	chainmaker.org/chainmaker-go/accesscontrol => ./module/accesscontrol
	chainmaker.org/chainmaker-go/blockchain => ./module/blockchain
	chainmaker.org/chainmaker-go/consensus => ./module/consensus
	chainmaker.org/chainmaker-go/core => ./module/core
	chainmaker.org/chainmaker-go/net => ./module/net
	chainmaker.org/chainmaker-go/rpcserver => ./module/rpcserver
	chainmaker.org/chainmaker-go/snapshot => ./module/snapshot
	chainmaker.org/chainmaker-go/subscriber => ./module/subscriber
	chainmaker.org/chainmaker-go/sync => ./module/sync
	chainmaker.org/chainmaker-go/txpool => ./module/txpool
	chainmaker.org/chainmaker-go/vm => ./module/vm
	github.com/libp2p/go-libp2p-core => chainmaker.org/chainmaker/libp2p-core v1.0.0
	google.golang.org/grpc v1.40.0 => google.golang.org/grpc v1.26.0
)
