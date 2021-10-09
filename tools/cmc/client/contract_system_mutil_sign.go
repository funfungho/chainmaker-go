package client

import (
	"chainmaker.org/chainmaker-go/tools/cmc/util"
	"chainmaker.org/chainmaker/pb-go/v2/common"
	sdkutils "chainmaker.org/chainmaker/sdk-go/v2/utils"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
)

var (
	payload *common.Payload
	pairs1  []*common.KeyValuePair
)

type ParamMultiSign struct {
	Key    string
	Value  string
	IsFile bool
}

func systemContractMultiSignCMD() *cobra.Command {
	systemContractMultiSignCmd := &cobra.Command{
		Use:   "multisign",
		Short: "system contract multi sign command",
		Long:  "system contract multi sign command",
	}

	systemContractMultiSignCmd.AddCommand(multiSignReqCMD())
	systemContractMultiSignCmd.AddCommand(multiSignVoteCMD())
	systemContractMultiSignCmd.AddCommand(multiSignQueryCMD())

	return systemContractMultiSignCmd
}

func multiSignReqCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "multisignreq",
		Short: "multi sign req",
		Long:  "multi sign req",
		RunE: func(_ *cobra.Command, _ []string) error {
			return multiSignReq()
		},
	}

	attachFlags(cmd, []string{
		flagUserSignKeyFilePath, flagUserSignCrtFilePath,
		flagConcurrency, flagTotalCountPerGoroutine, flagSdkConfPath, flagOrgId, flagChainId,
		flagParams, flagTimeout, flagUserTlsCrtFilePath, flagUserTlsKeyFilePath, flagEnableCertHash,
	})

	cmd.MarkFlagRequired(flagSdkConfPath)
	cmd.MarkFlagRequired(flagParams)

	return cmd
}

func multiSignVoteCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "multisignvote",
		Short: "multi sign vote",
		Long:  "multi sign vote",
		RunE: func(_ *cobra.Command, _ []string) error {
			return multiSignVote()
		},
	}

	attachFlags(cmd, []string{
		flagUserSignKeyFilePath, flagUserSignCrtFilePath,
		flagConcurrency, flagTotalCountPerGoroutine, flagSdkConfPath, flagOrgId, flagChainId, flagTxId,
		flagTimeout, flagUserTlsCrtFilePath, flagUserTlsKeyFilePath, flagEnableCertHash, flagAdminCrtFilePath, flagAdminKeyFilePath,
	})

	cmd.MarkFlagRequired(flagSdkConfPath)
	cmd.MarkFlagRequired(flagAdminCrtFilePath)
	cmd.MarkFlagRequired(flagAdminKeyFilePath)
	cmd.MarkFlagRequired(flagTxId)

	return cmd
}

func multiSignQueryCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "multisignquery",
		Short: "multi sign query",
		Long:  "multi sign query",
		RunE: func(_ *cobra.Command, _ []string) error {
			return multiSignQuery()
		},
	}

	attachFlags(cmd, []string{
		flagUserSignKeyFilePath, flagUserSignCrtFilePath,
		flagConcurrency, flagTotalCountPerGoroutine, flagSdkConfPath, flagOrgId, flagChainId,
		flagTimeout, flagUserTlsCrtFilePath, flagUserTlsKeyFilePath, flagEnableCertHash, flagTxId,
	})

	cmd.MarkFlagRequired(flagSdkConfPath)
	cmd.MarkFlagRequired(flagTxId)

	return cmd
}

func multiSignReq() error {
	var (
		err error
	)

	client, err := util.CreateChainClient(sdkConfPath, chainId, orgId, userTlsCrtFilePath, userTlsKeyFilePath,
		userSignCrtFilePath, userSignKeyFilePath)
	if err != nil {
		return err
	}
	defer client.Stop()
	var pms []*ParamMultiSign
	var pairs []*common.KeyValuePair
	if params != "" {
		err := json.Unmarshal([]byte(params), &pms)
		if err != nil {
			return err
		}
	}
	for _, pm := range pms {
		if pm.IsFile {
			byteCode, err := ioutil.ReadFile(pm.Value)
			if err != nil {
				panic(err)
			}
			pairs = append(pairs, &common.KeyValuePair{
				Key:   pm.Key,
				Value: byteCode,
			})

		} else {
			pairs = append(pairs, &common.KeyValuePair{
				Key:   pm.Key,
				Value: []byte(pm.Value),
			})
		}

	}
	payload = client.CreateMultiSignReqPayload(pairs)

	resp, err := client.MultiSignContractReq(payload)
	if err != nil {
		return fmt.Errorf("multi sign req failed, %s", err.Error())
	}

	fmt.Printf("multi sign req resp: %+v\n", resp)

	return nil
}

func multiSignVote() error {
	var (
		err error
	)

	client, err := util.CreateChainClient(sdkConfPath, chainId, orgId, userTlsCrtFilePath, userTlsKeyFilePath,
		userSignCrtFilePath, userSignKeyFilePath)
	if err != nil {
		return err
	}
	defer client.Stop()

	result, err := client.GetTxByTxId(txId)
	if err != nil {
		return fmt.Errorf("get tx by txid failed, %s", err.Error())
	}
	payload = result.Transaction.Payload
	endorser, err := sdkutils.MakeEndorserWithPath(adminKeyFilePaths, adminCrtFilePaths, payload)
	if err != nil {
		return fmt.Errorf("multi sign vote failed, %s", err.Error())
	}
	resp, err := client.MultiSignContractVote(payload, endorser)
	if err != nil {
		return fmt.Errorf("multi sign vote failed, %s", err.Error())
	}

	fmt.Printf("multi sign vote resp: %+v\n", resp)

	return nil
}

func multiSignQuery() error {
	var (
		err error
	)

	client, err := util.CreateChainClient(sdkConfPath, chainId, orgId, userTlsCrtFilePath, userTlsKeyFilePath,
		userSignCrtFilePath, userSignKeyFilePath)
	if err != nil {
		return err
	}
	defer client.Stop()

	resp, err := client.MultiSignContractQuery(txId)
	if err != nil {
		return fmt.Errorf("multi sign query failed, %s", err.Error())
	}

	fmt.Printf("multi sign query resp: %+v\n", resp)

	return nil
}
