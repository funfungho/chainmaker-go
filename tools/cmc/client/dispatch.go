/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package client

import (
	"fmt"
	"sync"

	"chainmaker.org/chainmaker-go/tools/cmc/util"
	sdkPbCommon "chainmaker.org/chainmaker/pb-go/common"
	sdk "chainmaker.org/chainmaker/sdk-go"
	sdkutils "chainmaker.org/chainmaker/sdk-go/utils"
)

func Dispatch(client *sdk.ChainClient, contractName, method string, params map[string]string) {
	var (
		wgSendReq sync.WaitGroup
	)

	for i := 0; i < concurrency; i++ {
		wgSendReq.Add(1)
		go runInvokeContract(client, contractName, method, params, &wgSendReq)
	}

	wgSendReq.Wait()
}
func DispatchTimes(client *sdk.ChainClient, contractName, method string, params map[string]string) {
	var (
		wgSendReq sync.WaitGroup
	)
	times := util.MaxInt(1, sendTimes)
	wgSendReq.Add(times)
	for i := 0; i < times; i++ {
		go runInvokeContractOnce(client, contractName, method, params, &wgSendReq)
	}
	wgSendReq.Wait()
}

func runInvokeContract(client *sdk.ChainClient, contractName, method string, params map[string]string,
	wg *sync.WaitGroup) {

	defer func() {
		wg.Done()
	}()

	for i := 0; i < totalCntPerGoroutine; i++ {
		txId := sdkutils.GetRandTxId()
		resp, err := client.InvokeContract(contractName, method, txId, util.ConvertParameters(params), timeout, syncResult)
		if err != nil {
			fmt.Printf("[ERROR] invoke contract failed, %s", err.Error())
			return
		}

		if resp.Code != sdkPbCommon.TxStatusCode_SUCCESS {
			fmt.Printf("[ERROR] invoke contract failed, [code:%d]/[msg:%s]/[txId:%s]\n", resp.Code, resp.Message, txId)
			return
		}

		fmt.Printf("INVOKE contract resp, [code:%d]/[msg:%s]/[contractResult:%+v]/[txId:%s]\n", resp.Code, resp.Message,
			resp.ContractResult, txId)
	}
}

func runInvokeContractOnce(client *sdk.ChainClient, contractName, method string, params map[string]string,
	wg *sync.WaitGroup) {

	defer func() {
		wg.Done()
	}()

	txId := sdkutils.GetRandTxId()
	resp, err := client.InvokeContract(contractName, method, txId, util.ConvertParameters(params), timeout, syncResult)
	if err != nil {
		fmt.Printf("[ERROR] invoke contract failed, %s", err.Error())
		return
	}

	if resp.Code != sdkPbCommon.TxStatusCode_SUCCESS {
		fmt.Printf("[ERROR] invoke contract failed, [code:%d]/[msg:%s]/[txId:%s]\n", resp.Code, resp.Message, txId)
		return
	}

	fmt.Printf("INVOKE contract resp, [code:%d]/[msg:%s]/[contractResult:%+v]/[txId:%s]\n", resp.Code, resp.Message,
		resp.ContractResult, txId)
}
