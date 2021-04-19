/*
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package raft

// SnapshotHeight stores block height in raft snapshot.
type SnapshotHeight struct {
	Height int64
}

// AdditionalData contains consensus specified data to be store in block
type AdditionalData struct {
	Signature    []byte
	AppliedIndex uint64
}
