module chainmaker.org/chainmaker-go/rpcserver

go 1.15

require (
	chainmaker.org/chainmaker-go/blockchain v0.0.0
	chainmaker.org/chainmaker-go/monitor v0.0.0
	chainmaker.org/chainmaker-go/subscriber v0.0.0
	chainmaker.org/chainmaker/common/v2 v2.0.1-0.20210909033927-2a4cfc146579
	chainmaker.org/chainmaker/localconf/v2 v2.0.0-20210913154622-9f9774ed7d1b
	chainmaker.org/chainmaker/logger/v2 v2.0.0-20210907134457-53647922a89d
	chainmaker.org/chainmaker/pb-go/v2 v2.0.1-0.20210901132412-435b75070bf2
	chainmaker.org/chainmaker/protocol/v2 v2.0.1-0.20210910112253-04256ae9c5ed
	chainmaker.org/chainmaker/store/v2 v2.0.0-20210911172624-6ea7e13fd4d6
	chainmaker.org/chainmaker/utils/v2 v2.0.0-20210907033606-84c6c841cbdb
	chainmaker.org/chainmaker/vm-native v0.0.0-20210914041545-8232478a91c4 // indirect
	github.com/gogo/protobuf v1.3.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/prometheus/client_golang v1.9.0
	golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba
	golang.org/x/tools v0.1.5 // indirect
	google.golang.org/grpc v1.37.0
)

replace (
	chainmaker.org/chainmaker-go/accesscontrol => ../accesscontrol
	chainmaker.org/chainmaker-go/blockchain => ../blockchain
	chainmaker.org/chainmaker-go/consensus => ../consensus
	chainmaker.org/chainmaker-go/consensus/dpos => ./../consensus/dpos
	chainmaker.org/chainmaker-go/core => ../core

	chainmaker.org/chainmaker-go/monitor => ../monitor
	chainmaker.org/chainmaker-go/net => ../net
	chainmaker.org/chainmaker-go/rpcserver => ../rpcserver
	chainmaker.org/chainmaker-go/snapshot => ../snapshot
	chainmaker.org/chainmaker-go/subscriber => ../subscriber
	chainmaker.org/chainmaker-go/sync => ../sync
	chainmaker.org/chainmaker-go/txpool => ../txpool
	github.com/libp2p/go-libp2p => ../net/p2p/libp2p
	github.com/libp2p/go-libp2p-core => ../net/p2p/libp2pcore
	github.com/libp2p/go-libp2p-pubsub => ../net/p2p/libp2ppubsub
)
