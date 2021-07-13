/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package accesscontrol

import (
	"chainmaker.org/chainmaker-go/logger"
	"chainmaker.org/chainmaker/common/concurrentlru"
	bcx509 "chainmaker.org/chainmaker/common/crypto/x509"
	commonPb "chainmaker.org/chainmaker/pb-go/common"
	"chainmaker.org/chainmaker/protocol"
	"sync"
)

var mockAcLogger = logger.GetLogger(logger.MODULE_ACCESS)

func MockAccessControl() protocol.AccessControlProvider {
	ac := &accessControl{
		authMode:              IdentityMode,
		orgList:               &sync.Map{},
		orgNum:                0,
		resourceNamePolicyMap: &sync.Map{},
		hashType:              "",
		dataStore:             nil,
		memberCache:           concurrentlru.New(0),
		certCache:             concurrentlru.New(0),
		crl:                   sync.Map{},
		frozenList:            sync.Map{},
		opts: bcx509.VerifyOptions{
			Intermediates: bcx509.NewCertPool(),
			Roots:         bcx509.NewCertPool(),
		},
		localOrg: nil,
		log:      mockAcLogger,
	}
	return ac
}

func MockAccessControlWithHash(hashAlg string) protocol.AccessControlProvider {
	ac := &accessControl{
		authMode:              IdentityMode,
		orgList:               &sync.Map{},
		orgNum:                0,
		resourceNamePolicyMap: &sync.Map{},
		hashType:              hashAlg,
		dataStore:             nil,
		memberCache:           concurrentlru.New(0),
		certCache:             concurrentlru.New(0),
		crl:                   sync.Map{},
		frozenList:            sync.Map{},
		opts: bcx509.VerifyOptions{
			Intermediates: bcx509.NewCertPool(),
			Roots:         bcx509.NewCertPool(),
		},
		localOrg: nil,
		log:      mockAcLogger,
	}
	return ac
}

func MockSignWithMultipleNodes(msg []byte, signers []protocol.SigningMember, hashType string) ([]*commonPb.EndorsementEntry, error) {
	var ret []*commonPb.EndorsementEntry
	for _, signer := range signers {
		sig, err := signer.Sign(hashType, msg)
		if err != nil {
			return nil, err
		}
		signerSerial, err := signer.GetMember()
		if err != nil {
			return nil, err
		}
		ret = append(ret, &commonPb.EndorsementEntry{
			Signer:    signerSerial,
			Signature: sig,
		})
	}
	return ret, nil
}
