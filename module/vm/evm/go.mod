module chainmaker.org/chainmaker-go/evm

go 1.15

require (
	chainmaker.org/chainmaker-go/logger v0.0.0
	chainmaker.org/chainmaker-go/utils v0.0.0
	chainmaker.org/chainmaker-go/wasmer v0.0.0
	chainmaker.org/chainmaker/common v0.0.0-20210709154839-e2c8e4fc62b4
	chainmaker.org/chainmaker/pb-go v0.0.0-20210709093937-9b3b422e24b1
	chainmaker.org/chainmaker/protocol v0.0.0-20210709171355-90bbfd38e3cc
	github.com/ethereum/go-ethereum v1.10.3
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2

)

replace (
	chainmaker.org/chainmaker-go/localconf => ../../conf/localconf
	chainmaker.org/chainmaker-go/logger => ../../logger

	chainmaker.org/chainmaker-go/store => ../../store
	chainmaker.org/chainmaker-go/utils => ../../utils
	chainmaker.org/chainmaker-go/wasi => ../wasi
	chainmaker.org/chainmaker-go/wasmer => ../wasmer
)
